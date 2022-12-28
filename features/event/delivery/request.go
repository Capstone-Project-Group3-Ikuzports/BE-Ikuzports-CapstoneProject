package delivery

import (
	"ikuzports/features/event"
	"time"
)

type EventRequest struct {
	Name          string `json:"name" form:"name"`
	Address       string `json:"address" form:"address"`
	City          string `json:"city" form:"city"`
	CategoryID    uint   `json:"category_id" form:"category_id"`
	StartDate     string `json:"start_date" form:"start_date"`
	EndDate       string `json:"end_date" form:"end_date"`
	ImageEvent    string `json:"image_event" form:"image_event"`
	MaximumPeople int    `json:"maximum_people" form:"maximum_people"`
}

func toCore(eventInput EventRequest) event.EventCore {
	eventCoreData := event.EventCore{
		Name:          eventInput.Name,
		Address:       eventInput.Address,
		City:          eventInput.City,
		CategoryID:    eventInput.CategoryID,
		StartDate:     timeParsing(eventInput.StartDate),
		EndDate:       timeParsing(eventInput.EndDate),
		ImageEvent:    eventInput.ImageEvent,
		MaximumPeople: eventInput.MaximumPeople,
	}
	return eventCoreData
}

func timeParsing(date string) (dateParsed time.Time) {
	layoutFormat := "2006-01-02"

	dateParsed, _ = time.Parse(layoutFormat, date)

	return dateParsed
}
