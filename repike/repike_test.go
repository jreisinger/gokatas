package repike

import "testing"

type test struct {
	name    string
	re      string
	text    string
	matched bool
}

var tests = []test{
	{"EmptyBoth", "", "", true},
	{"EmptyRegex", "", "foo", true},
	{"EmptyText", "foo", "", false},

	{"LiteralExact", "foo", "foo", true},
	{"LiteralSearch", "foo", "A food truck", true},
	{"LiteralNo", "foo", "foes", false},

	{"StartOnlyEmpty", "^", "", true},
	{"StartOnlySome", "^", "foo", true},
	{"StartExact", "^foo", "foo", true},
	{"StartInexact", "^foo", "food", true},
	{"StartIncomplete", "^foo", "fo", false},
	{"StartNo1", "^foo", "xfoo", false},
	{"StartNo2", "^foo", "A food truck", false},
	{"StartEndEmpty", "^$", "", true},
	{"StartEndSome", "^$", "x", false},
	{"StartDotEndZero", "^.$", "", false},
	{"StartDotEndOne", "^.$", "x", true},
	{"StartDotEndTwo", "^.$", "xy", false},

	{"EndOnlyEmpty", "$", "", true},
	{"EndOnlySome", "$", "foo", true},
	{"EndExact", "foo$", "foo", true},
	{"EndInexact", "foo$", "xfoo", true},
	{"EndIncomplete", "foo$", "oo", false},
	{"EndNo1", "foo$", "food", false},
	{"EndNo2", "foo$", "A food truck", false},
	{"EndAB", "ab$", "abcab", true},

	{"CStarOnlyZero", "a*", "", true},
	{"CStarOnlyOne", "a*", "a", true},
	{"CStarOnlyMany", "a*", "aaaa", true},
	{"CStarZero", "fo*d", "fd", true},
	{"CStarOne", "fo*d", "fod", true},
	{"CStarTwo", "fo*d", "food", true},
	{"CStarMany", "fo*d", "fooooood", true},
	{"CStarSearch", "fo*d", "A food truck", true},
	{"CStarNoEmpty", "fo*d", "", false},
	{"CStarNo1", "fo*d", "f", false},
	{"CStarNo2", "fo*d", "fx", false},
	{"CStarNo3", "fo*d", "fox", false},

	{"DotStarZero", "foo.*bar", "foobar", true},
	{"DotStarOne", "foo.*bar", "foodbar", true},
	{"DotStarMulti", "foo.*bar", "food and bar", true},
	{"DotStarSearch", "foo.*bar", "The food bar.", true},
	{"DotStarNoEmpty", "foo.*bar", "", false},
	{"DotStarNo1", "foo.*bar", "foo", false},
	{"DotStarNo2", "foo.*bar", "bar", false},
	{"DotStarNo3", "foo.*bar", "fooar", false},
	{"DotStarNo4", "foo.*bar", "fobar", false},

	{"DotStarSuffixZero", "foo.*", "foo", true},
	{"DotStarSuffixSearch", "foo.*", "A food truck", true},
	{"DotStarSuffixAnchored", "^foo.*$", "foodie", true},
	{"DotStarSuffixAnchoredNo", "^foo.*$", "A food truck", false},

	{"DotStarPrefixZero", ".*foo", "foo", true},
	{"DotStarPrefixSearch", ".*foo", "A food truck", true},
	{"DotStarPrefixAnchored", "^.*foo$", "A foo", true},
	{"DotStarPrefixAnchoredNo", "^.*foo$", "A food truck", false},

	{"DotStarOnlyZero", ".*", "foo", true},
	{"DotStarOnlySearch", ".*", "A food truck", true},
	{"DotStarOnlyAnchored", "^.*$", "A foo", true},

	{"DotDotExact", "abc..", "abcde", true},
	{"DotDotSmall", "abc..", "abcd", false},
	{"DotDotLarge", "abc..", "abcdefghijklm", true},
	{"DotDotSearch", "abc..", "_abcde_", true},
	{"DotDotNo", "abc..", "vwxyz", false},

	{"DotSameLengthA", "f.o", "fao", true},
	{"DotSameLengthZ", "f.o", "fzo", true},
	{"DotSameLengthDot", "f.o", "f.o", true},
	{"DotSearch", "f.o", "A fxod truck", true},
	{"DotIncomplete", "f.o", "fo", false},
	{"DotNo", "f.o", "fxy", false},
	{"DotStartYes", ".dog", "The dog", true},
	{"DotStartMin", ".dog", "_dog", true},
	{"DotStartNo", ".dog", "doggy", false},

	{"DotOnlyEmpty", ".", "", false},
	{"DotOnlyOne", ".", "a", true},
	{"DotOnlyMany", ".", "abcdef", true},
	{"DotOnlyAnchored", "^.$", "a", true},
	{"DotOnlyAnchoredNo", "^.$", "ab", false},

	{"DotCExact", ".a", "xa", true},
	{"DotCMore", ".a", "_ya_", true},
	{"DotCNo", ".a", "xb", false},

	{"CDotExact", "a.", "ax", true},
	{"CDotMore", "a.", "_ay_", true},
	{"CDotNo", "a.", "bx", false},
}

func TestMatch(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name+"/repike", func(t *testing.T) {
			matched := Match(test.re, test.text)
			if matched != test.matched {
				t.Fatalf("got %v, want %v", matched, test.matched)
			}
		})
	}
}
