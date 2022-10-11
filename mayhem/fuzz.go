package fuzz

import "github.com/ipfs/go-unixfs"

func mayhemit(bytes []byte) int {

    _ = unifxs.FromBytes(bytes)
    return 0
}

func Fuzz(data []byte) int {
    _ = mayhemit(data)
    return 0
}