package build

import (
	"archive/tar"
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/converged-computing/container-crafter/pkg/config"
	"github.com/converged-computing/container-crafter/pkg/utils"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

const (
	// Max layer size is 10GB, 10 billion bytes
	maxLayerSize = 10000000000
)

// A BuildMatrix generates a range of sizes and layer counts
type BuildMatrix struct {
	URI        string
	Study      *config.StudyConfig
	ConfigFile string
}

// NewBuildMatrix generates a new build matrix spec
func NewBuildMatrix(uri string, configFile string) (*BuildMatrix, error) {

	// Do we have a config file?
	cfg, err := config.Load(configFile)
	if err != nil {
		return nil, err
	}

	// URI provided in configuration file
	if cfg.URI != "" {
		uri = cfg.URI
	}

	// Remove any tag provided in the URI, we generate them based on the matrix coordinate
	parts := strings.Split(uri, ":")
	matrix := &BuildMatrix{
		URI:        parts[0],
		ConfigFile: configFile,
		Study:      cfg,
	}
	return matrix, nil
}

// Build generates the matrix and runs the builds. We will do these in
// serial for now, because my computer can only handle that.
func (m *BuildMatrix) Build() error {

	// Final list of tags
	tags := []string{}

	// Do builds through image sizes and layers for each
	for _, imageSize := range m.Study.IterSizes() {

		// busybox is 5MB, adjust for it so total image size includes base
		size := imageSize - 5000000

		for _, layerCount := range m.Study.IterLayers() {
			layerSize := size / layerCount

			// Assemble tag based on layers and total size
			tag := fmt.Sprintf("%s:%d-layers-size-%d-bytes", m.URI, layerCount, size)
			err := m.buildImage(tag, layerSize, layerCount)
			if err != nil {
				return err
			}
			tags = append(tags, tag)
		}
	}

	// Tell the user about it
	fmt.Println("⭐ Final image set built:")
	for _, tag := range tags {
		fmt.Println(tag)
	}
	return nil
}

// writeDockerfile writes the Dockerfile to the filesystem
func (m BuildMatrix) writeDockerfile(content string) (string, error) {
	f, err := os.CreateTemp("", "Dockerfile-")
	if err != nil {
		return "", err
	}
	defer f.Close()

	// write data to the temporary file
	_, err = f.Write([]byte(content))
	return f.Name(), err
}

// BuildImage builds the actual image with the layer size and layers
// The URI (tag) is returned to save for later
func (m BuildMatrix) buildImage(tag string, layerSize, layers int) error {
	if int(layerSize) > maxLayerSize {
		fmt.Printf("Warning: %d is larger than layer limit of %d, skipping %d layers\n", layerSize, maxLayerSize, layers)
		return nil
	}

	fmt.Printf("Building image %s [layers:%d, layerSize:%d]\n", tag, layers, layerSize)

	// Base image is busybox, 5MB
	dockerfile := "FROM busybox"

	// Make the filename unique too
	textFile := utils.GenerateRandomName()

	for i := 0; i < int(layers); i++ {
		dockerfile += fmt.Sprintf("\nRUN head -c %d /dev/zero > %s-%d.txt", layerSize, textFile, i)
	}

	filename, err := m.writeDockerfile(dockerfile)
	if err != nil {
		return err
	}
	defer os.Remove(filename)

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return fmt.Errorf("unable to init client %s", err)
	}
	defer cli.Close()

	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)
	defer tw.Close()

	tarHeader := &tar.Header{
		Name: filename,
		Size: int64(len(dockerfile)),
	}
	err = tw.WriteHeader(tarHeader)
	if err != nil {
		return fmt.Errorf("unable to write tar header %s", err)
	}
	_, err = tw.Write([]byte(dockerfile))
	if err != nil {
		return fmt.Errorf("unable to write tar body %s", err)
	}
	dockerFileTarReader := bytes.NewReader(buf.Bytes())

	imageBuildResponse, err := cli.ImageBuild(
		ctx,
		dockerFileTarReader,
		types.ImageBuildOptions{
			Context:    dockerFileTarReader,
			Tags:       []string{tag},
			Dockerfile: filename,
			Remove:     true,
		},
	)
	if err != nil {
		return fmt.Errorf("unable to build docker image %s", err)
	}
	defer imageBuildResponse.Body.Close()
	_, err = io.Copy(os.Stdout, imageBuildResponse.Body)
	return err
}
