package utils

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

var CONFIG_DIR string
var CONFIG_FILE string

type Network struct {
	Provider string `json:"provider"`
}

type ConfigFile struct {
	Networks map[string]Network `json:"networks"`
}

func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic("can not find home directory")
	}

	CONFIG_DIR = filepath.Join(homeDir, ".config", "supher")
	CONFIG_FILE = CONFIG_DIR + "/config.json"
}

func CheckConfigFile() bool {
	if _, err := os.Stat(CONFIG_FILE); err == nil {
		return true
	} else if errors.Is(err, os.ErrNotExist) {
		return false
	} else {
		panic("Unhandled config file exists situation!")
	}
}

func CreateConfigFile() error {
	if err := os.MkdirAll(CONFIG_DIR, os.ModePerm); err != nil {
		return err
	}

	_, err := os.Create(CONFIG_FILE)
	if err != nil {
		return err
	}

	ethereumNetwork := map[string]Network{
		"ethereum": Network{
			Provider: "YOUR_JSON_RPC_PROVIDER",
		},
	}

	// create structure for ethereum
	emptyConfigFile := ConfigFile{
		Networks: ethereumNetwork,
	}

	file, err := json.MarshalIndent(emptyConfigFile, "", " ")
	if err != nil {
		panic("error with marshaling empty data")
	}

	err = ioutil.WriteFile(CONFIG_FILE, file, os.ModePerm)
	if err != nil {
		panic("error with writing empty data to config file")
	}

	return nil
}
