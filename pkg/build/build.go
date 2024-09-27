package build

import "strings"

// A BuildMatrix generates a range of sizes and layer counts
type BuildMatrix struct {
	Layers    int
	ImageSize int
	URI       string
}

// NewBuildMatrix generates a new build matrix spec
func NewBuildMatrix(uri string, layers, imageSize int) *BuildMatrix {

	// Remove any tag provided in the URI, we generate them based on the matrix coordinate
	// We could (maybe should) do more validation of the Uri here.
	parts := strings.Split(uri, ":")
	return &BuildMatrix{URI: parts[0], ImageSize: imageSize, Layers: layers}
}

// Build generates the matrix and runs the builds. We will do these in
// serial for now, because my computer can only handle that.
func (*BuildMatrix) Build() error {
	return nil
}

// Push the URI with all tags to a registry.
// This assumes you have authentication
func (*BuildMatrix) Push() error {
	return nil
}
