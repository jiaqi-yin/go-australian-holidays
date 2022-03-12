package services

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/jiaqi-yin/go-australian-holidays/domain"
	"github.com/jiaqi-yin/go-australian-holidays/storage"
)

type HolidayService struct {
	Storage storage.Storage
}

func (hs *HolidayService) GetAndSave(endpoint string) {
	response, err := http.Get(endpoint)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject domain.Response
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		log.Fatal(err)
	}

	hs.Storage.Save(responseObject.Result.Holidays)
}

func (hs *HolidayService) LoadFromFile() []domain.Holiday {
	return hs.Storage.Load()
}

func (hs *HolidayService) IsHoliday(state string, date string) domain.HolidayCheckResponse {
	holidays := hs.LoadFromFile()
	var response domain.HolidayCheckResponse
	for _, holiday := range holidays {
		if holiday.State == state && holiday.Date == date {
			response.IsHoliday = true
			response.Holiday = holiday
			break
		}
	}
	return response
}

func NewHolidayService(storage storage.Storage) *HolidayService {
	return &HolidayService{
		Storage: storage,
	}
}
