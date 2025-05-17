package weather_test

import (
	"go/BasGo/geo"
	"go/BasGo/weather"
	"strings"
	"testing"
)

func TestGetWeather(t *testing.T) {
	expected := "Paris"
	geoData := geo.DataGeo{
		City: expected,
	}
	format := 3

	result := weather.GetWeather(geoData, format)
	if !strings.Contains(result, expected) {
		t.Errorf("Ожидалось %v, получено %v", expected, result)
	}
}
