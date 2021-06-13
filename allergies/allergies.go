package allergies

// Allergies lists all of the things that a person is allergic to,
// in order of increasing score for that allergen
func Allergies(k score) []string {
	allergies := make([]string, 0, 2)

	for key := score(1); key < _maxAllergen*2; key = key << 1 {
		if k&key > 0 {
			allergies = append(allergies, allergens[key])
		}
	}

	return allergies
}

// AllergicTo determines if the person is allergic to the given substance
func AllergicTo(k score, substance string) bool {
	return k&allergensInv[substance] > 0
}

type score = uint

const (
	eggs score = 1 << iota
	peanuts
	shellfish
	strawberries
	tomatoes
	chocolate
	pollen
	cats
	_maxAllergen = cats
)

var allergens = map[score]string{
	eggs:         "eggs",
	peanuts:      "peanuts",
	shellfish:    "shellfish",
	strawberries: "strawberries",
	tomatoes:     "tomatoes",
	chocolate:    "chocolate",
	pollen:       "pollen",
	cats:         "cats",
}

var allergensInv = map[string]score{
	"eggs":         eggs,
	"peanuts":      peanuts,
	"shellfish":    shellfish,
	"strawberries": strawberries,
	"tomatoes":     tomatoes,
	"chocolate":    chocolate,
	"pollen":       pollen,
	"cats":         cats,
}
