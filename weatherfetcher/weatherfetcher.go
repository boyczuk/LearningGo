package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type WeatherResponse struct {
	Main struct {
		Temperature float64 `json:"temp"`
		FeelsLike   float64 `json:"feels_like"`
		Humidity    int     `json:"humidity"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
}

func getWeather(city string) {
	apiKey := "f2f1d3faa0a6d3c0de4f69e96236d670"
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	body, _ := io.ReadAll(resp.Body)

	var weather WeatherResponse
	json.Unmarshal(body, &weather)

	tempCelsius := strconv.FormatFloat(weather.Main.Temperature-273.15, 'f', 2, 64)
	feelsLikeCelsius := strconv.FormatFloat(weather.Main.FeelsLike-273.15, 'f', 2, 64)

	fmt.Println("Temperature: " + tempCelsius + "°C")
	fmt.Println("Feels like temperature: " + feelsLikeCelsius + "°C")
	fmt.Println("Humidity: " + strconv.Itoa(weather.Main.Humidity) + "%")
	fmt.Println("In the sky: " + weather.Weather[0].Description)

}

func main() {
	getWeather("Toronto")
}
