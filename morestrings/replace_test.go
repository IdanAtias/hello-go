package morestrings

import "testing"

func TestReplaceRunes(t *testing.T) {
	cases := []struct {
		in                 string
		old_rune, new_rune rune
		out                string
	}{
		{"hello go!", 'l', '1', "he11o go!"},
		{"hello go!", 'x', 'y', "hello go!"},
		{"", 'x', 'y', ""},
		{"xxx", 'x', 'y', "yyy"},
		{"XXX", 'x', 'y', "XXX"},
		{"Abc", 'A', '4', "4bc"},
	}
	for _, c := range cases {
		out := ReplaceRunes(c.in, c.old_rune, c.new_rune)
		if out != c.out {
			t.Errorf("ReplaceRunes(%q, %q, %q) == %q, expected %q", c.in, c.old_rune, c.new_rune, out, c.out)
		}
	}
}
