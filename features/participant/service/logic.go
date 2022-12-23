package service

import (
	"errors"
	"ikuzports/features/participant"
)

type participantService struct {
	participantRepository participant.RepositoryInterface
}

func New(repo participant.RepositoryInterface) participant.ServiceInterface {
	return &participantService{
		participantRepository: repo,
	}
}

func (service *participantService) Create(data participant.ParticipantCore) (err error) {
	data.Status = "Participant"

	_, err = service.participantRepository.Create(data)
	if err != nil {
		return errors.New("failed to insert data. error query")
	}

	_, errUpdate := service.participantRepository.UpdateParticipant(data)
	if errUpdate != nil {
		return errors.New("failed to update total participant data. error query")
	}

	return nil
}
