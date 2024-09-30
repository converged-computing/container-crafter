package config

import (
	"os"

	yaml "gopkg.in/yaml.v3"
)

// A StudyConfig holds information for sizes and layers
// We will require both
type StudyConfig struct {
	Sizes  []StudySize   `yaml:"sizes,omitempty"`
	Layers []StudyLayers `yaml:"layers,omitempty"`
	URI    string        `yaml:"uri,omitempty"`
}

// Sizes are in bytes, the total size for the container
type StudySize struct {
	Total int `yaml:"total"`
}

// Layers are the number of layers to do for each size
// The build will only happen for layer / image size combinations that
// produce layer sizes <= 10MB (registry limit)
type StudyLayers struct {
	Exact int `yaml:"exact"`
	Min   int `yaml:"min"`
	Max   int `yaml:"max"`
}

// read the config and return a config type
func readConfig(yamlStr []byte) (*StudyConfig, error) {

	// First unmarshall into generic structure
	data := StudyConfig{}
	err := yaml.Unmarshal(yamlStr, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// Iterate through Layer sizes
func (s *StudyConfig) IterLayers() []int {
	layers := []int{}
	for _, layerSpec := range s.Layers {
		// Check for exact count first
		if layerSpec.Exact != 0 {
			layers = append(layers, layerSpec.Exact)
			continue
		}
		if layerSpec.Min != 0 && layerSpec.Max != 0 && layerSpec.Max > layerSpec.Min {
			for i := layerSpec.Min; i <= layerSpec.Max; i++ {
				layers = append(layers, i)
			}
		}
	}
	return layers
}

// Iterate through Image Sizes
func (s *StudyConfig) IterSizes() []int {
	sizes := []int{}
	for _, size := range s.Sizes {
		if size.Total != 0 {
			sizes = append(sizes, size.Total)
		}
	}
	return sizes
}

// Load a study config yaml file
func Load(yamlFile string) (*StudyConfig, error) {
	yamlContent, err := os.ReadFile(yamlFile)
	if err != nil {
		return nil, err
	}
	return readConfig(yamlContent)
}
