package service

import (
	"errors"
	"strings"
)

var textToMorse = map[rune]string{}
var morseToText = map[string]string{}

func init() {
	for k, v := range textToMorse {
		morseToText[v] = string(k)
	}
}

func AutoDefect(data string) (string, error) {
	if isMorse(data) {
		return decodeMorse(data), nil
	}
	return encodeText(data), nil
}

func isMorse(input string) bool {
	for _, char := range input {
		if char != '.' && char != '-' && char != ' ' {
			return false
		}
	}
	return true
}

// кодирование текста в морзе
func encodeText(text string) string {
	var result strings.Builder
	text = strings.ToUpper(text)

	for _, char := range text {
		code, exists := textToMorse[char]
		if exists {
			result.WriteString(code)
			result.WriteString(" ")
		}
	}
	return strings.TrimSpace(result.String())

}

// декодирование морзе в текст
func decodeMorse(morse string) string {
	var result strings.Builder
	words := strings.Split(morse, "  ")

	for _, word := range words {
		letters := strings.Split(word, " ")

		for _, letter := range letters {
			char, exists := morseToText[letter]
			if exists {
				result.WriteString(char)
			}
		}
		result.WriteString(" ")
	}

	return strings.TrimSpace(result.String())
}

// обработка ошибок
func HandleError(err error) error {
	if err != nil {
		return errors.New("Ошибка при конвертации: " + err.Error())

	}
	return nil
}
