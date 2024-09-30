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
	uri := createCmd.String("u", "uri", &argparse.Options{Default: "test", Help: "Unique resource identifier (URI) to use, without tag"})

	// Read most metadata from a config file
	config := createCmd.String("c", "config", &argparse.Options{Help: "Study yaml configuration file to replace arguments above"})

	// Now parse the arguments
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Println(Header)
		fmt.Println(parser.Usage(err))
		return
	}

	if *config == "" {
		log.Fatalf("You must provide a config file with --config to indicate preferences for your builds.")
	}
	if createCmd.Happened() {
		err := build.Run(*uri, *config)
		if err != nil {
			log.Fatalf("Issue with create: %s\n", err)
		}
	} else {
		fmt.Println(Header)
		fmt.Println(parser.Usage(nil))
	}
}
