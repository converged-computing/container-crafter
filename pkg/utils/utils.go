package utils

import (
	"fmt"

	"github.com/moby/moby/pkg/namesgenerator"
	"golang.org/x/exp/rand"
)

// GenerateRandomName uses docker's generator and adds a little
// more randomness to the end
func GenerateRandomName() string {
	name := namesgenerator.GetRandomName(2)
	contenders := "abcdefghijklmnopqrstuvwxyz1234567890"

	// Getting random character
	randomChars := []byte{}
	for i := 0; i < 8; i++ {
		randomChars = append(randomChars, contenders[rand.Intn(len(contenders))])
	}
	fmt.Println(randomChars)
	name = fmt.Sprintf("%s-%s", name, string(randomChars))
	fmt.Println(name)
	return name
}
