package participant

type ParticipantCore struct {
	ID      uint
	UserID  int
	EventID int
	Status  string
}

type RepositoryInterface interface {
	Create(data ParticipantCore) (row int, err error)
	UpdateParticipant(data ParticipantCore) (row int, err error)
	UpdateStatus(data ParticipantCore) (row int, err error)
	FindMember(eventID, userID int) (data ParticipantCore, err error)
}

type ServiceInterface interface {
	Create(data ParticipantCore) (err error)
}
