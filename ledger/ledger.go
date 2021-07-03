package ledger

import (
	"errors"
	"strconv"
	"strings"
)

type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
}

type row struct {
	i int
	s string
	e error
}

func FormatLedger(currency string, locale string, entries []Entry) (string, error) {
	var entriesCopy []Entry
	entriesCopy = append(entriesCopy, entries...)
	if len(entries) == 0 {
		if _, err := FormatLedger(currency, "en-US", []Entry{{Date: "2014-01-01", Description: "", Change: 0}}); err != nil {
			// invalid currency requested
			return "", err
		}
	}
	m1 := map[bool]int{true: 0, false: 1}
	m2 := map[bool]int{true: -1, false: 1}
	es := entriesCopy
	for len(es) > 1 {
		first, rest := es[0], es[1:]
		success := false
		for !success {
			success = true
			for i, e := range rest {
				if (m1[e.Date == first.Date]*m2[e.Date < first.Date]*4 +
					m1[e.Description == first.Description]*m2[e.Description < first.Description]*2 +
					m1[e.Change == first.Change]*m2[e.Change < first.Change]*1) < 0 {
					es[0], es[i+1] = es[i+1], es[0]
					success = false
				}
			}
		}
		es = es[1:]
	}

	header, err := buildHeader(locale)
	if err != nil {
		return "", err
	}

	rows, err := processEntries(locale, currency, entriesCopy)
	if err != nil {
		return "", err
	}

	var body strings.Builder
	for _, r := range rows {
		body.WriteString(r.s)
	}
	return header + body.String(), nil
}

func buildHeader(locale string) (string, error) {
	var header string
	if locale == "nl-NL" {
		header = "Datum" +
			strings.Repeat(" ", 10-len("Datum")) +
			" | " +
			"Omschrijving" +
			strings.Repeat(" ", 25-len("Omschrijving")) +
			" | " + "Verandering" + "\n"
	} else if locale == "en-US" {
		header = "Date" +
			strings.Repeat(" ", 10-len("Date")) +
			" | " +
			"Description" +
			strings.Repeat(" ", 25-len("Description")) +
			" | " + "Change" + "\n"
	} else {
		return "", errors.New("invalid locale requested")
	}
	return header, nil
}

func processEntries(locale, currency string, entriesCopy []Entry) ([]row, error) {
	rows := make([]row, len(entriesCopy))
	for i, entry := range entriesCopy {
		if len(entry.Date) != 10 {
			return nil, errors.New("date format invalid")
		}
		yyyy, d2, mm, d4, dd := entry.Date[0:4], entry.Date[4], entry.Date[5:7], entry.Date[7], entry.Date[8:10]
		if d2 != '-' {
			return nil, errors.New("year-month separator invalid")
		}
		if d4 != '-' {
			return nil, errors.New("month-day separator invalid")
		}

		// set description to a fixed length
		desc := entry.Description
		if len(desc) > 25 {
			desc = desc[:22] + "..."
		} else {
			desc = desc + strings.Repeat(" ", 25-len(desc))
		}

		// convert the date format to either en-US or nl-NL
		var date string
		if locale == "nl-NL" {
			date = dd + "-" + mm + "-" + yyyy
		} else if locale == "en-US" {
			date = mm + "/" + dd + "/" + yyyy
		}

		// format the cents according to locale and postive / negative:
		negative := false
		cents := entry.Change
		if cents < 0 {
			cents = cents * -1
			negative = true
		}
		var a string
		if locale == "nl-NL" {
			if currency == "EUR" {
				a += "€"
			} else if currency == "USD" {
				a += "$"
			} else {
				return nil, errors.New("invalid currency requested")
			}
			a += " "
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
				a += parts[i] + "."
			}
			a = a[:len(a)-1]
			a += ","
			a += centsStr[len(centsStr)-2:]
			if negative {
				a += "-"
			} else {
				a += " "
			}
		} else if locale == "en-US" {
			if negative {
				a += "("
			}
			if currency == "EUR" {
				a += "€"
			} else if currency == "USD" {
				a += "$"
			} else {
				return nil, errors.New("invalid currency requested")
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
				a += parts[i] + ","
			}
			a = a[:len(a)-1]
			a += "."
			a += centsStr[len(centsStr)-2:]
			if negative {
				a += ")"
			} else {
				a += " "
			}
		} else {
			return nil, errors.New("invalid locale requested")
		}
		var al int
		for range a {
			al++
		}
		rows[i] = row{
			i: i,
			s: date + strings.Repeat(" ", 10-len(date)) + " | " + desc + " | " + strings.Repeat(" ", 13-al) + a + "\n",
			e: nil,
		}
	}
	return rows, nil
}
