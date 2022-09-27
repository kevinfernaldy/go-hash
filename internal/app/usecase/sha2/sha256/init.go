package sha256

type SHA256 struct {
	isUsed  bool
	payload []byte
	hash    []byte
}

func New() *SHA256 {
	return &SHA256{}
}
