package cryptopals

import (
	"testing"
)

func TestHexToBase64(t *testing.T) {
	h := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	want := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	got, err := HexToBase64(h)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestFixedXOR(t *testing.T) {
	in := "1c0111001f010100061a024b53535009181c"
	against := "686974207468652062756c6c277320657965"
	want := "746865206b696420646f6e277420706c6179"
	got, err := FixedXOR(in, against)
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Errorf("want %q, got %q", want, got)
	}
}
