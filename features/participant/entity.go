package participant

type ParticipantCore struct {
	ID      int
	UserID  int
	EventID int
	Status  string
}

type RepositoryInterface interface {
	Create(data ParticipantCore) (row int, err error)
	UpdateParticipant(data ParticipantCore) (row int, err error)
	UpdateStatus(data ParticipantCore) (row int, err error)
}

type ServiceInterface interface {
	Create(data ParticipantCore) (err error)
}
