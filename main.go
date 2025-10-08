package main

import (
	"flag"
	"fmt"
	"go/BasGo/geo"
	"go/BasGo/weather"
)

func main() {
	city := flag.String("city", "", "Город пользователя")
	format := flag.Int("format", 1, "Формат вывода погоды")

	flag.Parse()

	fmt.Println(*city)
	geoData, err := geo.GetMyLocation(*city)
	if err != nil {
		fmt.Println("Ошибка геолокации", err.Error())
		return
	}
	weatherData := weather.GetWeather(*geoData, *format)
	fmt.Println(weatherData)
}
