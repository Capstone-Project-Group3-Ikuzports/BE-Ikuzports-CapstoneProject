package service

import (
	"errors"
	"ikuzports/features/event"
	"ikuzports/features/participant"
	"ikuzports/utils/helper"
	"time"

	"github.com/go-playground/validator/v10"
)

type eventService struct {
	eventRepository       event.RepositoryInterface
	participantRepository participant.RepositoryInterface
	validate              *validator.Validate
}

func New(repo event.RepositoryInterface, parRepo participant.RepositoryInterface) event.ServiceInterface {
	return &eventService{
		eventRepository:       repo,
		participantRepository: parRepo,
		validate:              validator.New(),
	}
}

func (service *eventService) Create(input event.EventCore) (err error) {
	helper.LogDebug(input)
	if errValidate := service.validate.Struct(input); errValidate != nil {
		return errValidate
	}
	if input.EndDate.Unix() < input.StartDate.Unix() {
		return errors.New("error input date")
	}

	input.Status = "Available"

	_, err = service.eventRepository.Create(input)
	if err != nil {
		return errors.New("failed to insert data. error query")
	}

	eventID, errGet := service.eventRepository.GetLastID()
	if errGet != nil {
		return errors.New("error get last ID. error query")
	}

	dataParticipant := participant.ParticipantCore{
		UserID:  input.UserID,
		EventID: eventID,
		Status:  "Organizer",
	}
	_, errCr := service.participantRepository.Create(dataParticipant)
	if errCr != nil {
		return errors.New("error create participant. error query")
	}

	_, errUpdate := service.eventRepository.UpdateTotal(eventID)
	if errUpdate != nil {
		return errors.New("error update total. error query")
	}

	return nil
}

func (service *eventService) GetAll(queryCategoryID, queryPage int, queryCity, queryStatus string) (data []event.EventCore, err error) {
	limit := 15
	offset := (queryPage - 1) * limit

	dataDate, errDate := service.eventRepository.GetDate()
	if errDate != nil {
		return nil, errDate
	}

	for _, v := range dataDate {
		if v.StartDate.Unix() < time.Now().Unix() || v.EndDate.Unix() < time.Now().Unix() {
			service.eventRepository.UpdateStatus(int(v.ID), "Not Available")
		}
	}

	if queryCategoryID == 0 && queryCity == "" && queryStatus == "" {
		data, err = service.eventRepository.GetAll(limit, offset)
	} else {
		data, err = service.eventRepository.GetAllFilter(limit, offset, queryCategoryID, queryCity, queryStatus)
	}

	if err != nil {
		helper.LogDebug(err)
		return nil, err
	}

	if len(data) == 0 {
		helper.LogDebug("Get data success. No data.")
		return nil, errors.New("Get data success. No data.")
	}

	return data, nil
}

func (service *eventService) GetByID(id int) (data event.EventCore, err error) {
	data, err = service.eventRepository.GetByID(id)
	if err != nil {
		return event.EventCore{}, err
	}

	return data, nil
}

func (service *eventService) Delete(id int) (err error) {
	_, err = service.eventRepository.Delete(id)
	if err != nil {
		return errors.New("error delete data")
	}

	return nil
}

// func (service *eventService) UpdateStatus() (err error) {

// 	dates, errDates := service.eventRepository.GetDate()
// 	if errDates != nil {
// 		helper.LogDebug(errDates)
// 		return errDates
// 	}

// 	for _, v := range dates {
// 		if v.StartDate.Unix() < time.Now().Unix() || v.EndDate.Unix() < time.Now().Unix() {
// 			service.eventRepository.UpdateStatus(int(v.ID), "Not Available")
// 		}
// 	}

// 	return nil
// }
