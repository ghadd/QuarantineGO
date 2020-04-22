package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"reflect"
	"strconv"
	"time"

	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
)

const (
	cAPIKey       = "90b34303c90966d9988b014c6388d0f4"
	celsiusKelvin = 273.15
	delaySec      = 300
	maxPolling    = 480
)

var (
	myClient = &http.Client{Timeout: 10 * time.Second}
	city     string
	prevdata *cityData
	data     *cityData
)

type weatherData struct {
	TempKelvins float32 `json:"temp"`
	FeelsLike   float32 `json:"feels_like"`
	MinKelvins  float32 `json:"temp_min"`
	MaxKelvins  float32 `json:"temp_max"`
	Pressure    int     `json:"pressure"`
	Humidity    int     `json:"humidity"`
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

func display(data *cityData) {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
	fmt.Println(time.Now().Format("2006-01-02 15:04:05 Monday"))
	fmt.Printf("Absolute temperature: %fK(%fºC)\n", data.Main.TempKelvins, data.Main.TempKelvins-celsiusKelvin)
	fmt.Printf("Feels like: %fK\n", data.Main.FeelsLike)
	fmt.Printf("Min/Max temperature: %fK / %fK\n", data.Main.MinKelvins, data.Main.MaxKelvins)
	fmt.Printf("Pressure: %dkPa\n", data.Main.Pressure)
	fmt.Printf("Humidity: %d%%\n", data.Main.Humidity)
	fmt.Printf("City: %s", data.Name)
}

func main() {
	err := os.RemoveAll("src/weatherTracker/data")
	if err != nil {
		log.Fatal(err)
	}
	os.MkdirAll("src/weatherTracker/data", os.ModeDir)

	fmt.Print("Enter city to check weather: ")
	fmt.Scan(&city)

	for longpoll := 0; longpoll < maxPolling; longpoll++ {
		getJSON(city)
		prevdata = data
		data = getJSON(city)

		display(data)

		if prevdata != nil && data != nil {
			if !reflect.DeepEqual(data, prevdata) {
				f, err := os.Open("src/weatherTracker/assets/beep.wav")
				if err != nil {
					log.Fatal(err)
				}
				streamer, format, err := wav.Decode(f)
				if err != nil {
					log.Fatal(err)
				}
				defer streamer.Close()

				speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
				speaker.Play(streamer)
			}
		}

		file, _ := json.MarshalIndent(data, "", " ")
		err := ioutil.WriteFile("src/weatherTracker/data/"+strconv.FormatInt(time.Now().UnixNano(), 10)+"data.json", file, 0644)

		if err != nil {
			panic(err)
		}

		time.Sleep(time.Second * delaySec)
	}
}
