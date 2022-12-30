package repository

import (
	"errors"
	"ikuzports/features/event"

	"gorm.io/gorm"
)

type eventRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) event.RepositoryInterface {
	return &eventRepository{
		db: db,
	}
}

func (repo *eventRepository) Create(input event.EventCore) (row int, err error) {
	eventGorm := fromCore(input)
	tx := repo.db.Create(&eventGorm)
	if tx.Error != nil {
		return -1, tx.Error
	}

	// var id int
	// ty := repo.db.Raw("SELECT LAST_INSERT_ID()").Scan(&id)
	// if ty.Error != nil {
	// 	return -1, errors.New("failed create data")
	// }

	// tx2 := repo.db.Exec(`INSERT INTO event_participants (created_at, updated_at, user_id, event_id, status)
	// VALUES(current_timestamp, current_timestamp, ?, ?, "Organizer")`, eventGorm.UserID, id)
	// if tx2.Error != nil {
	// 	return -1, errors.New("failed create data")
	// }

	// tz := repo.db.Exec("UPDATE events SET total_participant = (SELECT COUNT(user_id) FROM event_participants WHERE event_id = ?) WHERE id = ?", id, id)
	// if tz.Error != nil {
	// 	return -1, errors.New("failed create data")
	// }

	if tx.RowsAffected == 0 {
		return 0, errors.New("insert failed")
	}

	return int(tx.RowsAffected), nil
}

func (repo *eventRepository) GetLastID() (id int, err error) {
	var lastID int
	tx := repo.db.Raw("SELECT LAST_INSERT_ID()").Scan(&lastID)
	if tx.Error != nil {
		return -1, errors.New("failed create data")
	}
	return lastID, nil
}

func (repo *eventRepository) UpdateTotal(id int) (rows int, err error) {
	tx := repo.db.Exec("UPDATE events SET total_participant = (SELECT COUNT(user_id) FROM event_participants WHERE event_id = ?) WHERE id = ?", id, id)
	if tx.Error != nil {
		return -1, errors.New("failed update total_participant data")
	}

	if tx.RowsAffected == 0 {
		return 0, errors.New("failed update total_participant data")
	}

	return int(tx.RowsAffected), nil
}

func (repo *eventRepository) GetDate() (date []event.EventCore, err error) {
	var dates []event.EventCore
	tx := repo.db.Table("events").Select("id", "start_date", "end_date").Find(&dates)
	if tx.Error != nil {
		return nil, errors.New("failed read data")
	}

	if tx.RowsAffected == 0 {
		return nil, nil
	}

	return dates, nil
}

func (repo *eventRepository) UpdateStatus(id int, status string) (rows int, err error) {
	tx := repo.db.Exec("UPDATE events SET status = ? WHERE id = ?", status, id)
	if tx.Error != nil {
		return -1, errors.New("failed update data")
	}

	if tx.RowsAffected == 0 {
		return 0, errors.New("failed update data")
	}

	return int(tx.RowsAffected), nil
}

func (repo *eventRepository) GetAllFilter(limit, offset, queryCategoryID int, queryCity, queryStatus string) (data []event.EventCore, err error) {
	var event []Event
	tx := repo.db.Preload("User").Preload("Category").Where(&Event{CategoryID: uint(queryCategoryID), City: queryCity, Status: queryStatus}).Order("updated_at desc").Limit(limit).Offset(offset).Find(&event)

	if tx.Error != nil {
		return nil, tx.Error
	}

	dataCore := toCoreList(event)

	return dataCore, nil

}

func (repo *eventRepository) GetAll(limit, offset int) (data []event.EventCore, err error) {
	var event []Event

	tx := repo.db.Preload("User").Preload("Category").Order("updated_at desc").Limit(limit).Offset(offset).Find(&event)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var dataCore = toCoreList(event)

	return dataCore, nil
}

func (repo *eventRepository) GetByID(id int) (data event.EventCore, err error) {
	var event Event

	tx := repo.db.Preload("User").Preload("Category").First(&event, id)
	if tx.Error != nil {
		return data, tx.Error
	}

	var dataCore = event.toCore()

	return dataCore, nil
}

func (repo *eventRepository) Delete(id int) (row int, err error) {
	var event Event
	tx := repo.db.Delete(&event, id)
	if tx.Error != nil {
		return -1, tx.Error
	}

	if tx.RowsAffected == 0 {
		return 0, errors.New("delete failed")
	}

	return int(tx.RowsAffected), nil
}
