package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Convert(data string) (string, error) {
	data = strings.TrimSpace(data)
	if isMorse(data) {
		return morse.ToText(data), nil
	}
	return morse.ToMorse(data), nil
}

func isMorse(s string) bool {
	hasMorseChar := false
	for _, char := range s {
		switch char {
		case '.', '-':
			hasMorseChar = true
		case ' ':
			continue
		default:
			return false
		}
	}
	return hasMorseChar
}
