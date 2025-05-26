package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func isMorse(str string) bool {
	return strings.ContainsAny(str, ".-")
}
func Translate(str string) (string, error) {
	if len(str) == 0 {
		return "", errors.New("empty input")
	}

	if isMorse(str) {
		return morse.ToText(str), nil
	} else {
		return morse.ToMorse(str), nil
	}

}
