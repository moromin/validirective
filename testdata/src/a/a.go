package a

import _ "embed"

//go:embed sample.json
var s1 []byte

// go:embed sample.json
var s2 []byte

//go:embed sample.json
//  go:embed sample.json
var s3 []byte
