package stringutil

// Reverse takes string s, converts it to array of runes, 
// reverses the runes in it, then converts array of runes to string and returns
// it.

func Reverse(s string) string {

	r := []rune(s);	// convet string to array of runes
	for i, j := 0, len(r) - 1; i < len(r)/2; i, j = i + 1, j - 1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r);
}
