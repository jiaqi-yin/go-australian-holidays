package main

import (
	"github.com/jiaqi-yin/go-australian-holidays/app"
	"github.com/jiaqi-yin/go-australian-holidays/services"
	"github.com/jiaqi-yin/go-australian-holidays/storage"
)

const (
	FILE = "holiday.csv"
)

func main() {
	app := app.App{
		HolidayService: services.NewHolidayService(storage.NewFileSystem(FILE)),
	}
	app.Init()
	app.Run()
}
