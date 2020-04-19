package receiver

import (
	"fmt"
	"log"
	"vk-to-telegram/config"
	"vk-to-telegram/tools"
)

type LongPollInit struct {
	Response struct {
		Key    string `json:"key"`
		Server string `json:"server"`
		Ts     string `json:"ts"`
	} `json:"response"`
}

func StartPolling() {
	fmt.Println("Starting vk receiver...")
	// get key, server and ts
	group_id := config.Data.VkGroupId
	version := config.Data.VkApiVersion
	token := config.Data.VkAccessToken
	query_url := "https://api.vk.com/method/groups.getLongPollServer?group_id=" + group_id + "&v=" + version + "&access_token=" + token

	lp_data := LongPollInit{}
	err := tools.GetJson(query_url, &lp_data)

	if err != nil {
		log.Fatal("Cannot init Long Poll")
	}

	key := lp_data.Response.Key
	server := lp_data.Response.Server
	ts := lp_data.Response.Ts

	fmt.Println("Key: " + key)
	fmt.Println("Server: " + server)
	fmt.Println("TS: " + ts)
}
