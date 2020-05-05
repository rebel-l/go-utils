package mapof

import (
	"github.com/rebel-l/go-utils/slice"
)

// StringSliceMap represents a map of string slices.
type StringSliceMap map[string]slice.StringSlice

// AddUniqueValue adds a value to the slice of strings by a given key.
// It ensures that the value is added only once to the slice.
func (s StringSliceMap) AddUniqueValue(key, value string) {
	values, ok := s[key]
	if ok {
		if values.IsNotIn(value) {
			s[key] = append(values, value)
		}
	} else {
		s[key] = []string{value}
	}
}

// KeyExists checks if a given key exists in the map.
func (s StringSliceMap) KeyExists(key string) bool {
	_, ok := s[key]
	return ok
}

// KeyNotExists checks if a given key doesn't exits in the map.
func (s StringSliceMap) KeyNotExists(key string) bool {
	return !s.KeyExists(key)
}

// GetValuesForKey returns the StringSlice of a given key. If key doesn't exist, it returns an empty StringSlice.
func (s StringSliceMap) GetValuesForKey(key string) slice.StringSlice {
	values, ok := s[key]
	if ok {
		return values
	}

	return slice.StringSlice{}
}
