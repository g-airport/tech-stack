// Copyright © 2015-2018 Anker Innovations Technology Limited All Rights Reserved.
package varint

const maxByte = 10

func EncodeVarint(x uint64) []byte {
	var buf [maxByte]byte
	var n int
	for n = 0; x > 127; n++ {
		buf[n] = 0x80 | uint8(x&0x7F)
		x >>= 7
	}
	buf[n] = uint8(x)
	n++
	return buf[0:n]

}

func DecodeVarint(buf []byte) (x uint64, n int) {
	for shift := uint(0); shift < 64; shift += 7 {
		if n >= len(buf) {
			return 0, 0
		}
		b := uint64(buf[n])
		n++
		x |= (b & 0x7F) << shift
		if (b & 0x80) == 0 {
			return x, n
		}
	}
	// The number is too large to represent in a 64-bit value.
	return 0, 0
}