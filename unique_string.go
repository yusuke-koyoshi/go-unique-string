package unique_string

import (
	"math/bits"
	"strings"
)

const (
	c1 = 2246822507
	c2 = 3266489909
	c3 = 2869860233
	c4 = 597399067
	c5 = 1444728091
	c6 = 197830471
	c7 = -1425107063
)

func base32Encode(input uint64) string {
	text := "abcdefghijklmnopqrstuvwxyz234567"

	var sb strings.Builder
	for i := 0; i < 13; i++ {
		char := string(text[input>>59])
		sb.WriteString(char)
		input = input << 5
	}
	return sb.String()
}

func murmurHash64(data []byte) uint64 {
	size := len(data)
	var k1, k2 uint32
	var i int
	for i = 0; i+7 < size; i += 8 {
		p1 := uint32(uint32(data[i]) | (uint32(data[i+1]) << 8) | (uint32(data[i+2]) << 16) | (uint32(data[i+3]) << 24))
		p2 := uint32(uint32(data[i+4]) | (uint32(data[i+5]) << 8) | (uint32(data[i+6]) << 16) | (uint32(data[i+7]) << 24))
		p1 = p1 * c4
		p1 = bits.RotateLeft32(p1, 15)
		p1 = p1 * uint32(c3)
		k1 = k1 ^ p1
		k1 = bits.RotateLeft32(k1, 19)
		k1 = k1 + k2
		k1 = k1*5 + c5
		p2 = p2 * uint32(c3)
		p2 = bits.RotateLeft32(p2, 17)
		p2 = p2 * c4
		k2 = k2 ^ p2
		k2 = bits.RotateLeft32(k2, 13)
		k2 = k2 + k1
		k2 = k2*5 + c6
	}

	o := size - i

	if o > 0 {
		h1 := uint32(uint32(data[i]) | (uint32(data[i+1]) << 8) | (uint32(data[i+2]) << 16) | (uint32(data[i+3]) << 24))

		if o < 4 {
			switch o {
			case 2:
				h1 = uint32(uint32(data[i]) | (uint32(data[i+1]) << 8))
			case 3:
				h1 = uint32(uint32(data[i]) | (uint32(data[i+1]) << 8) | (uint32(data[i+2]) << 16))
			default:
				h1 = uint32(data[i])
			}
		}

		h1 *= uint32(c4)
		h1 = bits.RotateLeft32(h1, 15)
		h1 *= uint32(c3)
		k1 = k1 ^ h1

		if o > 4 {
			var h2 uint32
			switch o {
			case 6:
				h2 = uint32(uint32(data[i+4]) | (uint32(data[i+5]) << 8))
			case 7:
				h2 = uint32(uint32(data[i+4]) | (uint32(data[i+5]) << 8) | (uint32(data[i+6]) << 16))
			default:
				h2 = uint32(data[i+4])
			}
			var neg int = c7
			h3 := h2 * uint32(neg)
			h3 = bits.RotateLeft32(h3, 17)
			h3 *= c4
			k2 ^= h3
		}
	}

	k1 ^= uint32(size)
	k2 ^= uint32(size)
	k1 += k2
	k2 += k1
	k1 ^= (k1 >> 16)
	k1 *= uint32(c1)
	k1 ^= (k1 >> 13)
	k1 *= uint32(c2)
	k1 ^= (k1 >> 16)
	k2 ^= (k2 >> 16)
	k2 *= uint32(c1)
	k2 ^= (k2 >> 13)
	k2 *= uint32(c2)
	k2 ^= (k2 >> 16)
	k1 += k2
	k2 += k1

	return (uint64(k2) << 32) | uint64(k1)
}

func GenerateUniqueString(input ...string) string {
	hash64 := murmurHash64([]byte(strings.Join(input, "-")))
	return base32Encode(hash64)
}
