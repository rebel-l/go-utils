package option

import (
	"fmt"
	"strings"
)

// Options provides slice of Option.
type Options []Option

// IsValidOption checks if Options contains a specific key.
func (o Options) IsValidOption(key string) bool {
	for _, v := range o {
		if v.Key == key {
			return true
		}
	}

	return false
}

// IsValidOptionCI checks if Options contains a specific key as IsValidOption() but works case insensitive.
func (o Options) IsValidOptionCI(key string) bool {
	key = strings.ToLower(key)
	for _, v := range o {
		if strings.ToLower(v.Key) == key {
			return true
		}
	}

	return false
}

// ForAll iterates over all options and executes a callback on each option.
func (o Options) ForAll(callback func(option Option) error) error {
	for i, v := range o {
		err := callback(v)
		if err != nil {
			return fmt.Errorf("failed to execute callback on entry %d: %w", i, err)
		}
	}

	return nil
}
