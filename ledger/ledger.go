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

func FormatLedger(currency string, locale string, entries []Entry) (string, error) {
	if currency != "EUR" && currency != "USD" {
		return "", errors.New("invalid currency requested")
	}

	if locale != "nl-NL" && locale != "en-US" {
		return "", errors.New("invalid locale requested")
	}

	header := buildHeader(locale)

	entriesCopy := make([]Entry, 0, len(entries))
	entriesCopy = append(entriesCopy, entries...)

	sortEntries(entriesCopy)

	body, err := buildBody(locale, currency, entriesCopy)
	if err != nil {
		return "", err
	}

	return header + body, nil
}

func sortEntries(es []Entry) {
	m1 := map[bool]int{true: 0, false: 1}
	m2 := map[bool]int{true: -1, false: 1}
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
}

func buildHeader(locale string) string {
	switch locale {
	case "nl-NL":
		return "Datum" + strings.Repeat(" ", 10-len("Datum")) +
			" | " + "Omschrijving" + strings.Repeat(" ", 25-len("Omschrijving")) +
			" | " + "Verandering" + "\n"
	default:
		return "Date" + strings.Repeat(" ", 10-len("Date")) +
			" | " + "Description" + strings.Repeat(" ", 25-len("Description")) +
			" | " + "Change" + "\n"
	}
}

func buildBody(locale, currency string, entriesCopy []Entry) (string, error) {
	var body strings.Builder
	for _, entry := range entriesCopy {
		if len(entry.Date) != 10 {
			return "", errors.New("date format invalid")
		}
		yyyy, d2, mm, d4, dd := entry.Date[0:4], entry.Date[4], entry.Date[5:7], entry.Date[7], entry.Date[8:10]
		if d2 != '-' {
			return "", errors.New("year-month separator invalid")
		}
		if d4 != '-' {
			return "", errors.New("month-day separator invalid")
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
		var amount string
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
		var al int
		for range amount {
			al++
		}
		body.WriteString(date + strings.Repeat(" ", 10-len(date)) + " | " + desc + " | " + strings.Repeat(" ", 13-al) + amount + "\n")
	}

	return body.String(), nil
}
