package utils

import (
	"encoding/binary"
	"errors"
	"io"
	"os"
)

// Check for any error while reading the file
func Check(e error) {
	if e != nil && e != io.EOF {
		panic(e)
	} else if e == io.EOF {
		err := errors.New("Reached unexpected EOF. Either file is corrupt or does not follow RosBag 2.0 spec")
		panic(err)
	}
}

// CheckEOF checks whether EOF has occurred
func CheckEOF(e error) bool {
	if e == io.EOF {
		return true
	}
	return false
}

// ReadInt32 reads a 32 bit integer in Little Endina format
func ReadInt32(f *os.File) uint32 {
	val := make([]byte, 4)
	_, err := f.Read(val)
	Check(err)
	return binary.LittleEndian.Uint32(val)
}
