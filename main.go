package hasher

import (
	"crypto/sha256"
	"os"
)

// Compute computes a Dropbox Content Hash according to
// https://www.dropbox.com/developers/reference/content-hash
func Compute(f *os.File) [32]byte {
	var shas []byte
	for {
		sum, n, _ := func(f *os.File) ([]byte, int, error) {
			bytes := make([]byte, 4*1024*1024)
			n, err := f.Read(bytes)
			if err != nil {
				return nil, n, err
			}
			sum := sha256.Sum256(bytes[0:n])
			return sum[:], n, nil
		}(f)
		if n == 0 {
			break
		}
		shas = append(shas, sum[:]...)
	}

	return sha256.Sum256(shas)
}
