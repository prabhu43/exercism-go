package bob

import (
	"regexp"
	"strings"
)

// Hey returns the Bob's response for the given remark
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	if isNotSayingAnything(remark) {
		return "Fine. Be that way!"
	} else if isYellingQuestion(remark) {
		return "Calm down, I know what I'm doing!"
	} else if isYelling(remark) {
		return "Whoa, chill out!"
	} else if isQuestioning(remark) {
		return "Sure."
	}

	return "Whatever."
}

func isNotSayingAnything(remark string) bool {
	emptyString := ""
	return remark == emptyString
}

func isYellingQuestion(remark string) bool {
	return isYelling(remark) && isQuestioning(remark)
}

func isYelling(remark string) bool {
	isUpper := strings.ToUpper(remark) == remark
	containsAtleastAnAlphabet, _ := regexp.MatchString("[a-zA-Z]", remark)
	return isUpper && containsAtleastAnAlphabet
}

func isQuestioning(remark string) bool {
	questionChar := "?"
	if strings.HasSuffix(remark, questionChar) {
		return true
	}
	return false
}
