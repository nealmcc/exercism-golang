package chromatic

type interval = rune

// each interval determines how many notes to go up by
var sizes = map[interval]int{
	'm': 1,
	'M': 2,
	'A': 3,
}
