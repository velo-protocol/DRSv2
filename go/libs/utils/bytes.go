package utils

import "bytes"

func StringToByte32(s string) [32]byte {
	var byteArr [32]byte
	copy(byteArr[:], s)
	return byteArr
}

func Byte32ToString(b [32]byte) string {
	return string(bytes.Trim(b[:], "\x00"))
}

func BytesToString(b []byte) string {
	return string(bytes.Trim(b[:], "\x00"))
}

func BytesToBytes32(b []byte) [32]byte {
	var byteArr [32]byte
	copy(byteArr[:], b)
	return byteArr
}
