[![Build Status](https://travis-ci.org/chambbj/content-hasher.png?branch=master)](https://travis-ci.org/chambbj/content-hasher)

# content-hasher

A pure Go Dropbox Content Hasher.

See https://www.dropbox.com/developers/reference/content-hash for more
information.

# Installation

``` console
$ go get gopkg.in/chambbj/content-hasher.v0
```

# Example

Here is a sample program.

``` go
package main

import (
	"fmt"
	"os"

	hasher "gopkg.in/chambbj/content-hasher.v0"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	fmt.Printf("%x\n", hasher.Compute(f))
}
```

Build it.

``` console
$ go build main.go
```

Download the test image used in the Dropbox Content Hasher documentation
[here](https://www.dropbox.com/static/images/developers/milky-way-nasa.jpg), and
run it through the sample program and verify that the resulting hash is

```
$ sample-program milky-way-nasa.jpg
485291fa0ee50c016982abbfa943957bcd231aae0492ccbaa22c58e3997b35e0
```
