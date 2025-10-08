package weather

import (
	"fmt"
	"go/BasGo/geo"
	"io"
	"net/http"
	"net/url"
)

func GetWeather(geo geo.DataGeo, format int) string {
	baseUrl, err := url.Parse("https://wttr.in/" + geo.City)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	params := url.Values{}
	params.Add("format", fmt.Sprint(format))
	baseUrl.RawQuery = params.Encode()
	resp, err := http.Get(baseUrl.String())
	if err != nil {
		defer resp.Body.Close()
		fmt.Println(err.Error())
		return ""
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return fmt.Sprintf("Location: %s\nWeather: %s", geo.City, string(body))
}
