package build

import (
	"github.com/converged-computing/container-crafter/pkg/build"
)

// Create a matrix of builds with a particular layer count and size
func Run(uri, configFile string) error {
	builder, err := build.NewBuildMatrix(uri, configFile)
	if err != nil {
		return err
	}
	return builder.Build()
}
