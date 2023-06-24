package areader_test

import (
	"bytes"
	"testing"

	"github.com/jreisinger/gokatas/areader"
)

func TestA_Read(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name string
		p    []byte
		want []byte
	}{
		{
			name: "empty slice",
			p:    []byte{},
			want: []byte{},
		},
		{
			name: "short slice",
			p:    []byte{0, 0, 0},
			want: []byte{'A', 'A', 'A'},
		},
		{
			name: "long slice",
			p:    make([]byte, 100),
			want: bytes.Repeat([]byte{'A'}, 100),
		},
		{
			name: "nil slice",
			p:    nil,
			want: []byte{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			a := areader.A{}
			n, err := a.Read(tc.p)
			if err != nil {
				t.Errorf("A.Read() error = %v", err)
			}
			if n != len(tc.want) {
				t.Errorf("A.Read() n = %v, want %v", n, len(tc.want))
			}
			if !bytes.Equal(tc.p, tc.want) {
				t.Errorf("A.Read() p = %v, want %v", tc.p, tc.want)
			}
		})
	}
}
