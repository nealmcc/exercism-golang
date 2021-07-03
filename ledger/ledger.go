package ledger

import (
	"errors"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Entry is a transaction in the ledger
type Entry struct {
	Date        string // "YYYY-mm-dd"
	strongDate  time.Time
	Description string
	Change      int // in cents
}

// FormatLedger returns the list of entries formatted according to the given
// currency and locale, with descriptions set to a fixed length
func FormatLedger(currency, locale string, entries []Entry) (string, error) {
	currencyOpt, err := WithCurrency(currency)
	if err != nil {
		return "", err
	}

	localeOpt, err := WithLocale(locale)
	if err != nil {
		return "", err
	}

	entriesCopy, err := cloneAndParse(entries)
	if err != nil {
		return "", err
	}
	sortEntries(entriesCopy)

	ledger := New(currencyOpt, localeOpt)
	var b strings.Builder

	ledger.WriteHeader(&b)
	for _, e := range entriesCopy {
		ledger.WriteEntry(&b, e, locale, currency)
	}
	return b.String(), nil
}

// Ledger is a formatting tool which can write a heading
// and sequence of Entries to a given io.Writer.
// The formatting of a Ledger is controlled by Options.
type Ledger struct {
	options
}

// New creates a new Ledger which will use the given formatting options
func New(opts ...Option) *Ledger {
	// defaults:
	options := options{
		currencySymbol: "$",
		dateLayout:     "2006-01-02",
		negPrefix:      "(",
		negSuffix:      ")",
		posPrefix:      "",
		posSuffix:      "",
		thousands:      ",",
		decimal:        ".",
		terms: map[string]string{
			"Date":        "Date",
			"Description": "Description",
			"Amount":      "Amount",
		},
		maxDescription: 25,
	}

	for _, o := range opts {
		o.apply(&options)
	}

	return &Ledger{options}
}

// WriteHeader writes the header for this ledger to the given writer,
// with a trailing newline character.
// Language-specific terminology can be provided when creating a new
// Ledger by passing in either a WithTerms() Option, or a WithLocale() Option.
// If multiple such Options are passed in, then the last one is used.
func (l *Ledger) WriteHeader(w io.Writer) error {
	format := "%-10s | %-" + strconv.Itoa(l.maxDescription) + "s | %s\n"
	date, ok := l.terms["Date"]
	if !ok {
		date = "Date"
	}

	desc, ok := l.terms["Description"]
	if !ok {
		desc = "Description"
	}

	amount, ok := l.terms["Amount"]
	if !ok {
		amount = "Amount"
	}

	_, err := fmt.Fprintf(w, format, date, desc, amount)
	return err
}

func (l *Ledger) WriteEntry(w io.Writer, e Entry, locale, currency string) error {
	date := e.strongDate.Format(l.dateLayout)
	desc := formatDesc(e.Description, 25 /* maxLen */)
	amount := formatAmount(locale, currency, e.Change)
	_, err := fmt.Fprintf(w, "%-10s | %-25s | %13s\n", date, desc, amount)
	return err
}

func formatDesc(desc string, maxLen int) string {
	if len(desc) > maxLen {
		return desc[:maxLen-3] + "..."
	}
	return fmt.Sprintf("%-"+strconv.Itoa(maxLen)+"s", desc)
}

func formatAmount(locale, currency string, cents int) string {
	switch locale {
	case "nl-NL":
		return amountDutch(cents, currency)
	default:
		return amountUSA(cents, currency)
	}
}

func amountDutch(cents int, currency string) string {
	var b strings.Builder
	negative := cents < 0
	if negative {
		cents = cents * -1
	}
	writeCurrency(&b, currency)
	b.WriteByte(' ')
	whole, cents := cents/100, cents%100
	fmt.Fprintf(&b, "%s,%02d", formatThousands(whole, '.'), cents)
	if negative {
		b.WriteByte('-')
	} else {
		b.WriteByte(' ')
	}
	return b.String()
}

func amountUSA(cents int, currency string) string {
	var b strings.Builder
	negative := cents < 0
	if negative {
		cents = cents * -1
	}
	if negative {
		b.WriteByte('(')
	}
	writeCurrency(&b, currency)
	whole, cents := cents/100, cents%100
	fmt.Fprintf(&b, "%s.%02d", formatThousands(whole, ','), cents)
	if negative {
		b.WriteByte(')')
	} else {
		b.WriteByte(' ')
	}
	return b.String()
}

func writeCurrency(b *strings.Builder, currency string) {
	if currency == "EUR" {
		b.WriteRune('€')
	} else {
		b.WriteByte('$')
	}
}

// formatThousands separates the the given integer using a thousands separator
func formatThousands(n int, sep byte) string {
	var b strings.Builder
	rest := strconv.Itoa(n)
	parts := make([]string, 0, len(rest)/3+1)
	for len(rest) > 3 {
		parts = append(parts, rest[len(rest)-3:])
		rest = rest[:len(rest)-3]
	}
	if len(rest) > 0 {
		parts = append(parts, rest)
	}
	for i := len(parts) - 1; i >= 1; i-- {
		b.WriteString(parts[i])
		b.WriteByte(sep)
	}
	b.WriteString(parts[0])
	return b.String()
}

func cloneAndParse(entries []Entry) ([]Entry, error) {
	r := make([]Entry, len(entries))
	for i, e := range entries {
		date, err := time.Parse("2006-01-02", e.Date)
		if err != nil {
			return nil, errors.New("malformed date")
		}
		r[i] = Entry{
			strongDate:  date,
			Description: e.Description,
			Change:      e.Change,
		}
	}
	return r, nil
}

// sortEntries orders the entries by date, description, cents.
func sortEntries(entries []Entry) {
	sort.Slice(entries, func(i, j int) bool {
		if entries[i].strongDate.Before(entries[j].strongDate) {
			return true
		}
		if entries[i].strongDate.After(entries[j].strongDate) {
			return false
		}
		if entries[i].Description < entries[j].Description {
			return true
		}
		if entries[i].Description > entries[j].Description {
			return false
		}
		return entries[i].Change < entries[j].Change
	})
}
