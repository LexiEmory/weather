package weather

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jarcoal/httpmock"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func SetUpTestingEnv(t *testing.T) *gin.Engine {
	r := gin.Default()

	godotenv.Load("../.env.testing")

	Route(r)

	return r
}

func Test_Get_Weather(t *testing.T) {
	r := SetUpTestingEnv(t)

	url := fmt.Sprintf("%s/data/%s/weather?lat=%s&lon=%s&appid=%s&units=imperial", os.Getenv("OPEN_WEATHER_API_URL"), os.Getenv("OPEN_WEATHER_API_VERSION"), "0.0", "0.0", os.Getenv("OPEN_WEATHER_API_KEY"))
	openWeatherResponse := getMockOpenWeatherResponse()
	str, _ := json.Marshal(openWeatherResponse)
	stringResponse := httpmock.NewStringResponder(200, string(str))
	httpmock.RegisterResponder("GET", url, stringResponse)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	emptyRequest := ``
	jsonValue := io.NopCloser(strings.NewReader(emptyRequest))

	req, _ := http.NewRequest("GET", "/w/?lat=0.0&lng=0.0", jsonValue)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func Test_Get_Weather_Bad_Params(t *testing.T) {
	r := SetUpTestingEnv(t)

	url := fmt.Sprintf("%s/data/%s/weather?lat=%s&lon=%s&appid=%s&units=imperial", os.Getenv("OPEN_WEATHER_API_URL"), os.Getenv("OPEN_WEATHER_API_VERSION"), "0.0", "0.0", os.Getenv("OPEN_WEATHER_API_KEY"))
	openWeatherResponse := getMockOpenWeatherResponse()
	str, _ := json.Marshal(openWeatherResponse)
	stringResponse := httpmock.NewStringResponder(200, string(str))
	httpmock.RegisterResponder("GET", url, stringResponse)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	emptyRequest := ``
	jsonValue := io.NopCloser(strings.NewReader(emptyRequest))

	req, _ := http.NewRequest("GET", "/w/", jsonValue)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func Test_Get_Weather_Bad_OpenWeather_Response(t *testing.T) {
	r := SetUpTestingEnv(t)

	url := fmt.Sprintf("%s/data/%s/weather?lat=%s&lon=%s&appid=%s&units=imperial", os.Getenv("OPEN_WEATHER_API_URL"), os.Getenv("OPEN_WEATHER_API_VERSION"), "0.0", "0.0", os.Getenv("OPEN_WEATHER_API_KEY"))
	openWeatherResponse := getMockOpenWeatherResponse()
	str, _ := json.Marshal(openWeatherResponse)
	// sim a server failure
	stringResponse := httpmock.NewStringResponder(500, string(str))
	httpmock.RegisterResponder("GET", url, stringResponse)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	emptyRequest := ``
	jsonValue := io.NopCloser(strings.NewReader(emptyRequest))

	req, _ := http.NewRequest("GET", "/w/?lat=0.0&lng=0.0", jsonValue)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func Test_Response_Gen(t *testing.T) {
	SetUpTestingEnv(t)

	coldResPrebuild := getMockOpenWeatherResponseWithTemp(50)
	coldPost := processWeatherResponse(coldResPrebuild)
	moderateResPrebuild := getMockOpenWeatherResponseWithTemp(70)
	moderatePost := processWeatherResponse(moderateResPrebuild)
	hotResPrebuild := getMockOpenWeatherResponseWithTemp(100)
	hotPost := processWeatherResponse(hotResPrebuild)

	assert.Equal(t, coldPost.Temperature, "cold")
	assert.Equal(t, hotPost.Temperature, "hot")
	assert.Equal(t, moderatePost.Temperature, "moderate")
	assert.Equal(t, moderatePost.Condition, "Weather Status Not Available")
}

func getMockOpenWeatherResponse() GetOpenWeatherResponse {
	return GetOpenWeatherResponse{
		Weather: []Weather{
			{
				ID:          0,
				Main:        "Cloudy",
				Description: "testing desc",
				Icon:        "b11",
			},
		},
		Main: Main{
			Temp: 69.0,
		},
	}
}

func getMockOpenWeatherResponseWithTemp(temp float64) GetOpenWeatherResponse {
	return GetOpenWeatherResponse{
		Weather: []Weather{},
		Main: Main{
			Temp: temp,
		},
	}
}
