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

	entriesCopy, err := cloneAndParseDate(entries)
	if err != nil {
		return "", err
	}
	sortEntries(entriesCopy)

	ledger := New(currencyOpt, localeOpt)
	var b strings.Builder

	ledger.WriteHeader(&b)
	for _, e := range entriesCopy {
		ledger.WriteEntry(&b, e)
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
		dateLayout:     "2006-01-02",
		currencySymbol: " $",
		symbolPosition: afterSuffix,
		negPrefix:      "(",
		negSuffix:      ")",
		posPrefix:      " ",
		posSuffix:      " ",
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

func cloneAndParseDate(entries []Entry) ([]Entry, error) {
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

// WriteHeader writes the header for this ledger to the given writer,
// with a trailing newline character.
// Language-specific terminology can be provided when creating a new
// Ledger by passing in either a WithTerms() Option, or a WithLocale() Option.
// If multiple such Options are passed in, then the last one is used.
func (j *Ledger) WriteHeader(w io.Writer) error {
	format := "%-10s | %-" + strconv.Itoa(j.maxDescription) + "s | %s\n"
	date, ok := j.terms["Date"]
	if !ok {
		date = "Date"
	}

	desc, ok := j.terms["Description"]
	if !ok {
		desc = "Description"
	}

	amount, ok := j.terms["Amount"]
	if !ok {
		amount = "Amount"
	}

	_, err := fmt.Fprintf(w, format, date, desc, amount)
	return err
}

// WriteEntry writes the given entry to the writer, using this Ledger's format
func (j *Ledger) WriteEntry(w io.Writer, e Entry) error {
	date := e.strongDate.Format(j.dateLayout)
	desc := j.formatDesc(e.Description)
	change := j.formatChange(e.Change)
	_, err := fmt.Fprintf(w, "%s | %s | %13s\n", date, desc, change)
	return err
}

func (j *Ledger) formatDesc(desc string) string {
	maxLen := j.maxDescription
	if len(desc) > maxLen {
		return desc[:maxLen-3] + "..."
	}
	return fmt.Sprintf("%-"+strconv.Itoa(maxLen)+"s", desc)
}

func (j *Ledger) formatChange(cents int) string {
	var b strings.Builder
	negative := cents < 0
	if negative {
		cents = cents * -1
	}

	if j.symbolPosition == beforePrefix {
		b.WriteString(j.currencySymbol)
	}

	if negative {
		b.WriteString(j.negPrefix)
	} else {
		b.WriteString(j.posPrefix)
	}

	if j.symbolPosition == afterPrefix {
		b.WriteString(j.currencySymbol)
	}

	whole, cents := cents/100, cents%100
	b.WriteString(formatThousands(whole, j.thousands))
	b.WriteString(j.decimal)
	fmt.Fprintf(&b, "%02d", cents)

	if j.symbolPosition == beforeSuffix {
		b.WriteString(j.currencySymbol)
	}

	if negative {
		b.WriteString(j.negSuffix)
	} else {
		b.WriteString(j.posSuffix)
	}

	if j.symbolPosition == afterSuffix {
		b.WriteString(j.currencySymbol)
	}

	return b.String()
}

// formatThousands separates the the given integer using a thousands separator
func formatThousands(n int, sep string) string {
	rest := strconv.Itoa(n)
	if len(sep) == 0 {
		return rest
	}

	var b strings.Builder
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
		b.WriteString(sep)
	}
	b.WriteString(parts[0])
	return b.String()
}
