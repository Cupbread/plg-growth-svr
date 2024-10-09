package util

import (
	"encoding/binary"
)

func Md5(message []byte) ([16]byte, error) {
	var s = [64]int{
		7, 12, 17, 22, 7, 12, 17, 22, 7, 12, 17, 22, 7, 12, 17, 22,
		5, 9, 14, 20, 5, 9, 14, 20, 5, 9, 14, 20, 5, 9, 14, 20,
		4, 11, 16, 23, 4, 11, 16, 23, 4, 11, 16, 23, 4, 11, 16, 23,
		6, 10, 15, 21, 6, 10, 15, 21, 6, 10, 15, 21, 6, 10, 15, 21,
	}

	var K = [64]uint32{
		0xd76aa478, 0xe8c7b756, 0x242070db, 0xc1bdceee,
		0xf57c0faf, 0x4787c62a, 0xa8304613, 0xfd469501,
		0x698098d8, 0x8b44f7af, 0xffff5bb1, 0x895cd7be,
		0x6b901122, 0xfd987193, 0xa679438e, 0x49b40821,
		0xf61e2562, 0xc040b340, 0x265e5a51, 0xe9b6c7aa,
		0xd62f105d, 0x02441453, 0xd8a1e681, 0xe7d3fbc8,
		0x21e1cde6, 0xc33707d6, 0xf4d50d87, 0x455a14ed,
		0xa9e3e905, 0xfcefa3f8, 0x676f02d9, 0x8d2a4c8a,
		0xfffa3942, 0x8771f681, 0x6d9d6122, 0xfde5380c,
		0xa4beea44, 0x4bdecfa9, 0xf6bb4b60, 0xbebfbc70,
		0x289b7ec6, 0xeaa127fa, 0xd4ef3085, 0x04881d05,
		0xd9d4d039, 0xe6db99e5, 0x1fa27cf8, 0xc4ac5665,
		0xf4292244, 0x432aff97, 0xab9423a7, 0xfc93a039,
		0x655b59c3, 0x8f0ccc92, 0xffeff47d, 0x85845dd1,
		0x6fa87e4f, 0xfe2ce6e0, 0xa3014314, 0x4e0811a1,
		0xf7537e82, 0xbd3af235, 0x2ad7d2bb, 0xeb86d391,
	}

	var h0 uint32 = 0x67452301
	var h1 uint32 = 0xefcdab89
	var h2 uint32 = 0x98badcfe
	var h3 uint32 = 0x10325476

	originalLength := uint64(len(message) * 8)
	message = append(message, 0x80)
	for len(message)%64 != 56 {
		message = append(message, 0x00)
	}

	lengthBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(lengthBytes, originalLength)
	message = append(message, lengthBytes...)

	for i := 0; i < len(message); i += 64 {
		var M [16]uint32
		for j := 0; j < 16; j++ {
			M[j] = binary.LittleEndian.Uint32(message[i+j*4:])
		}

		a, b, c, d := h0, h1, h2, h3

		for j := 0; j < 64; j++ {
			var f, g uint32
			switch {
			case j < 16:
				f = (b & c) | ((^b) & d)
				g = uint32(j)
			case j < 32:
				f = (d & b) | ((^d) & c)
				g = (5*uint32(j) + 1) % 16
			case j < 48:
				f = b ^ c ^ d
				g = (3*uint32(j) + 5) % 16
			default:
				f = c ^ (b | (^d))
				g = (7 * uint32(j)) % 16
			}

			temp := d
			d = c
			c = b
			b = b + leftRotate((a+f+K[j]+M[g]), s[j])
			a = temp
		}

		h0 += a
		h1 += b
		h2 += c
		h3 += d
	}
	var result [16]byte
	binary.LittleEndian.PutUint32(result[0:4], h0)
	binary.LittleEndian.PutUint32(result[4:8], h1)
	binary.LittleEndian.PutUint32(result[8:12], h2)
	binary.LittleEndian.PutUint32(result[12:16], h3)

	return result, nil
}

func leftRotate(x uint32, c int) uint32 {
	return (x << c) | (x >> (32 - c))
}
