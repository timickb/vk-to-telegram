package receiver

import (
	"fmt"
	"net/http"
	"vk-to-telegram/config"
)

func StartPolling() {
	fmt.Println("Starting vk receiver...")
	key := config.Data.VkKey
	server := config.Data.VkServer
	ts := config.Data.VkTs
	url := server + "?act=a_check&key=" + key + "&ts=" + string(ts) + "&wait=25&mode=2&version=3"
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Cannot complete query to API")
	}
	fmt.Println(resp.Body)
}
