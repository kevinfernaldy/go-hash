package main

import (
	"fmt"

	"github.com/kevinfernaldy/go-hash/internal/app/interface/hash"
	"github.com/kevinfernaldy/go-hash/internal/constant"
)

func main() {
	sha224, _ := hash.CreateHash(constant.HashSHA224)
	sha256, _ := hash.CreateHash(constant.HashSHA256)
	sha384, _ := hash.CreateHash(constant.HashSHA384)
	sha512, _ := hash.CreateHash(constant.HashSHA512)
	md5, _ := hash.CreateHash(constant.HashMD5)

	sha224.Update("Hello world")
	sha256.Update("Hello world")
	sha384.Update("Hello world")
	sha512.Update("Hello world")
	md5.Update("Hello world")

	fmt.Println("Plain text: Hello world")
	fmt.Printf("SHA224: %s\n", sha224.Digest())
	fmt.Printf("SHA256: %s\n", sha256.Digest())
	fmt.Printf("SHA384: %s\n", sha384.Digest())
	fmt.Printf("SHA512: %s\n", sha512.Digest())
	fmt.Printf("MD5   : %s\n", md5.Digest())
}
