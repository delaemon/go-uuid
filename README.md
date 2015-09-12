# go-uuidv4
[![Build Status](https://travis-ci.org/delaemon/go-uuidv4.svg?branch=master)](https://travis-ci.org/delaemon/go-uuidv4) [![GoCover](http://gocover.io/_badge/github.com/delaemon/go-uuidv4)](http://gocover.io/github.com/delaemon/go-uuidv4) [![GoDoc](https://godoc.org/github.com/delaemon/go-uuidv4?status.png)](https://godoc.org/github.com/delaemon/go-uuidv4)

##Description
This package to generate a UUID according to the RFC 4122 standard. Only support for version 4 are built-in.

##Installation
This package can be installed with the go get command:
```sh
go get github.com/delaemon/go-uuidv4
```

##Usage
```
package main

import (
	"github.com/delaemon/go-uuidv4"
)

func main() {
	uuid, err := uuidv4.Generate()
}
```

