# Burrows-Wheeler Transform (BWT)

[![Go](https://github.com/rossmerr/bwt/actions/workflows/go.yml/badge.svg)](https://github.com/rossmerr/bwt/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/rossmerr/bwt)](https://goreportcard.com/report/github.com/rossmerr/bwt)
[![Read the Docs](https://pkg.go.dev/badge/golang.org/x/pkgsite)](https://pkg.go.dev/github.com/rossmerr/bwt)

Rearranges a character string into runs of similar characters. This is useful for compression, since it tends to be easy to compress a string that has runs of repeated characters by techniques such as move-to-front transform and run-length encoding.

Given the following int `abaaba`

The Burrows-Wheeler Matrix would look like the this :-

```
abaaba
aabaab
aabaab
abaaba
abaaba
baabaa
baabaa
```

The last column is then `abbaaa`, were the `` rune is the End-of-Text code.

```go
matrix, err := bwt.Matrix("abaaba") // 'matrix' is a [][]rune
```

```go
last, err := bwt.Last("abaaba") // 'last' is a []rune

fmt.Println(string(last)) // abbaaa
```

```go
first, last, err := bwt.FirstLast("abaaba")  // 'first' is a []rune

fmt.Println(string(first)) // aaaabb

fmt.Println(string(last)) // abbaaa
```

```go
str := "abaaba"
text := []rune(str)

first, last, sa, err := bwt.FirstLastSuffix(str)  // 'sa' is a suffixarray.Suffix

fmt.Println(string(first)) // aaaabb

fmt.Println(string(last)) // abbaaa

// You want to find the original offset of the first 'b' in the 'str'
// 6 is the index of rune 'b' from the first column,
//
// you could Enumerate over 'sa' to find the index
// and the first column of the BWT has consecutivity
// so we would know the first 'b' must have been the first 'b' in the 'str'
offset := sa.Get(6)

fmt.Println(offset) // 1
fmt.Println(string(text[offset])) // b

```
