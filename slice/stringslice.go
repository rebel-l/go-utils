package slice

// StringSlice represents a slice of strings
type StringSlice []string

// IsIn searches for a given value in the slice. If search matches a value it returns true, otherwise false
func (s StringSlice) IsIn(search string) bool {
	for _, v := range s {
		if v == search {
			return true
		}
	}

	return false
}

// IsNotIn searches for a given value in the slice. If search doesn't match a value it returns true, otherwise false
func (s StringSlice) IsNotIn(search string) bool {
	return !s.IsIn(search)
}

// Diff returns a StringSlice containing all values which are not in the given StringSlice
func (s StringSlice) Diff(b StringSlice) StringSlice {
	var result StringSlice
	for _, a := range s { // nolint: wsl
		if b.IsNotIn(a) {
			result = append(result, a)
		}
	}

	return result
}
