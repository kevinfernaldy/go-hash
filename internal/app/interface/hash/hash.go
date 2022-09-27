package hash

import (
	"errors"

	"github.com/kevinfernaldy/go-hash/internal/app/usecase/md5"
	"github.com/kevinfernaldy/go-hash/internal/app/usecase/sha2/sha224"
	"github.com/kevinfernaldy/go-hash/internal/app/usecase/sha2/sha256"
	"github.com/kevinfernaldy/go-hash/internal/app/usecase/sha2/sha384"
	"github.com/kevinfernaldy/go-hash/internal/app/usecase/sha2/sha512"

	"github.com/kevinfernaldy/go-hash/internal/constant"
)

func CreateHash(hashType int) (Hasher, error) {
	switch hashType {
	case constant.HashSHA224:
		{
			return sha224.New(), nil
		}
	case constant.HashSHA256:
		{
			return sha256.New(), nil
		}
	case constant.HashSHA384:
		{
			return sha384.New(), nil
		}
	case constant.HashSHA512:
		{
			return sha512.New(), nil
		}
	case constant.HashMD5:
		{
			return md5.New(), nil
		}
	default:
		{
			return nil, errors.New("unknown hash type requested")
		}
	}

}
