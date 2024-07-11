package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)


func Hello(name string) (string, error){
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf(randomFormat() , name)
	return message, nil
}


func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names{
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}


func init() {
	rand.Seed(time.Now().UnixNano())
}


func randomFormat() string{
	mess := []string{
		"Hi, %v.",
		"Hello, %v.",
		"Welcom, %v",
	}
	return mess[rand.Intn(len(mess))]
}
