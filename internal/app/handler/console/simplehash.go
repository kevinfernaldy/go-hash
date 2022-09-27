package console

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/kevinfernaldy/go-hash/internal/app/interface/hash"
	"github.com/kevinfernaldy/go-hash/internal/constant"
)

func (c *Console) getHashType() constant.HashType {
	hashList := &constant.HashTypeList

	fmt.Println("Select the hash method you want to use")
	for index, item := range *hashList {
		fmt.Printf("%d. %s\n", index+1, item.Name)
	}

	selection := c.getSelection(1, len(*hashList)) - 1

	return (*hashList)[selection]
}

func (c *Console) getPayload() string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the message you want to hash: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}

	return strings.TrimRight(text, c.newline)
}

func (c *Console) handleSimpleHash() {
	// Get the hash ID
	c.clearScreen()
	hashType := c.getHashType()

	// Get the payload
	c.clearScreen()
	fmt.Printf("Algorithm : %s\n", hashType.Name)
	fmt.Println()
	payload := c.getPayload()

	// Do the hash
	hash, err := hash.CreateHash(hashType.ID)
	if err != nil {
		panic(err)
	}
	hash.Update(payload)
	hashedPayload := hash.Digest()

	// Print the result
	c.clearScreen()
	fmt.Printf("Algorithm : %s\n", hashType.Name)
	fmt.Printf("Payload   : %s\n", payload)
	fmt.Println()
	fmt.Printf("Hash      : %s\n", hashedPayload)
}
