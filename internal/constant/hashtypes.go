package constant

type HashType struct {
	ID   int
	Name string
}

const (
	HashMD5 int = iota
	HashSHA224
	HashSHA256
	HashSHA384
	HashSHA512
)

var HashTypeList = []HashType{
	{HashMD5, "MD5"},
	{HashSHA224, "SHA224"},
	{HashSHA256, "SHA256"},
	{HashSHA384, "SHA384"},
	{HashSHA512, "SHA512"},
}
