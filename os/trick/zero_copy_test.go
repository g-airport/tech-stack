package trick

import (
	"bytes"
	"reflect"
	"testing"
)

func TestBytes2Data(t *testing.T) {
	buf := new(bytes.Buffer)
	buf.Write(Data2Bytes(&Data{Int32: 32, String: "string"}))
	d := Bytes2Data(buf.Bytes())
	t.Log(d)
}

func TestData2Bytes(t *testing.T) {
	b := Data2Bytes(&Data{Int32: 32, String: "string"})
	t.Log(b)
}

func TestBytes2StringFast(t *testing.T) {
	s := Bytes2StringFast([]byte("string"))
	t.Log(s)

	b := String2ByteFast("string")
	buf := new(bytes.Buffer)
	buf.Write(b)
	t.Log(reflect.DeepEqual(buf.Bytes(), b))
}

func TestString2ByteFast(t *testing.T) {
	b := String2ByteFast("string")
	t.Log(b)
}

func BenchmarkBytes2StringFast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Bytes2StringFast([]byte("string"))
	}

}

func BenchmarkString2ByteFast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = String2ByteFast("string")
	}
}

// use bigger string to test
func BenchmarkString2ByteFastV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = String2ByteFastV2("string")
	}
}
