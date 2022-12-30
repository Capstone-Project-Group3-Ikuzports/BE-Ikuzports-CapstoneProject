package service

import (
	"errors"
	"ikuzports/features/event"
	"ikuzports/features/participant"
	"ikuzports/utils/helper"

	"github.com/labstack/gommon/log"
)

type participantService struct {
	participantRepository participant.RepositoryInterface
	eventRepository       event.RepositoryInterface
}

func New(repo participant.RepositoryInterface, eventRepo event.RepositoryInterface) participant.ServiceInterface {
	return &participantService{
		participantRepository: repo,
		eventRepository:       eventRepo,
	}
}

func (service *participantService) Create(data participant.ParticipantCore) (err error) {
	dataParticipant, _ := service.participantRepository.FindMember(data.EventID, data.UserID)

	if dataParticipant.EventID == data.EventID && dataParticipant.UserID == data.UserID {
		return errors.New(" failed to join, you are already in this event")
	}

	dataEvent, errEvent := service.eventRepository.GetByID(data.EventID)
	if errEvent != nil {
		log.Error(err.Error())
		return helper.ServiceErrorMsg(errEvent)
	}

	if dataEvent.TotalParticipant+1 >= dataEvent.MaximumPeople {
		_, err = service.participantRepository.UpdateStatus(data)
		if err != nil {
			return errors.New("failed to update status. error query")
		}
	}

	if dataEvent.TotalParticipant >= dataEvent.MaximumPeople {
		return errors.New("cannot post data. maximum people reached")
	} else {
		data.Status = "Participant"

		_, err = service.participantRepository.Create(data)
		if err != nil {
			return errors.New("failed to insert data. error query")
		}

		_, errUpdate := service.participantRepository.UpdateParticipant(data)
		if errUpdate != nil {
			return errors.New("failed to update total participant data. error query")
		}
	}

	return nil
}
