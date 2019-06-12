package trick

import "testing"

func TestIs64(t *testing.T) {
	Is64()
}

func TestIsLittleEndian(t *testing.T){
	t.Log(IsLittleEndian()) // darwin littleEndian
}