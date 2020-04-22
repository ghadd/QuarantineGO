package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const cAPIKey = "90b34303c90966d9988b014c6388d0f4"

var (
	myClient = &http.Client{Timeout: 10 * time.Second}
	city     string
	data     *cityData
)

type weatherData struct {
	TempKelvins float32 `json:"temp"`
	FeelsLike   float32 `json:"feels_like"`
	MinKelvins  float32 `json:"temp_min"`
	MaxKelvins  float32 `json:"temp_max"`
	Pressure    int     `json:"pressure"`
	Humidity    int     `json:"hunidity"`
}

type wind struct {
	Speed float32 `json:"speed"`
	Deg   float32 `json:"deg"`
}

type cityData struct {
	Main       weatherData `json:"main"`
	Visibility int         `json:"visibility"`
	Wind       wind        `json:"wind"`
	Name       string      `json:"name"`
}

func getURL(city string) (url string) {
	url = "http://api.openweathermap.org/data/2.5/weather?q=" + city + "&appid=" + cAPIKey
	return
}

func getJSON(city string) *cityData {
	r, err := myClient.Get(getURL(city))
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	data = new(cityData)
	json.Unmarshal(body, data)

	return data
}

func main() {
	fmt.Print("Enter city to check weather: ")
	fmt.Scan(&city)

	getJSON(city)
	data = getJSON(city)

	file, _ := json.MarshalIndent(data, "", " ")

	_ = ioutil.WriteFile("src/weatherTracker/data/data.json", file, 0644)
}
