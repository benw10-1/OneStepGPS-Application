package api

import (
	. "server/config"
	"net/http"
	"fmt"
	"io/ioutil"
)

func GetAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting API info")
	
	URL := "https://track.onestepgps.com/v3/api/public/device?latest_point=true&api-key=" + GetEnv("API_KEY", "")

	resp, err := http.Get(URL)

	if err != nil {
		fmt.Println(err)
	}

	resBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Got API info")

	w.Write(resBody)
}