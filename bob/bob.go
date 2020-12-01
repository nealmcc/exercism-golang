// Package bob solves the side exercise 'Bob'
package bob

import "strings"

// Hey initiates a conversation with Bob
func Hey(remark string) string {
	trimmed := strings.TrimSpace(remark)
	tone := readTone(&trimmed)
	return teenResponses[tone]
}

type tone int

const (
	question = tone(iota)
	yell
	yellQuestion
	standard
	zero
)

var teenResponses = map[tone]string{
	question:     "Sure.",
	yell:         "Whoa, chill out!",
	yellQuestion: "Calm down, I know what I'm doing!",
	standard:     "Whatever.",
	zero:         "Fine. Be that way!",
}

func readTone(remark *string) tone {
	if isEmpty(remark) {
		return zero
	}
	if isQuestion(remark) {
		if isYell(remark) {
			return yellQuestion
		}
		return question
	}
	if isYell(remark) {
		return yell
	}
	return standard
}

func isEmpty(s *string) bool {
	return *s == ""
}

func isQuestion(s *string) bool {
	str := *s
	length := len(str)
	lastChar := str[length-1]
	return lastChar == '?'
}

func isYell(s *string) bool {
	const caps = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return strings.ContainsAny(*s, caps) && (*s == strings.ToUpper(*s))
}
