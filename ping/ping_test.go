// Copyright 2009 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// taken from http://golang.org/src/pkg/net/ipraw_test.go

package ping

import "testing"

func TestPing(t *testing.T) {
	type args struct {
		address string
		count   int
	}
	tests := []struct {
		name    string
		args    args
		wantCnt int
		wantErr bool
	}{
		// TODO: Add test cases.
		{"testlocal", args{"127.0.0.1", 5}, 5, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCnt, gotLat, err := Ping(tt.args.address, tt.args.count)
			if (err != nil) != tt.wantErr {
				t.Errorf("Ping() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotCnt != tt.wantCnt {
				t.Errorf("Ping() gotCnt = %v, want %v", gotCnt, tt.wantCnt)
				return
			}
			t.Log("Ping latency:", float64(gotLat)*1e-6/float64(gotCnt), " us")
		})
	}
}
