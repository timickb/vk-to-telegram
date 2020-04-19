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
	VkApiVersion   string `json:"vk_api_version"`
	VkGroupId      string `json:"vk_group_id"`
	VkAccessToken  string `json:"vk_access_token"`
	TelegramToken  string `json:"telegram_token"`
	TelegramChatId string `json:"telegram_chat_id"`
	Port           int    `json:"port"`
}

var Data = Config{}

func ReadConfig() {
	file, err := ioutil.ReadFile("config/config.json")
	check(err)

	_ = json.Unmarshal([]byte(file), &Data)
}
