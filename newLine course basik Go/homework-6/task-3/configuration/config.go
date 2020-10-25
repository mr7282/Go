package configuration

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"
)

//Configuration - contains configuration fields
type Configuration struct {
	ADDRESS string `json: "address"`
	TYPE    string `json: "type"`
}

//Config - application setup
func Config() Configuration {
	jsonFile, err := os.Open("config.json")
	if err != nil {
		logrus.Fatalln(err)
	}
	logrus.Info("Succesfuly opened config.json")
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		logrus.Fatalln(err)
	}
	logrus.Info("Succesfuly read config.json")

	var conf Configuration

	json.Unmarshal(byteValue, &conf)
	logrus.Info("Configuration completed")
	return conf
}
