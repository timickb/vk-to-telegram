package config

import (
	"encoding/json"
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Config struct {
	VkKey         string `json:"vk_key"`
	VkServer      string `json:"vk_server"`
	VkTs          int    `json:"vk_ts"`
	TelegramToken string `json:"telegram_token"`
	Port          int    `json:"port"`
}

var Data = Config{}

func ReadConfig() {
	file, err := ioutil.ReadFile("config/config.json")
	check(err)

	_ = json.Unmarshal([]byte(file), &Data)
}
