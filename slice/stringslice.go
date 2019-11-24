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

// Len returns the length of the slice
func (s StringSlice) Len() int {
	return len(s)
}

// IsEqual checks if two slices contains the same values. It don't cares about the order.
// Note: IsSame() is faster and is preferred as long as it is ensured the order of the values is same in both slices.
func (s StringSlice) IsEqual(b StringSlice) bool {
	if s.Len() != b.Len() {
		return false
	}

	for _, item := range s {
		if b.IsNotIn(item) {
			return false
		}
	}

	return true
}

// IsSame checks if two slices contains the same values in exactly the same order.
func (s StringSlice) IsSame(b StringSlice) bool {
	if s.Len() != b.Len() {
		return false
	}

	for k, item := range s {
		if item != b[k] {
			return false
		}
	}

	return true
}
