// MurmurHash3 was written by Austin Appleby, and is placed in the public
package murmurhash3

// #include <MurmurHash3.h>
import "C"
import "unsafe"

func Murmur3A(key []byte, seed uint32) uint32 {
	var p *byte
	if len(key) > 0 {
		p = &key[0]
	}
	var hash uint32
	C.MurmurHash3_x86_32(unsafe.Pointer(p), C.int(len(key)), C.uint32_t(seed), unsafe.Pointer(&hash))
	return hash
}

func Murmur3C(key []byte, seed uint32) [4]uint32 {
	var p *byte
	if len(key) > 0 {
		p = &key[0]
	}
	var hash [4]uint32
	C.MurmurHash3_x86_128(unsafe.Pointer(p), C.int(len(key)), C.uint32_t(seed), unsafe.Pointer(&hash))
	return hash
}

func Murmur3F(key []byte, seed uint32) [2]uint64 {
	var p *byte
	if len(key) > 0 {
		p = &key[0]
	}
	var hash [2]uint64
	C.MurmurHash3_x64_128(unsafe.Pointer(p), C.int(len(key)), C.uint32_t(seed), unsafe.Pointer(&hash))
	return hash
}
