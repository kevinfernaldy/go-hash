package md5

type MD5 struct {
	isUsed  bool
	payload []byte
	hash    []byte
}

func New() *MD5 {
	return &MD5{}
}
