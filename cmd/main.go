package main

import (
	"fmt"

	sha "github.com/kevinfernaldy/go-hash/internal/app/usecase/sha2/sha256"
)

func main() {
	sha := sha.SHA256{}

	sha.Update("hello world")

	hash, err := sha.Digest()
	if err != nil {
		panic(err)
	}

	fmt.Println(hash)
}
