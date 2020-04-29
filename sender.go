package main

import (
	"log"
	"net/url"
	"vk-to-telegram/config"
	"vk-to-telegram/structs"
	"vk-to-telegram/tools"
)

func SendToTelegram(msg structs.Message) int {
	token := config.Data.TelegramToken
	chat_id := config.Data.TelegramChatId

	text := url.QueryEscape(GetFullMessageText(msg))

	base_url := "https://api.telegram.org/bot" + token + "/sendMessage?chat_id=" + chat_id

	query_url := base_url + "&text=" + text + "&parse_mode=HTML"

	log.Println("Query to telegram: ", query_url)

	res := structs.TelegramResponse{}
	err := tools.GetJson(query_url, &res)

	if err != nil {
		return 1
	}

	if res.Ok == false {
		return 2
	}
	if len(msg.Attachments) > 0 {
		log.Println("Found", len(msg.Attachments), "attachments. Let's send it.")
		for _, att := range msg.Attachments {
			if att.Type == "photo" {
				last_url_index := len(att.Photo.Sizes) - 1
				url := att.Photo.Sizes[last_url_index].URL
				photo_query := "https://api.telegram.org/bot" + token + "/sendPhoto?url=" + url
				res := structs.TelegramResponse{}
				err := tools.GetJson(photo_query, &res)
				if err != nil {
					log.Panic(err)
					log.Println("Couldn't send attachment(")
				}
			}
			if att.Type == "doc" {
				url := att.Doc.URL
				doc_query := "https://api.telegram.org/bot" + token + "/sendDocument?url=" + url
				res := structs.TelegramResponse{}
				err := tools.GetJson(doc_query, &res)
				if err != nil {
					log.Panic(err)
					log.Println("Couldn't send attachment(")
				}
			}
		}
	}

	return 0
}

func GetFullMessageText(graph structs.Message) string {
	result := graph.Text
	// TODO how to organize forwarded message in telegram
	/*var queue []structs.Message
	queue = append(queue, graph)

	depth := 1

	for len(queue) > 0 {
		msg := queue[0]
		queue = queue[1:]
		size := len(msg.FwdMessages)
		if size > 0 {
			result += "\n" + strings.Repeat("-", size) + "Forwarded:\n"
		}
		for i := 0; i < size; i++ {
			result += "\n===============\n"
			result += msg.FwdMessages[i].Text
			queue = append(queue, msg.FwdMessages[i].FwdMessages...)
		}
		if size > 0 {
			result += "\n" + strings.Repeat("-", size) + "Forwarded:\n"
		}
	}*/

	return result
}

func GetDocURL(doc structs.Doc) string {
	return doc.URL
}

func GetPhotoURL(photo structs.Photo) string {
	last_url_index := len(photo.Sizes) - 1
	return photo.Sizes[last_url_index].URL
}
