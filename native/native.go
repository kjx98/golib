// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package binary implements simple translation between numbers and byte
// sequences and encoding and decoding of varints.
//
// Numbers are translated by reading and writing fixed-size values.
// A fixed-size value is either a fixed-size arithmetic
// type (bool, int8, uint8, int16, float32, complex64, ...)
// or an array or struct containing only fixed-size values.
//
// The varint functions encode and decode single integer values using
// a variable-length encoding; smaller values require fewer bytes.
// For a specification, see
// https://developers.google.com/protocol-buffers/docs/encoding.
//
// This package favors simplicity over efficiency. Clients that require
// high-performance serialization, especially for large data structures,
// should look at more advanced solutions such as the encoding/gob
// package or protocol buffers.

// +build !386

package native

import (
	"unsafe"
	"encoding/binary"
)

// NativeEndian is the native-endian implementation of ByteOrder.
var NativeEndian nativeEndian

type nativeEndian struct{}

var HostEndian binary.ByteOrder

func init() {
	bb := []byte{1,2,3,4}
	hl := NativeEndian.Uint32(bb)
	switch hl {
	case uint32(0x01020304):
		HostEndian = binary.BigEndian
	case uint32(0x04030201):
		HostEndian = binary.LittleEndian
	}
}

func (nativeEndian) Uint16(b []byte) (res uint16) {
	_ = b[1]
	bp := (* [2]byte)(unsafe.Pointer(&res))
	copy(bp[:2], b)
	return
}

func (nativeEndian) PutUint16(b []byte, v uint16) {
	_ = b[1]
	bp := (* [2]byte)(unsafe.Pointer(&v))
	copy(b, bp[:2])
}

func (nativeEndian) Uint32(b []byte) (res uint32) {
	_ = b[3]
	bp := (* [4]byte)(unsafe.Pointer(&res))
	copy(bp[:4], b)
	return
}

func (nativeEndian) Float32(b []byte) (res float32) {
	_ = b[3]
	bp := (* [4]byte)(unsafe.Pointer(&res))
	copy(bp[:4], b)
	return
}

func (nativeEndian) PutUint32(b []byte, v uint32) {
	_ = b[3]
	bp := (* [4]byte)(unsafe.Pointer(&v))
	copy(b, bp[:4])
}

func (nativeEndian) PutFloat32(b []byte, v float32) {
	_ = b[3]
	bp := (* [4]byte)(unsafe.Pointer(&v))
	copy(b, bp[:4])
}

func (nativeEndian) Uint64(b []byte) (res uint64) {
	_ = b[7]
	bp := (* [8]byte)(unsafe.Pointer(&res))
	copy(bp[:8], b)
	return
}

func (nativeEndian) Float64(b []byte) (res float64) {
	_ = b[7]
	bp := (* [8]byte)(unsafe.Pointer(&res))
	copy(bp[:8], b)
	return
}

func (nativeEndian) PutUint64(b []byte, v uint64) {
	_ = b[7]
	bp := (* [8]byte)(unsafe.Pointer(&v))
	copy(b, bp[:8])
}

func (nativeEndian) PutFloat64(b []byte, v float64) {
	_ = b[7]
	bp := (* [8]byte)(unsafe.Pointer(&v))
	copy(b, bp[:8])
}

func (nativeEndian) String() string { return "NativeEndian" }

func (nativeEndian) GoString() string { return "binary.NativeEndian" }
