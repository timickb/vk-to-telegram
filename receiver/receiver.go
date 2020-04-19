package receiver

import (
	"log"
	"time"
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

type ReceivedData struct {
	Ts      string `json:"ts"`
	Updates []struct {
		Type   string `json:"type"`
		Object struct {
			Message struct {
				Date                  int           `json:"date"`
				FromID                int           `json:"from_id"`
				ID                    int           `json:"id"`
				Out                   int           `json:"out"`
				PeerID                int           `json:"peer_id"`
				Text                  string        `json:"text"`
				ConversationMessageID int           `json:"conversation_message_id"`
				FwdMessages           []interface{} `json:"fwd_messages"`
				Important             bool          `json:"important"`
				RandomID              int           `json:"random_id"`
				Attachments           []interface{} `json:"attachments"`
				IsHidden              bool          `json:"is_hidden"`
			} `json:"message"`
			ClientInfo struct {
				ButtonActions  []string `json:"button_actions"`
				Keyboard       bool     `json:"keyboard"`
				InlineKeyboard bool     `json:"inline_keyboard"`
				LangID         int      `json:"lang_id"`
			} `json:"client_info"`
		} `json:"object"`
		GroupID int    `json:"group_id"`
		EventID string `json:"event_id"`
	} `json:"updates"`
}

func StartPolling() {
	log.Println("Initializing VK Bot Long Poll API...")
	// get key, server and ts
	group_id := config.Data.VkGroupId
	version := config.Data.VkApiVersion
	token := config.Data.VkAccessToken
	initial_url := "https://api.vk.com/method/groups.getLongPollServer?group_id=" + group_id + "&v=" + version + "&access_token=" + token

	log.Println("Long Poll initial url is " + initial_url)

	lp_data := LongPollInit{}
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

	recv_data := ReceivedData{}

	// main update receiver loop
	for {
		err = tools.GetJson(query_url, &recv_data)
		if err != nil {
			log.Fatal("An error occured while making request.")
		}
		// update last update id
		ts = recv_data.Ts
		log.Println("New TS: " + recv_data.Ts)
		log.Println("Received", len(recv_data.Updates), "updates")

		for i := 0; i < len(recv_data.Updates); i++ {
			log.Println("Update number", (i + 1), "type is "+recv_data.Updates[i].Type)
			if recv_data.Updates[i].Type == "message_new" {
				log.Println("New message text:", recv_data.Updates[i].Object.Message.Text)
			}
		}

		time.Sleep(time.Second * 10)
	}

}
