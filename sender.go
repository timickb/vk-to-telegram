package main

import (
	"log"
	"vk-to-telegram/config"
	"vk-to-telegram/structs"
	"vk-to-telegram/tools"
)

func SendToTelegram(msg structs.Message) int {
	token := config.Data.TelegramToken
	chat_id := config.Data.TelegramChatId

	text := GetFullMessageText(msg)

	query_url := "https://api.telegram.org/bot" + token + "/sendMessage?chat_id=" + chat_id + "&text=\"" + text + "\"&parse_mode=HTML"

	log.Println("Query to telegram: ", query_url)

	res := structs.TelegramResponse{}
	err := tools.GetJson(query_url, &res)

	if err != nil {
		return 1
	}

	if res.Ok == false {
		return 2
	}
	return 0
}

func GetFullMessageText(graph structs.Message) string {
	result := graph.Text
	var queue []structs.Message
	queue = append(queue, graph)

	for len(queue) > 0 {
		msg := queue[0]
		queue = queue[1:]
		for i := 0; i < len(msg.FwdMessages); i++ {
			result += "<br>===============<br>"
			result += msg.FwdMessages[i].Text
			queue = append(queue, msg.FwdMessages[i].FwdMessages...)
		}
	}

	return result
}

func GetAttachments(msg structs.Message) []string {
	var result []string
	for i := 0; i < len(msg.Attachments); i++ {
		att := msg.Attachments[i]
		if att.Type == "photo" {
			last_url_index := len(msg.Attachments[i].Photo.Sizes) - 1
			result = append(result, msg.Attachments[i].Photo.Sizes[last_url_index].URL)
		}
		if att.Type == "doc" {
			result = append(result, msg.Attachments[i].Doc.URL)
		}
	}
	return result
}
