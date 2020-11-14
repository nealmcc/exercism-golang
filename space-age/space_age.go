// Package space holds exercise 4 on the Exercism golang track.
package space

// Planet is the name of a planet in Earth's solar system
type Planet string

const earthYear = 31557600.0

var planetaryYear = map[Planet]float64{
	"Mercury": 0.2408467 * earthYear,
	"Venus":   0.61519726 * earthYear,
	"Earth":   1.0 * earthYear,
	"Mars":    1.8808158 * earthYear,
	"Jupiter": 11.862615 * earthYear,
	"Saturn":  29.447498 * earthYear,
	"Uranus":  84.016846 * earthYear,
	"Neptune": 164.79132 * earthYear,
}

// Age is how many times the planet has orbited the sun after s seconds.
func Age(s float64, p Planet) float64 {
	return s / planetaryYear[p]
}
