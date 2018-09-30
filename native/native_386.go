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
package native

import "unsafe"
import "encoding/binary"

// NativeEndian is the native-endian implementation of ByteOrder.
var NativeEndian nativeEndian

type nativeEndian struct{}

var HostEndian binary.ByteOrder

func init() {
	HostEndian = binary.LittleEndian
}


func (nativeEndian) Uint16(b []byte) uint16 {
	_ = b[1]
	return *(* uint16)(unsafe.Pointer(&b[0]))
}

func (nativeEndian) PutUint16(b []byte, v uint16) {
	_ = b[1]
	*(* uint16)(unsafe.Pointer(&b[0])) = v
}

func (nativeEndian) Uint32(b []byte) uint32 {
	_ = b[3]
	return *(* uint32)(unsafe.Pointer(&b[0]))
}

func (nativeEndian) PutUint32(b []byte, v uint32) {
	_ = b[3]
	*(* uint32)(unsafe.Pointer(&b[0])) = v
}

func (nativeEndian) Uint64(b []byte) uint64 {
	_ = b[7]
	return *(* uint64)(unsafe.Pointer(&b[0]))
}

func (nativeEndian) PutUint64(b []byte, v uint64) {
	_ = b[7]
	*(* uint64)(unsafe.Pointer(&b[0])) = v
}

func (nativeEndian) String() string { return "NativeEndian" }

func (nativeEndian) GoString() string { return "binary.NativeEndian" }
