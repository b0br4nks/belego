package belego

import (
    "encoding/binary"
    "fmt"
)

type Endianness int

const (
    Big Endianness = iota
    Little
)

func init() {
    fmt.Println("Belego is a program to detect the endianness. (Big or Little)")
}

func Endian() Endianness {
    var i int32 = 0x01020304
    b := make([]byte, 4)
    binary.LittleEndian.PutUint32(b, uint32(i))
    if b[0] == 0x04 {
        return Little
    }
    return Big
}
