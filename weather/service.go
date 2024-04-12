package weather

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func getWeatherForLatLng(lat string, lng string) (GetWeatherResponse, error) {
	url := fmt.Sprintf("%s/data/%s/weather?lat=%s&lon=%s&appid=%s&units=imperial", os.Getenv("OPEN_WEATHER_API_URL"), os.Getenv("OPEN_WEATHER_API_VERSION"), lat, lng, os.Getenv("OPEN_WEATHER_API_KEY"))

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return GetWeatherResponse{}, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return GetWeatherResponse{}, err
	}

	if res.StatusCode != 200 {
		return GetWeatherResponse{}, errors.New("openweather api returned " + res.Status)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return GetWeatherResponse{}, err
	}

	var apiResponse GetOpenWeatherResponse
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		return GetWeatherResponse{}, err
	}

	return processWeatherResponse(apiResponse), nil
}

func processWeatherResponse(obj GetOpenWeatherResponse) GetWeatherResponse {
	returnObj := GetWeatherResponse{
		Condition:   "",
		Temperature: "",
	}

	if len(obj.Weather) < 1 {
		returnObj.Condition = "Weather Status Not Available"
	} else {
		returnObj.Condition = obj.Weather[0].Description
	}

	if obj.Main.Temp < 69.0 {
		returnObj.Temperature = "cold"
	} else if obj.Main.Temp > 80.0 {
		returnObj.Temperature = "hot"
	} else {
		returnObj.Temperature = "moderate"
	}

	return returnObj
}
