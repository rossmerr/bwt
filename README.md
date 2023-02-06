# Burrows-Wheeler Transform (BWT)

![Go](https://github.com/rossmerr/bwt/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/rossmerr/bwt)](https://goreportcard.com/report/github.com/rossmerr/bwt)
[![Read the Docs](https://pkg.go.dev/badge/golang.org/x/pkgsite)](https://pkg.go.dev/github.com/rossmerr/bwt)

Rearranges a character string into runs of similar characters. This is useful for compression, since it tends to be easy to compress a string that has runs of repeated characters by techniques such as move-to-front transform and run-length encoding.
