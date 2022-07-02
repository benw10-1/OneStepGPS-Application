package api

import (
	. "server/config"
	// . "server/store"
	"net/http"
	"fmt"
	"io/ioutil"
)


// func GetAPI(w http.ResponseWriter, r *http.Request) {
// 	resp, err := http.Get(URL)

// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	resBody, err := ioutil.ReadAll(resp.Body)

// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	fmt.Println("Got API info")

// 	

// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Access-Control-Allow-Headers", "*")

// 	w.Write(resBody)
// }

