package main

import (
	"fmt"
	"vk-to-telegram/config"
	"vk-to-telegram/receiver"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("VK-to-Telegram message resender")
	config.ReadConfig()
	fmt.Println(config.Data.TelegramToken)
	receiver.StartPolling()
}
