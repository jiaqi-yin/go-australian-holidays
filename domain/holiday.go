package domain

import "fmt"

type Response struct {
	Success bool   `json:"success"`
	Result  Result `json:"result"`
}

type Result struct {
	Limit    int       `json:"limit"`
	Holidays []Holiday `json:"records"`
}

type Holiday struct {
	Date  string `json:"Date"`
	Name  string `json:"Holiday Name"`
	State string `json:"Jurisdiction"`
}

func (h *Holiday) ToString() string {
	return fmt.Sprintf("%v,%v,%v", h.Date, h.Name, h.State)
}

func NewHoliday(date, name, info, state string) *Holiday {
	return &Holiday{
		Date:  date,
		Name:  name,
		State: state,
	}
}
