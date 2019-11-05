package secret

import "fmt"

// Secret to encrypt/decrypt
type Secret struct {
	Key       string
	CipherHex string
}

func (s *Secret) String() string {
	return fmt.Sprintf("%s: %s", s.Key, s.CipherHex)
}
