package ledger

import (
	"errors"
)

type options struct {
	currencySymbol       string
	dateLayout           string
	negPrefix, negSuffix string
	posPrefix, posSuffix string
	thousands            string
	decimal              string
	terms                map[string]string
	maxDescription       int
}

type translateFn = func(string) string

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
		return nil, errors.New("invalid currency requested")
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
		return nil, errors.New("invalid locale requested")
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
		dateLayout: "02-01-2006",
		terms: map[string]string{
			"Date":        "Datum",
			"Description": "Omschrijving",
			"Amount":      "Verandering",
		},
	}

	usaOpt localeOpt = localeOpt{
		dateLayout: "01/02/2006",
		terms: map[string]string{
			"Amount": "Change",
		},
	}
)

type localeOpt struct {
	// the date format as used by (t time.Time) Format()
	dateLayout string
	// prefixes and suffixes for positive and negative monetary amounts
	negPrefix, negSuffix string
	posPrefix, posSuffix string
	// thousands separator
	thousands string
	// decimal point separator
	decimal string
	// terms can selectively override the default header names
	terms map[string]string
}

func (l localeOpt) apply(o *options) {
	o.dateLayout = l.dateLayout
	for k, v := range l.terms {
		o.terms[k] = v
	}
}

type maxDescription int

func (max maxDescription) apply(o *options) {
	o.maxDescription = int(max)
}
