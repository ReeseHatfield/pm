package data

import "fmt"

type PMDictionary map[string]Credentials

type Credentials struct {
	Username string
	Password string
}

func (c Credentials) String() string {
	return fmt.Sprintf("Username: %s \nPassword: %s", c.Username, c.Password)
}
