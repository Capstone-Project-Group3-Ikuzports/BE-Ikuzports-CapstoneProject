package repository

import (
	"ikuzports/features/participant"

	"gorm.io/gorm"
)

type EventParticipant struct {
	gorm.Model
	UserID  int
	EventID int
	Status  string
}

func fromCore(dataModel participant.ParticipantCore) EventParticipant {
	eventGorm := EventParticipant{
		UserID:  dataModel.UserID,
		EventID: dataModel.EventID,
		Status:  dataModel.Status,
	}
	return eventGorm
}

// func (dataModel *EventParticipant) toCore() participant.ParticipantCore {
// 	return participant.ParticipantCore{
// 		ID:      dataModel.ID,
// 		UserID:  dataModel.UserID,
// 		EventID: dataModel.EventID,
// 		Status:  dataModel.Status,
// 	}
// }

// func toCoreList(dataModel []EventParticipant) []participant.ParticipantCore {
// 	var dataCore []participant.ParticipantCore
// 	for _, v := range dataModel {
// 		dataCore = append(dataCore, v.toCore())
// 	}
// 	return dataCore
// }
