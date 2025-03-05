package temporarydb

import (
	"encoding/json"
	"log"
)

type Account struct {
	Email string `json:"email"`
	Hash  string `json:"hash"`
}

func (acc *Account) ToBytes() []byte {
	file, err := json.Marshal(acc)
	if err != nil {
		log.Panic(err.Error())
	}
	return file
}


