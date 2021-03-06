package structs

type TelegramResponse struct {
	Ok     bool `json:"ok"`
	Result struct {
		MessageID int `json:"message_id"`
		From      struct {
			ID        int    `json:"id"`
			IsBot     bool   `json:"is_bot"`
			FirstName string `json:"first_name"`
			Username  string `json:"username"`
		} `json:"from"`
		Chat struct {
			ID        int    `json:"id"`
			FirstName string `json:"first_name"`
			Username  string `json:"username"`
			Type      string `json:"type"`
		} `json:"chat"`
		Date int    `json:"date"`
		Text string `json:"text"`
	} `json:"result"`
}

type LongPollInit struct {
	Response struct {
		Key    string `json:"key"`
		Server string `json:"server"`
		Ts     string `json:"ts"`
	} `json:"response"`
}

type Message struct {
	Date                  int          `json:"date"`
	FromID                int          `json:"from_id"`
	ID                    int          `json:"id"`
	Out                   int          `json:"out"`
	PeerID                int          `json:"peer_id"`
	Text                  string       `json:"text"`
	ConversationMessageID int          `json:"conversation_message_id"`
	FwdMessages           []Message    `json:"fwd_messages"`
	Important             bool         `json:"important"`
	RandomID              int          `json:"random_id"`
	Attachments           []Attachment `json:"attachments"`
	IsHidden              bool         `json:"is_hidden"`
}

type Doc struct {
	ID        int    `json:"id"`
	OwnerID   int    `json:"owner_id"`
	Title     string `json:"title"`
	Size      int    `json:"size"`
	Ext       string `json:"ext"`
	URL       string `json:"url"`
	Date      int    `json:"date"`
	Type      int    `json:"type"`
	AccessKey string `json:"access_key"`
}

type Photo struct {
	ID      int `json:"id"`
	AlbumID int `json:"album_id"`
	OwnerID int `json:"owner_id"`
	Sizes   []struct {
		Type   string `json:"type"`
		URL    string `json:"url"`
		Width  int    `json:"width"`
		Height int    `json:"height"`
	} `json:"sizes"`
	Text      string `json:"text"`
	Date      int    `json:"date"`
	AccessKey string `json:"access_key"`
}

type Attachment struct {
	Type  string `json:"type"`
	Photo Photo  `json:"photo"`
	Doc   Doc    `json:"doc"`
}

type ReceivedData struct {
	Ts      string `json:"ts"`
	Updates []struct {
		Type   string `json:"type"`
		Object struct {
			Message    Message `json:"message"`
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
