package domain

type HolidayCheckRequest struct {
	State string `json:"state" binding:"required"`
	Date  string `json:"date" binding:"required"`
}
