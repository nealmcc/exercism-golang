package ledger

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Entry is a transaction in the ledger
type Entry struct {
	Date        string // "YYYY-mm-dd"
	Description string
	Change      int // in cents
}

// FormatLedger returns the list of entries formatted according to the given
// currency and locale, with descriptions set to a fixed length
func FormatLedger(currency, locale string, entries []Entry) (string, error) {
	if currency != "EUR" && currency != "USD" {
		return "", errors.New("invalid currency requested")
	}

	if locale != "nl-NL" && locale != "en-US" {
		return "", errors.New("invalid locale requested")
	}

	rows, err := parse(entries)
	if err != nil {
		return "", err
	}

	// sort the rows by date, description, cents
	sort.Slice(rows, func(i, j int) bool {
		if rows[i].date.Before(rows[j].date) {
			return true
		}
		if rows[i].date.After(rows[j].date) {
			return false
		}
		if rows[i].description < rows[j].description {
			return true
		}
		if rows[i].description > rows[j].description {
			return false
		}
		return rows[i].cents < rows[j].cents
	})

	header := buildHeader(locale)
	body := buildBody(locale, currency, rows)
	return header + body, nil
}

// row is an entry with a strongly-typed date
type row struct {
	date        time.Time
	description string
	cents       int
}

// parse checks all the incoming entries to make sure the date is formatted
// correctly, and returns the strongly-typed rows
func parse(entries []Entry) ([]row, error) {
	r := make([]row, len(entries))
	for i, e := range entries {
		date, err := time.Parse("2006-01-02", e.Date)
		if err != nil {
			return nil, errors.New("date format invalid")
		}
		r[i] = row{
			date:        date,
			description: e.Description,
			cents:       e.Change,
		}
	}
	return r, nil
}

func buildHeader(locale string) string {
	const format string = "%-10s | %-25s | %s\n"
	switch locale {
	case "nl-NL":
		return fmt.Sprintf(format, "Datum", "Omschrijving", "Verandering")
	default:
		return fmt.Sprintf(format, "Date", "Description", "Change")
	}
}

func buildBody(locale, currency string, rows []row) string {
	var body strings.Builder
	for _, r := range rows {
		date := formatDate(locale, r.date)
		desc := fixedLength(r.description, 25)
		amount := formatAmount(locale, currency, r.cents)
		fmt.Fprintf(&body, "%-10s | %-25s | %13s\n", date, desc, amount)
	}
	return body.String()
}

func formatDate(locale string, d time.Time) string {
	if locale == "nl-NL" {
		return d.Format("02-01-2006")
	} else {
		return d.Format("01/02/2006")
	}
}

func fixedLength(s string, length int) string {
	if len(s) > length {
		return s[:length-3] + "..."
	}
	return fmt.Sprintf("%-"+strconv.Itoa(length)+"s", s)
}

func formatAmount(locale, currency string, cents int) string {
	negative := false
	if cents < 0 {
		cents = cents * -1
		negative = true
	}
	amount := ""
	if locale == "nl-NL" {
		if currency == "EUR" {
			amount += "€"
		} else {
			amount += "$"
		}
		amount += " "
		centsStr := strconv.Itoa(cents)
		switch len(centsStr) {
		case 1:
			centsStr = "00" + centsStr
		case 2:
			centsStr = "0" + centsStr
		}
		rest := centsStr[:len(centsStr)-2]
		var parts []string
		for len(rest) > 3 {
			parts = append(parts, rest[len(rest)-3:])
			rest = rest[:len(rest)-3]
		}
		if len(rest) > 0 {
			parts = append(parts, rest)
		}
		for i := len(parts) - 1; i >= 0; i-- {
			amount += parts[i] + "."
		}
		amount = amount[:len(amount)-1]
		amount += ","
		amount += centsStr[len(centsStr)-2:]
		if negative {
			amount += "-"
		} else {
			amount += " "
		}
	} else {
		if negative {
			amount += "("
		}
		if currency == "EUR" {
			amount += "€"
		} else {
			amount += "$"
		}
		centsStr := strconv.Itoa(cents)
		switch len(centsStr) {
		case 1:
			centsStr = "00" + centsStr
		case 2:
			centsStr = "0" + centsStr
		}
		rest := centsStr[:len(centsStr)-2]
		var parts []string
		for len(rest) > 3 {
			parts = append(parts, rest[len(rest)-3:])
			rest = rest[:len(rest)-3]
		}
		if len(rest) > 0 {
			parts = append(parts, rest)
		}
		for i := len(parts) - 1; i >= 0; i-- {
			amount += parts[i] + ","
		}
		amount = amount[:len(amount)-1]
		amount += "."
		amount += centsStr[len(centsStr)-2:]
		if negative {
			amount += ")"
		} else {
			amount += " "
		}
	}
	return amount
}
