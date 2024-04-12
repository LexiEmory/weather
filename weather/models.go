package weather

import "github.com/gin-gonic/gin"

type GetWeatherResponse struct {
	Condition   string `json:"condition"`
	Temperature string `json:"temperature"`
}

type GetOpenWeatherResponse struct {
	Weather []Weather `json:"weather"`
	Main    Main      `json:"main"`
}

type Weather struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Main struct {
	Temp float64 `json:"temp"`
}

func JsonError(e error) gin.H {
	return gin.H{
		"error": e.Error(),
	}
}
