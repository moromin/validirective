# validirective
`validirective` detects directives that have been commented out by mistake.

## Install
```bash
go install github.com/moromin/validirective/cmd/validirective@latest
```

## Usage
```bash
go vet -vettool=`which validirective` <target-file>
```

## Example
```go
// target file
// sample.go
package sample

import _ "embed"

//go:embed sample.json
var s1 []byte

// go:embed sample.json
var s2 []byte
```
```bash
$ go vet -vettool=`which validirective` sample.go
# command-line-arguments
./sample.go:8:1: this directive has become a comment...
```
