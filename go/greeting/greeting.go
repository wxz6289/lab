package greeting

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Name is empty")
	}
	message := fmt.Sprintf("Hello %v", name)
	return  message, nil
}