package build

import (
	"github.com/converged-computing/container-crafter/pkg/build"
)

// Create a matrix of builds with a particular layer count and size
func Run(uri string, layers, imageSize int) error {
	builder := build.NewBuildMatrix(uri, layers, imageSize)
	return builder.Build()
}
