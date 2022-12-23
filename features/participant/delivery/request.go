package delivery

import "ikuzports/features/participant"

type ParticipantRequest struct {
	EventID int `json:"event_id" form:"event_id"`
}

func toCore(participantInput ParticipantRequest) participant.ParticipantCore {
	participantCoreData := participant.ParticipantCore{
		EventID: participantInput.EventID,
	}
	return participantCoreData
}
