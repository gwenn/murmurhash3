package murmurhash3

import "testing"

func TestMurmur3A(t *testing.T) {
	expected := uint32(3127628307)
  hash := Murmur3A([]byte("test"), 0)
	if hash != expected {
		t.Errorf("Expected %d but was %d for Murmur3A\n", expected, hash)
	}
}

func TestMurmur3C(t *testing.T) {
	expected := []uint32{1862463280, 1426881896, 1426881896, 1426881896}
  hash := Murmur3C([]byte("test"), 0)
	for i, e := range expected {
		if hash[i] != e {
			t.Errorf("Expected %d but was %d for Murmur3C[%d]\n", e, hash[i], i)
		}
	}
}

func TestMurmur3F(t *testing.T) {
	expected := []uint64{12429135405209477533, 11102079182576635266}
  hash := Murmur3F([]byte("test"), 0)
	for i, e := range expected {
		if hash[i] != e {
			t.Errorf("Expected %d but was %d for Murmur3F[%d]\n", e, hash[i], i)
		}
	}
}

//func BenchmarkMurmur3A(b *testing.B) {
//}
