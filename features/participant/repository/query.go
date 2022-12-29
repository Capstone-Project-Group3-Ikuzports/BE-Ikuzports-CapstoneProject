package repository

import (
	"errors"
	"ikuzports/features/participant"

	"gorm.io/gorm"
)

type participantRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) participant.RepositoryInterface {
	return &participantRepository{
		db: db,
	}
}

func (repo *participantRepository) Create(data participant.ParticipantCore) (row int, err error) {
	gormData := fromCore(data)

	tx := repo.db.Create(&gormData)
	if tx.Error != nil {
		return -1, errors.New("create data failed")
	}

	// tz := repo.db.Exec("UPDATE events SET total_participant = (SELECT COUNT(user_id) FROM event_participants WHERE event_id = ?) WHERE id = ?", gormData.EventID, gormData.EventID)
	// if tz.Error != nil {
	// 	return -1, errors.New("failed create data")
	// }

	if tx.RowsAffected == 0 {
		return 0, errors.New("create data failed")
	}

	return int(tx.RowsAffected), nil
}

func (repo *participantRepository) UpdateParticipant(data participant.ParticipantCore) (row int, err error) {
	gormData := fromCore(data)

	tx := repo.db.Exec("UPDATE events SET total_participant = (SELECT COUNT(user_id) FROM event_participants WHERE event_id = ?) WHERE id = ?", gormData.EventID, gormData.EventID)
	if tx.Error != nil {
		return -1, errors.New("failed create data")
	}

	if tx.RowsAffected == 0 {
		return 0, errors.New("create data failed")
	}

	return int(tx.RowsAffected), nil
}

func (repo *participantRepository) UpdateStatus(data participant.ParticipantCore) (row int, err error) {
	gormData := fromCore(data)

	tx := repo.db.Exec("UPDATE events SET status = ? WHERE id = ?", "Not Available", gormData.EventID)
	if tx.Error != nil {
		return -1, errors.New("failed create data")
	}

	if tx.RowsAffected == 0 {
		return 0, errors.New("create data failed")
	}

	return int(tx.RowsAffected), nil
}
