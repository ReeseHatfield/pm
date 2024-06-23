package data

type PMDictionary map[string]Credentials

type Credentials struct {
	Username string
	Password string
}
