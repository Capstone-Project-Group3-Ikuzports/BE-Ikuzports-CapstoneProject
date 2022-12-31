package repository

import (
	"errors"
	"ikuzports/features/chat"
	"ikuzports/features/club"
	"ikuzports/features/clubActivity"
	"ikuzports/features/galery"

	_members "ikuzports/features/clubMember"

	"gorm.io/gorm"
)

type clubRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) club.RepositoryInterface {
	return &clubRepository{
		db: db,
	}
}

// Create implements club.RepositoryInterface
func (repo *clubRepository) Create(input club.Core, id int) error {
	clubGorm := fromCore(input)
	tx := repo.db.Create(&clubGorm) // proses insert data
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failed")
	}

	// var idClub int
	// ty := repo.db.Raw("SELECT LAST_INSERT_ID()").Scan(&idClub)
	// if ty.Error != nil {
	// 	return errors.New("failed create data")
	// }
	// status := "Owner"
	// tz := repo.db.Exec("insert into club_members (user_id, club_id, status) Values(?, ?, ?)", id, idClub, status)
	// if tz.Error != nil {
	// 	return errors.New("failed create data")
	// }
	return nil
}

// GetLastID implements club.RepositoryInterface
func (repo *clubRepository) GetLastID() (id int, err error) {
	var lastID int
	tx := repo.db.Raw("SELECT LAST_INSERT_ID()").Scan(&lastID)
	if tx.Error != nil {
		return -1, errors.New("failed create data")
	}
	return lastID, nil
}

// UpdateTotal implements club.RepositoryInterface
func (repo *clubRepository) UpdateMember(id int) (rows int, err error) {
	tx := repo.db.Exec("UPDATE clubs SET joined_member = (SELECT COUNT(user_id) FROM club_members WHERE club_id = ?) WHERE id = ?", id, id)
	if tx.Error != nil {
		return -1, errors.New("failed update total_participant data")
	}

	if tx.RowsAffected == 0 {
		return 0, errors.New("failed update total_participant data")
	}

	return int(tx.RowsAffected), nil
}

// GetAll implements club.RepositoryInterface
func (repo *clubRepository) GetAll(offset, limit int) (data []club.Core, page int, err error) {
	var club []Club
	var jumlahData int64

	ty := repo.db.Preload("Category").Model(&Club{}).Order("created_at desc").Count(&jumlahData)
	if ty.Error != nil {
		return nil, 0, ty.Error
	}
	tx := repo.db.Preload("Category").Offset(offset).Limit(limit).Order("created_at desc").Find(&club)
	if tx.Error != nil {
		return nil, 0, tx.Error
	}

	var dataCore = toCoreList(club)

	return dataCore, int(jumlahData), nil
}

// GetAllWithSearch implements club.RepositoryInterface
func (repo *clubRepository) GetAllWithSearch(queryName, queryCity string, queryCategoryID, offset, limit int) (data []club.Core, page int, err error) {
	var club []Club

	var jumlahData int64
	ty := repo.db.Preload("Category").Model(&Club{}).Order("created_at desc").Count(&jumlahData)
	if ty.Error != nil {
		return nil, 0, ty.Error
	}

	tx := repo.db.Preload("Category").Where("name LIKE ?", "%"+queryName+"%").Where(&Club{CategoryID: uint(queryCategoryID), City: queryCity}).Offset(offset).Limit(limit).Order("created_at desc").Find(&club)
	if tx.Error != nil {
		return nil, 0, tx.Error
	}
	var dataCore = toCoreList(club)

	return dataCore, int(jumlahData), nil
}

// GetById implements club.RepositoryInterface
func (repo *clubRepository) GetById(id int) (data club.Core, err error) {
	var club Club

	tx := repo.db.Preload("Category").First(&club, id)

	if tx.Error != nil {
		return data, tx.Error
	}

	if tx.RowsAffected == 0 {
		return data, tx.Error
	}

	var dataCore = club.toCore()
	return dataCore, nil
}

// Update implements club.RepositoryInterface
func (repo *clubRepository) Update(input club.Core, id int) error {
	clubGorm := fromCore(input)
	tx := repo.db.Where("id= ?", id).Updates(clubGorm)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("update failed")
	}
	return nil
}

// GetStatus implements club.RepositoryInterface
func (repo *clubRepository) GetStatus(id int, userId int) (data club.Status, err error) {
	var member ClubMember

	tx := repo.db.Where("club_id = ?", id).Where("user_id = ?", userId).First(&member)
	if tx.Error != nil {
		return data, tx.Error
	}
	var dataCore = member.toCoreMember()

	return dataCore, nil
}

// Delete implements club.RepositoryInterface
func (repo *clubRepository) Delete(id int) error {
	var club Club
	tx := repo.db.Delete(&club, id) // proses delete
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("delete failed")
	}
	return nil
}

// GetMembers implements club.RepositoryInterface
func (repo *clubRepository) GetMembers(id int) (data []_members.Core, err error) {
	var members []ClubMember
	tx := repo.db.Where("club_id = ?", id).Preload("User").Find(&members)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var dataCore = toCoreMembersList(members)
	return dataCore, nil
}

// GetChats implements club.RepositoryInterface
func (repo *clubRepository) GetChats(id int) (data []chat.Core, err error) {
	var chats []Chat
	tx := repo.db.Where("club_id = ?", id).Preload("User").Find(&chats)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var dataCore = toCoreChatList(chats)
	return dataCore, nil
}

// GetGaleries implements club.RepositoryInterface
func (repo *clubRepository) GetGaleries(id int) (data []galery.Core, err error) {
	var gallery []Galery
	tx := repo.db.Where("club_id = ?", id).Find(&gallery)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var dataCore = toCoreGaleryList(gallery)
	return dataCore, nil
}

// GetActivities implements club.RepositoryInterface
func (repo *clubRepository) GetActivities(id int) (data []clubActivity.Core, err error) {
	var activities []ClubActivity
	tx := repo.db.Where("club_id = ?", id).Find(&activities)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var dataCore = toCoreActivityList(activities)
	return dataCore, nil
}
