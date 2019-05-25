package option

// Options provides slice of Option
type Options []Option

// IsValidOption checks if Options contains a specific key
func (o Options) IsValidOption(key string) bool {
	for _, v := range o {
		if v.Key == key {
			return true
		}
	}
	return false
}
