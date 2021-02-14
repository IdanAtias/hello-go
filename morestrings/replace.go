package morestrings

// replace runes in a string
func ReplaceRunes(s string, oldRune, newRune rune) string {
	r := []rune(s)
	for i := 0; i < len(s); i++ {
		if r[i] == oldRune {
			r[i] = newRune
		}
	}
	return string(r)
}
