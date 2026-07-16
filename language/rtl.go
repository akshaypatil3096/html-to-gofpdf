package language

// IsRTL returns true if the direction is RTL.
func IsRTL(dir TextDirection) bool {
	return dir == DirectionRTL
}
