package sha224

type SHA224 struct {
	isUsed  bool
	payload []byte
	hash    []byte
}

func New() *SHA224 {
	return &SHA224{}
}
