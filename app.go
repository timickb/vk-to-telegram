package main

import (
	"log"
	"time"
	"vk-to-telegram/config"
	"vk-to-telegram/structs"
	"vk-to-telegram/tools"
)

func StartPolling() {
	log.Println("Initializing VK Bot Long Poll API...")
	// get key, server and ts
	group_id := config.Data.VkGroupId
	version := config.Data.VkApiVersion
	token := config.Data.VkAccessToken
	initial_url := "https://api.vk.com/method/groups.getLongPollServer?group_id=" + group_id + "&v=" + version + "&access_token=" + token

	log.Println("Long Poll initial url is " + initial_url)

	lp_data := structs.LongPollInit{}
	err := tools.GetJson(initial_url, &lp_data)

	if err != nil {
		log.Fatal("Cannot init Long Poll")
	}

	key := lp_data.Response.Key
	server := lp_data.Response.Server
	ts := lp_data.Response.Ts

	log.Println("Long Poll initialized with variables: ")
	log.Println("Key: " + key)
	log.Println("Server: " + server)
	log.Println("TS: " + ts)

	query_url := "" + server + "?act=a_check&key=" + key + "&ts=" + ts + "&wait=1"

	recv_data := structs.ReceivedData{}

	// main update receiver loop
	for {
		err = tools.GetJson(query_url, &recv_data)
		if err != nil {
			log.Fatal("An error occured while making request.")
		}
		// update last update id
		ts = recv_data.Ts
		query_url = "" + server + "?act=a_check&key=" + key + "&ts=" + ts + "&wait=1"

		log.Println("New TS: " + recv_data.Ts)
		log.Println("Received", len(recv_data.Updates), "updates")

		for i := 0; i < len(recv_data.Updates); i++ {
			log.Println("Update number", (i + 1), "type is "+recv_data.Updates[i].Type)
			if recv_data.Updates[i].Type == "message_new" {
				log.Println("New message from:", recv_data.Updates[i].Object.Message.FromID)
				result := SendToTelegram(recv_data.Updates[i].Object.Message)
				if result == 0 {
					log.Println("Message sent to telegram!")
				}
				if result == 1 {
					log.Println("An error occured while sending GET query.")
				}
				if result == 2 {
					log.Println("Telegram denied this query.")
				}
			}
		}

		time.Sleep(time.Second * 10)
	}

}

func main() {
	log.Println("VK TO TELEGRAM MESSAGE RESENDER")
	config.ReadConfig()
	StartPolling()
}
