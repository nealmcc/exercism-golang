package phonenumber

import (
	"errors"
	"fmt"
	"regexp"
)

var (
	re *regexp.Regexp = regexp.MustCompile(`^` +
		`\+?1?[\W.-]*` +
		`\(?(?P<area>[2-9]\d{2})\)?[\W.-]*` +
		`(?P<exchange>[2-9]\d{2})[\W.-]*` +
		`(?P<subscriber>\d{4})[\W]*` +
		`$`,
	)
)

func tryParse(in string) (tel, error) {
	matches := re.FindAllStringSubmatch(in, -1)
	if len(matches) == 0 {
		return tel{}, errors.New("invalid telephone format")
	}
	t := tel{
		area:       matches[0][re.SubexpIndex("area")],
		exchange:   matches[0][re.SubexpIndex("exchange")],
		subscriber: matches[0][re.SubexpIndex("subscriber")],
	}
	return t, nil
}

func Number(in string) (string, error) {
	t, err := tryParse(in)
	if err != nil {
		return "", err
	}
	return t.area + t.exchange + t.subscriber, nil
}

func AreaCode(in string) (string, error) {
	t, err := tryParse(in)
	if err != nil {
		return "", err
	}
	return t.area, nil
}

func Format(in string) (string, error) {
	t, err := tryParse(in)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("(%s) %s-%s", t.area, t.exchange, t.subscriber), nil
}

type tel struct {
	area       string
	exchange   string
	subscriber string
}
