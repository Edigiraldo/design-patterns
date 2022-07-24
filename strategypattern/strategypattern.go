package strategypattern

import (
	"fmt"
)

type PasswordProtector struct {
	username      string
	password      string
	hashAlgorithm HashAlgorithm
}

func (p *PasswordProtector) SetHashAlgorithm(hash HashAlgorithm) {
	p.hashAlgorithm = hash
}

func (p *PasswordProtector) Hash() {
	p.hashAlgorithm.Hash(p)
}

func NewPasswordProtector(username string, password string, hash HashAlgorithm) *PasswordProtector {
	return &PasswordProtector{
		username:      username,
		password:      password,
		hashAlgorithm: hash,
	}
}

type HashAlgorithm interface {
	Hash(*PasswordProtector)
}

type MD5 struct{}

func (md5 *MD5) Hash(p *PasswordProtector) {
	fmt.Printf("Hashing using MD5 for password of %s\n", p.username)
}

type SHA struct{}

func (sha *SHA) Hash(p *PasswordProtector) {
	fmt.Printf("Hashing using SHA for password of %s\n", p.username)
}

func RunExample() {
	sha := &SHA{}
	md5 := &MD5{}

	passwordProtector := NewPasswordProtector("Francisco", "email password", sha)
	passwordProtector.Hash()
	passwordProtector.SetHashAlgorithm(md5)
	passwordProtector.Hash()
}
