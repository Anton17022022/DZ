package temporarydb

import (
	"3-validation-api/pkg/processing"
	"bytes"
	"encoding/json"
	"log"
	"os"
	"strings"
)

const fileName = "internal/temporarydb/temporarydb.json"

type TemporaryDb struct {
}

func NewTemporaryDb() *TemporaryDb {
	return &TemporaryDb{}
}

func (db *TemporaryDb) CheckAccountRegistration(hash string) bool {
	file := db.ReadFile()
	lines := strings.Split(string(*file), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		var lineMap map[string]any
		err := json.Unmarshal([]byte(line), &lineMap)
		if err != nil {
			log.Println(err.Error())
		}
		if lineMap["hash"] == hash {
			return true
		}
	}
	return false
}

func (db *TemporaryDb) RegisteryAcc(email *string) (string, error) {
	hash := processing.HashCreated(email)
	account := Account{
		Email: *email,
		Hash:  hash,
	}
	if !bytes.Contains(*db.ReadFile(), account.ToBytes()) {
		err := db.RegisteryAccDb(db.ReadFile(), account.ToBytes())
		if err != nil {
			return "", err
		}
	}
	return hash, nil
}

func (db *TemporaryDb) ReadFile() *[]byte {
	file, err := os.ReadFile(fileName)
	if err != nil {
		log.Panic(err.Error())
	}
	return &file
}

func (db *TemporaryDb) RegisteryAccDb(data *[]byte, acc []byte) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err

	}
	defer file.Close()
	dataForWrite := make([]byte, 0, len(*data)+len(acc)+2)
	dataForWrite = append(dataForWrite, (*data)...)
	dataForWrite = append(dataForWrite, []byte("\n")...)
	dataForWrite = append(dataForWrite, acc...)
	_, err = file.Write(dataForWrite)
	if err != nil {
		return err
	}
	return nil
}
