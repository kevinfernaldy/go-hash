package sha384

type SHA384 struct {
	isUsed  bool
	payload []byte
	hash    []byte
}

func New() *SHA384 {
	return &SHA384{}
}
