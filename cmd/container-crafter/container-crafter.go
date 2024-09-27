package main

import (
	"fmt"
	"log"
	"os"

	"github.com/akamensky/argparse"
	"github.com/converged-computing/container-crafter/cmd/container-crafter/build"
	"github.com/converged-computing/container-crafter/pkg/types"
)

// I know this text is terrible, just having fun for now
var (
	Header = `              
   _   _   _   _   _   _   _   _   _     _   _   _   _   _   _   _  
  / \ / \ / \ / \ / \ / \ / \ / \ / \   / \ / \ / \ / \ / \ / \ / \ 
 ( c | o | n | t | a | i | n | e | r ) ( c | r | a | f | t | e | r )
  \_/ \_/ \_/ \_/ \_/ \_/ \_/ \_/ \_/   \_/ \_/ \_/ \_/ \_/ \_/ \_/  
`
)

func RunVersion() {
	fmt.Printf("⭐️ container-crafter version %s\n", types.Version)
}

func main() {

	parser := argparse.NewParser("container-crafter", "Container generator based on size and layer count matrix")
	createCmd := parser.NewCommand("create", "Create matrix of containers according to a max size and number of layers")

	// Create arguments
	uri := createCmd.String("u", "uri", &argparse.Options{Help: "Unique resource identifier (URI) to use, without tag"})
	layers := createCmd.Int("l", "layers", &argparse.Options{Help: "Number of layers"})
	imageSize := createCmd.Int("s", "image-size", &argparse.Options{Help: "Total image size to generate"})

	// Now parse the arguments
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Println(Header)
		fmt.Println(parser.Usage(err))
		return
	}

	if createCmd.Happened() {
		err := build.Run(*uri, *layers, *imageSize)
		if err != nil {
			log.Fatalf("Issue with create: %s\n", err)
		}
	} else {
		fmt.Println(Header)
		fmt.Println(parser.Usage(nil))
	}
}
