package secret

type signal = uint

const (
	wink signal = 1 << iota
	doubleBlink
	closeEyes
	jump
	reverse
)

// Handshake converts a set of signals (expressed as a bitwise union)
// into the corresponding string values
func Handshake(flags signal) []string {
	h := make([]string, 0, 4)

	if flags&wink > 0 {
		h = append(h, "wink")
	}
	if flags&doubleBlink > 0 {
		h = append(h, "double blink")
	}
	if flags&closeEyes > 0 {
		h = append(h, "close your eyes")
	}
	if flags&jump > 0 {
		h = append(h, "jump")
	}
	if flags&reverse > 0 {
		for i, j := 0, len(h)-1; i < j; i, j = i+1, j-1 {
			h[i], h[j] = h[j], h[i]
		}
	}

	return h
}
