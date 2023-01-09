package fuzzUnwrapData

import "github.com/ipfs/go-unixfs"

func mayhemit(bytes []byte) int {

    unixfs.UnwrapData(bytes)
    return 0
}

func Fuzz(data []byte) int {
    _ = mayhemit(data)
    return 0
}