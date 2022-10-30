package utility

import (
	"encoding/json"
	f "fmt"
	"io/ioutil"
)


type Config struct{
	Token string `json:"Token"`
	BotPrefix string `json:"BotPrefix"`
}


func (a *Config) GetConfigerror() error {
	file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		f.Println(err.Error())
		return err
	}

	f.Println(string(file))

	err = json.Unmarshal(file, &a)

	if err != nil {
		f.Println(err.Error())
		return err
	}

	return nil
}
