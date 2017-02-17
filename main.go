// Copyright 2017, Bradley J Chambers (brad.chambers@gmail.com)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
