package ledger

import (
	"errors"
)

type options struct {
	dateLayout           string            // used by (t time.Time) Format()
	currencySymbol       string            // the currency symbol
	symbolPosition       symbolPosition    // where to position the symbol
	negPrefix, negSuffix string            // prefixes and suffixes for positive
	posPrefix, posSuffix string            // and negative monetary amounts
	thousands            string            // thousands separator
	decimal              string            // decimal point separator
	terms                map[string]string // override the default header names
	maxDescription       int               // max length of the description
}

// Option overrides the default formatting behaviour of the Ledger.
// To obtain Options for use, call the various With...() functions.
type Option interface {
	apply(o *options)
}

// WithCurrency provides the option to override the currency symbol
func WithCurrency(c string) (Option, error) {
	switch c {
	case "EUR":
		return currencyOpt("€"), nil
	case "USD":
		return currencyOpt("$"), nil
	default:
		return nil, errors.New("invalid currency")
	}
}

// WithLocale provides a bundle of options to override ledger formatting,
// including the heading terminology, the date format, and monetary format
func WithLocale(l string) (Option, error) {
	switch l {
	case "nl-NL":
		return dutchOpt, nil
	case "en-US":
		return usaOpt, nil
	default:
		return nil, errors.New("invalid locale")
	}
}

// compile-time interface checks
var (
	_ Option = currencyOpt("$")
	_ Option = localeOpt{}
	_ Option = maxDescription(0)
)

type currencyOpt string

func (c currencyOpt) apply(o *options) {
	o.currencySymbol = string(c)
}

var (
	dutchOpt localeOpt = localeOpt{
		dateLayout:     "02-01-2006",
		posPrefix:      " ",
		posSuffix:      " ",
		negPrefix:      " ",
		negSuffix:      "-",
		thousands:      ".",
		decimal:        ",",
		symbolPosition: beforePrefix,
		terms: map[string]string{
			"Date":        "Datum",
			"Description": "Omschrijving",
			"Amount":      "Verandering",
		},
	}

	usaOpt localeOpt = localeOpt{
		dateLayout:     "01/02/2006",
		posPrefix:      " ",
		posSuffix:      " ",
		negPrefix:      "(",
		negSuffix:      ")",
		thousands:      ",",
		decimal:        ".",
		symbolPosition: afterPrefix,
		terms: map[string]string{
			"Amount": "Change",
		},
	}
)

type symbolPosition uint

const (
	beforePrefix symbolPosition = iota
	afterPrefix
	beforeSuffix
	afterSuffix
)

type localeOpt struct {
	dateLayout           string
	negPrefix, negSuffix string
	posPrefix, posSuffix string
	symbolPosition       symbolPosition
	thousands            string
	decimal              string
	terms                map[string]string
}

func (l localeOpt) apply(o *options) {
	o.dateLayout = l.dateLayout
	o.negPrefix, o.negSuffix = l.negPrefix, l.negSuffix
	o.posPrefix, o.posSuffix = l.posPrefix, l.posSuffix
	o.symbolPosition = l.symbolPosition
	o.thousands, o.decimal = l.thousands, l.decimal
	for k, v := range l.terms {
		o.terms[k] = v
	}
}

type maxDescription int

func (max maxDescription) apply(o *options) {
	o.maxDescription = int(max)
}
