package domain

type HolidayCheckResponse struct {
	IsHoliday bool    `json:"is_holiday"`
	Holiday   Holiday `json:"holiday"`
}
