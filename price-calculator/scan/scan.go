package scan

import (
	"bufio"
	"os"
	"strings"
)

func TrimReturnKey(text *string) {
	*text = strings.TrimSuffix(*text, "\n")
	*text = strings.TrimSuffix(*text, "\r")

}

func ReadTextFromCML() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	TrimReturnKey(&text)
	return text, nil
}
