package main

import (
	"github.com/timickb/vk-to-telegram/internal/config"
	"github.com/timickb/vk-to-telegram/internal/sender"
	"github.com/timickb/vk-to-telegram/internal/structs"
	"github.com/timickb/vk-to-telegram/internal/tools"
	"log"
	"time"
)

func StartPolling() {
	log.Println("Initializing VK Bot Long Poll API...")
	// get key, server and ts
	groupId := config.Data.VkGroupId
	version := config.Data.VkApiVersion
	token := config.Data.VkAccessToken
	initialUrl := "https://api.vk.com/method/groups.getLongPollServer?group_id=" + groupId + "&v=" + version + "&access_token=" + token

	log.Println("Long Poll initial url is " + initialUrl)

	lpData := structs.LongPollInit{}
	err := tools.GetJson(initialUrl, &lpData)

	if err != nil {
		log.Fatal("Cannot init Long Poll")
	}

	key := lpData.Response.Key
	server := lpData.Response.Server
	ts := lpData.Response.Ts

	log.Println("Long Poll initialized with variables: ")
	log.Println("Key: " + key)
	log.Println("Server: " + server)
	log.Println("TS: " + ts)

	queryUrl := "" + server + "?act=a_check&key=" + key + "&ts=" + ts + "&wait=1"

	recvData := structs.ReceivedData{}

	// main update receiver loop
	for {
		err = tools.GetJson(queryUrl, &recvData)
		if err != nil {
			log.Fatal("An error occured while making request.")
		}
		// update last update id
		ts = recvData.Ts
		queryUrl = "" + server + "?act=a_check&key=" + key + "&ts=" + ts + "&wait=1"

		log.Println("New TS: " + recvData.Ts)
		log.Println("Received", len(recvData.Updates), "updates")

		for i := 0; i < len(recvData.Updates); i++ {
			log.Println("Update number", i+1, "type is "+recvData.Updates[i].Type)
			if recvData.Updates[i].Type == "message_new" {
				log.Println("New message from:", recvData.Updates[i].Object.Message.FromID)
				result := sender.SendToTelegram(recvData.Updates[i].Object.Message)
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
