package sha512

type SHA512 struct {
	isUsed  bool
	payload []byte
	hash    []byte
}

func New() *SHA512 {
	return &SHA512{}
}
