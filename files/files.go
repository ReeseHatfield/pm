package files

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
	"os/user"
	"pm/crypt"
	"pm/data"
)

func LoadPmFile(key []byte) (data.PMDictionary, error) {
	var pmDict data.PMDictionary

	path := GetPMPath()
	file, err := os.Open(path)
	if err != nil {
		return pmDict, fmt.Errorf("failed to open PM file: %w", err)
	}
	defer file.Close()

	encryptedDict, err := os.ReadFile(path)
	if err != nil {
		return pmDict, fmt.Errorf("failed to read PM file: %w", err)
	}

	serializedDict, err := crypt.Decrypt(encryptedDict, key)
	if err != nil {
		return pmDict, fmt.Errorf("failed to decrypt PM file: %w", err)
	}

	buffer := bytes.NewBuffer(serializedDict)
	deserializer := gob.NewDecoder(buffer)
	err = deserializer.Decode(&pmDict)
	if err != nil {
		return pmDict, fmt.Errorf("failed to deserialize PMDictionary: %w", err)
	}

	return pmDict, nil
}

// this should be made to return an error not a status to be more go like
func SavePmFile(key []byte, dict data.PMDictionary) (status bool) {

	// serialize PMDictionary
	buffer := new(bytes.Buffer)
	serializer := gob.NewEncoder(buffer)

	err := serializer.Encode(dict)

	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	serializedDict := buffer.Bytes()

	encryptedDict, err := crypt.Encrypt(serializedDict, key)

	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	path := GetPMPath()

	file, err := os.Create(path)

	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	defer file.Close()

	_, err = file.Write(encryptedDict)

	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	return true
}

func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	if err == nil {
		return true
	}

	if os.IsNotExist(err) {
		return false
	}

	return false
}

func GetPMPath() string {

	usr, err := user.Current()
	if err != nil {
		fmt.Println(err)
	}

	pmDotDat := usr.HomeDir + "/.pm/pm.dat"

	return pmDotDat
}
