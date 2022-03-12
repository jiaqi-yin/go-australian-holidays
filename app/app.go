package app

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jiaqi-yin/go-australian-holidays/domain"
	"github.com/jiaqi-yin/go-australian-holidays/services"
)

const (
	ENDPOINT = "https://data.gov.au/data/api/3/action/datastore_search?resource_id=33673aca-0857-42e5-b8f0-9981b4755686&limit=500"
)

type App struct {
	Router         *gin.Engine
	HolidayService *services.HolidayService
}

func (app *App) Init() {
	app.Router = gin.Default()
	app.initializeRoutes()
	app.HolidayService.GetAndSave(ENDPOINT)
}

func (app *App) initializeRoutes() {
	v1 := app.Router.Group("/v1")
	{
		v1.GET("/ping", app.ping)
		v1.GET("/is_holiday", app.checkHoliday)
	}
}

func (app *App) Run() {
	log.Fatal(app.Router.Run())
}

func (app *App) ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
	return
}

func (app *App) checkHoliday(c *gin.Context) {
	var request domain.HolidayCheckRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := app.HolidayService.IsHoliday(request.State, request.Date)
	c.JSON(http.StatusOK, response)
	return
}
