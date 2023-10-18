package utility

import (
	"encoding/json"
	_ "github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
	"io"
	"log"
	"os"
)

func Handle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func GenerateMnemonicKey() string {
	entropy, _ := bip39.NewEntropy(128)
	mnemonic, _ := bip39.NewMnemonic(entropy)
	return mnemonic
}

func ValidateMnemonic(mnemonic string) bool {
	return bip39.IsMnemonicValid(mnemonic)
}

func SaveFile(address string, mnemonic string) {
	path := "accounts/" + address + ".json"
	file, err := os.Create(path)
	defer file.Close()
	Handle(err)

	// Marshal the data
	data, err := json.MarshalIndent(map[string]string{
		"Mnemonic": mnemonic,
		"Address": address,
	}, "", "  ")

	// Write data to file
	io.WriteString(file, string(data))
}