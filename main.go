package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"vk-to-telegram/config"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	config.ReadConfig()
	fmt.Println(config.Data.TelegramToken)
}
