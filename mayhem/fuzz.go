package fuzz

import "github.com/ipfs/go-unixfs"

func mayhemit(bytes []byte) int {

    unixfs.WrapData(bytes)
    return 0
}

func Fuzz(data []byte) int {
    _ = mayhemit(data)
    return 0
}