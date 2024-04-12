package weather

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Route(router *gin.Engine) {
	weatherRouter := router.Group("/w")

	weatherRouter.GET("/", GetWeatherRoute())
}

func GetWeatherRoute() func(w *gin.Context) {
	return func(w *gin.Context) {
		lng := w.Query("lng")
		lat := w.Query("lat")

		if lng == "" || lat == "" {
			w.AbortWithStatusJSON(http.StatusBadRequest, JsonError(errors.New("lat or lng is missing")))
			return
		}

		res, err := getWeatherForLatLng(lat, lng)
		if err != nil {
			w.AbortWithStatusJSON(http.StatusInternalServerError, JsonError(err))
			return
		}

		w.JSON(http.StatusOK, res)
	}
}
