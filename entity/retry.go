package entity

// Retry represents tag for retry settings.
type Retry struct {
	// define later
}

// IsForLogger always return false.
func (r Retry) IsForLogger() bool {
	return false
}

// IsForTracer always return false.
func (r Retry) IsForTracer() bool {
	return false
}

// IsForRetrier always return false.
func (r Retry) IsForRetrier() bool {
	return true
}
