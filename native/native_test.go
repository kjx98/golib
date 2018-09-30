package native

import (
	"testing"
	"encoding/binary"
	"bytes"
	"unsafe"
)

type A struct {
    // should be exported member when read back from buffer
    One uint32
    Two uint32
    Three uint64
}

//测试获取数据
func TestHostEndian(t *testing.T) {
	//测试获取数据
	switch HostEndian {
	case binary.BigEndian:
		t.Log("Host Endian: Big Endian")
	case binary.LittleEndian:
		t.Log("Host Endian: Little Endian")
	default:
		t.Fatal("Invalid Host Endian")
	}

	//获取数据
	aa := uint32(0x05060708)
	ab := make([] byte, 8)
	ac := make([] byte, 8)
	NativeEndian.PutUint32(ab, aa)
	HostEndian.PutUint32(ac, aa)
	if bytes.Compare(ab, ac) != 0 {
		t.Error("hostEndian != nativeEndian")
	}

	a1 := A{One: 1, Two: 2, Three: 3}
	buf := new(bytes.Buffer)
	if err:= binary.Write(buf,NativeEndian,a1); err != nil {
		t.Error("binary.Write", err)
	} else {
		t.Log("struct write:", buf.Bytes())
	}
	var a2 A
	buf1 := bytes.NewReader(buf.Bytes())
	if err:= binary.Read(buf1, NativeEndian, &a2); err != nil {
		t.Error("binary.Read", err)
	}
	if a1 != a2 {
		t.Error("Write/Read struct diff")
	} else {
		t.Log("binary.Read:", a2)
	}
	var a22 A
	a2l := unsafe.Sizeof(a22)
	if a2l != 16 {
		t.Error("struct size diff: 16 vs", a2l)
	}
	a2p := (* [16]byte)(unsafe.Pointer(&a22))
	copy(a2p[0:], buf.Bytes())
	if a1 != a22 {
		t.Error("Write/copy struct diff")
	} else {
		t.Log("copy struct:", a22)
	}
}


//解析性能测试
func BenchmarkPut64(b *testing.B) {
	var buf [8]byte
	aa := uint64(0x0102030405060708)
	for i := 0; i < b.N; i++ {
		NativeEndian.PutUint64(buf[0:], aa)
		NativeEndian.Uint64(buf[0:])
	}
}

func BenchmarkPut64Host(b *testing.B) {
	var buf [8]byte
	aa := uint64(0x0102030405060708)
	for i := 0; i < b.N; i++ {
		HostEndian.PutUint64(buf[0:], aa)
		HostEndian.Uint64(buf[0:])
	}
}

func BenchmarkStruct(b *testing.B) {
	aa := A{One:11, Two: 22, Three: 33}
	for i := 0; i < b.N; i++ {
		buf := new(bytes.Buffer)
		binary.Write(buf, NativeEndian, aa)
	}
}

func BenchmarkStructMy(b *testing.B) {
	aa := A{One:11, Two: 22, Three: 33}
	var buf	[64]byte
	for i := 0; i < b.N; i++ {
		NativeEndian.PutUint32(buf[0:], uint32(aa.One))
		NativeEndian.PutUint32(buf[4:], uint32(aa.Two))
		NativeEndian.PutUint64(buf[8:], uint64(aa.Three))
	}
}

func BenchmarkStructCopy(b *testing.B) {
	aa := A{One:11, Two: 22, Three: 33}
	var buf	[64]byte
	aap := (* [16]byte)(unsafe.Pointer(&aa))
	for i := 0; i < b.N; i++ {
		copy(buf[0:], aap[0:])
	}
}

func BenchmarkStructHost(b *testing.B) {
	aa := A{One:11, Two: 22, Three: 33}
	for i := 0; i < b.N; i++ {
		buf := new(bytes.Buffer)
		binary.Write(buf, HostEndian, aa)
	}
}

func BenchmarkStructHostMy(b *testing.B) {
	aa := A{One:11, Two: 22, Three: 33}
	var buf	[64]byte
	for i := 0; i < b.N; i++ {
		HostEndian.PutUint32(buf[0:], uint32(aa.One))
		HostEndian.PutUint32(buf[4:], uint32(aa.Two))
		HostEndian.PutUint64(buf[8:], uint64(aa.Three))
	}
}
