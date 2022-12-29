package delivery

import (
	"ikuzports/features/event"
	"time"
)

type EventResponse struct {
	ID               uint      `json:"id"`
	Name             string    `json:"name"`
	Address          string    `json:"address"`
	City             string    `json:"city"`
	OrganizerName    string    `json:"organizer_name"`
	CategoryName     string    `json:"category_name"`
	StartDate        time.Time `json:"start_date"`
	EndDate          time.Time `json:"end_date"`
	TotalParticipant uint      `json:"total_participant"`
	ImageEvent       string    `json:"image_event"`
	Status           string    `json:"status"`
	MaximumPeople    int       `json:"maximum_people"`
	Description      string    `json:"description"`
}

func fromCore(dataCore event.EventCore) EventResponse {
	return EventResponse{
		ID:               dataCore.ID,
		Name:             dataCore.Name,
		Address:          dataCore.Address,
		City:             dataCore.City,
		OrganizerName:    dataCore.User.Name,
		CategoryName:     dataCore.Category.Name,
		StartDate:        dataCore.StartDate,
		EndDate:          dataCore.EndDate,
		TotalParticipant: uint(dataCore.TotalParticipant),
		ImageEvent:       dataCore.ImageEvent,
		Status:           dataCore.Status,
		MaximumPeople:    dataCore.MaximumPeople,
		Description:      dataCore.Description,
	}
}

func fromCoreList(dataCore []event.EventCore) []EventResponse {
	var dataResponse []EventResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
