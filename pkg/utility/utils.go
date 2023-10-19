package utility

import (
	"encoding/json"
	"errors"
	_ "github.com/tyler-smith/go-bip32"
	env "github.com/joho/godotenv"
	"github.com/tyler-smith/go-bip39"
	"github.com/onflow/flow-go-sdk/access/http"
	"io"
	"log"
	"os"
)
var(
	Host string
	NetworkKey string
	NetworkPrivateKey string
)

var err error

func init() {
	err := env.Load()
	if err != nil {
		log.Fatal(err)
	}

	var testnet string = os.Getenv("TESTNET")
	var mainnet = os.Getenv("MAINNET")
	var emulator = os.Getenv("EMULATOR")

	var testnet_pkey = os.Getenv("TESTNET_PRIVATE_KEY")
	var mainnet_pkey = os.Getenv("MAINNET_PRIVATE_KEY")
	var emulator_pkey = os.Getenv("EMULATOR_PRIVATE_KEY")

	switch(os.Getenv("NETWORK")) {
	case "testnet":
		Host = http.TestnetHost
		NetworkKey, NetworkPrivateKey, err = validate(testnet, testnet_pkey)
		Handle(err)
	case "emulator":
		Host = http.EmulatorHost
		NetworkKey, NetworkPrivateKey, err = validate(emulator, emulator_pkey)
		Handle(err)
		break
	case "mainnet":
		Host = http.MainnetHost
		NetworkKey, NetworkPrivateKey, err = validate(mainnet, mainnet_pkey)
		Handle(err)
		break
	default:
		log.Fatal("Network Host is invalid")
	}
}

func validate(address, pkey string) (string, string, error) {
	if (len(address) != 0 || len(pkey) != 0) {
		return address, pkey, nil
	} else {
		return "", "", errors.New("Keys cannot be empty")	
	}
}

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

func ValidateMnemonicKey(mnemonic string) bool {
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