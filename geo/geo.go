package geo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type DataGeo struct {
	City string `json:"city"`
}
type CityPopulationResponce struct {
	Error bool `json:"error"`
}

func GetMyLocation(city string) (*DataGeo, error) {
	if city != "" {
		isCity := checkCity(city)
		if !isCity {
			return nil, errors.New("NOCITY")
		}
		return &DataGeo{
			City: city,
		}, nil
	}
	resp, err := http.Get("https://ipapi.co/json/")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		return nil, errors.New("NOT200: " + resp.Status + " " + string(body))
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var geo DataGeo
	err = json.Unmarshal(body, &geo)
	if err != nil {
		return nil, err
	}
	return &geo, nil
}

func checkCity(city string) bool {
	postBody, _ := json.Marshal(map[string]string{
		"city": city,
	})
	resp, err := http.Post("https://countriesnow.space/api/v0.1/countries/population/cities", "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		fmt.Println(errors.New("NOT200: " + resp.Status + " " + string(body)))
		return false
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false
	}
	var populationResponce CityPopulationResponce
	err = json.Unmarshal(body, &populationResponce)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return !populationResponce.Error
}
