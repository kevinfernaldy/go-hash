package hash

type Hasher interface {
	Update(payload string) error
	Digest() string
}
