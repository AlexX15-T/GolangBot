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


func GetConfig(a *Config) error {
	file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		f.Println(err.Error())
		return err
	}

	err = json.Unmarshal(file, &a)

	if err != nil {
		f.Println(err.Error())
		return err
	}
	
	return nil
}
