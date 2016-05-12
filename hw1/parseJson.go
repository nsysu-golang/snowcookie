package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Weather struct {
	Response            map[string]interface{}
	Current_observation map[string]interface{}
}

func main() {

	contents, err := getJson("https://gist.githubusercontent.com/PichuChen/6ce430b5474b037b9a4dcafb719f9db1/raw/f6d9485db787f3b860da41c523e6745ee8b3fd53/NSYSU.json")
	if err != nil {
		fmt.Println("get json failed : %s", err)
		os.Exit(1)
	}

	data, err := parseJson2Weather(contents)
	if err != nil {
		fmt.Println("parse json failed : %s", err)
		os.Exit(2)
	}

	if temp_c, ok := data.Current_observation["temp_c"]; ok {
		fmt.Println("temp_c   : ", temp_c)
	}

	if icon_url, ok := data.Current_observation["icon_url"]; ok {
		fmt.Println("icon_url : ", icon_url)
	}

}

func getJson(url string) (contents string, err error) {
	var response *http.Response
	var contents_bytes []byte
	response, err = http.Get(url)
	if err != nil {
		return
	} else {
		defer response.Body.Close()
		contents_bytes, err = ioutil.ReadAll(response.Body)
		if err != nil {
			return
		}
		contents = string(contents_bytes)
		return
	}
}

func parseJson2Weather(raw string) (d Weather, err error) {
	err = json.Unmarshal([]byte(raw), &d)
	return
}
