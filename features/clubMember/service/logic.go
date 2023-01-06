package service

import (
	"errors"
	"ikuzports/features/club"
	"ikuzports/features/clubMember"
	"ikuzports/utils/helper"
	"ikuzports/utils/thirdparty"
	"log"
)

type clubMemberService struct {
	clubMemberRepository clubMember.RepositoryInterface
	clubRepository       club.RepositoryInterface
	// validate       *validator.Validate
}

func New(repo clubMember.RepositoryInterface, repoClub club.RepositoryInterface) clubMember.ServiceInterface {
	return &clubMemberService{
		clubMemberRepository: repo,
		clubRepository:       repoClub,
		// validate:       validator.New(),
	}
}

// Create implements clubMember.ServiceInterface
func (service *clubMemberService) Create(input clubMember.Core) error {
	input.Status = "Requested"
	dataMember, errCr := service.clubMemberRepository.FindMember(int(input.ClubID), int(input.UserID))
	if errCr != nil {
		errCreate := service.clubMemberRepository.Create(input)
		if errCreate != nil {
			return errors.New("failed to insert data, error query")
		}
	}
	if dataMember.ClubID == input.ClubID && dataMember.UserID == input.UserID && dataMember.DeletedAt == input.DeletedAt {
		if dataMember.Status == "Requested" {
			return errors.New(" failed to join, you are already register in this club. Please wait until the owner approve your request")
		} else {
			return errors.New(" failed to join, you are already join in this club")
		}
	}
	// errPut := service.clubMemberRepository.UpdateMember(int(input.ClubID))
	// if errPut != nil {
	// 	return errors.New("failed to update joined member, error query")
	// }
	return nil
}

// GetAll implements clubMember.ServiceInterface
func (service *clubMemberService) GetAll(queryStatus string) (data []clubMember.Core, err error) {
	if queryStatus == "" {
		data, err = service.clubMemberRepository.GetAll()
	} else {
		data, err = service.clubMemberRepository.GetAllByStatus(queryStatus)
	}
	if err != nil {
		helper.LogDebug(err)
		return nil, helper.ServiceErrorMsg(err)
	}

	return data, nil
}

// GetById implements clubMember.ServiceInterface
func (service *clubMemberService) GetById(id int) (data clubMember.Core, err error) {
	data, err = service.clubMemberRepository.GetById(id)
	if err != nil {
		return clubMember.Core{}, helper.ServiceErrorMsg(err)
	}
	return data, err
}

// Delete implements clubMember.ServiceInterface
func (service *clubMemberService) Delete(id int, userId int) error {
	data, err := service.clubMemberRepository.GetById(id)
	if err != nil {
		return helper.ServiceErrorMsg(err)
	}
	dataMember, errCr := service.clubRepository.GetStatus(int(data.ClubID), userId)
	if errCr != nil {
		return errors.New("error delete member. no data or you are not joined in this club")
	}

	if dataMember.Status != "Owner" {
		return errors.New("failed delete member, you are not the owner of the club")
	}

	errCreate := service.clubMemberRepository.Delete(id)
	if errCreate != nil {
		return errors.New("failed to delete data, error query")
	}
	errPut := service.clubMemberRepository.UpdateMember(int(data.ClubID))
	if errPut != nil {
		return errors.New("failed to update joined member, error query")
	}
	return nil
}

// Update implements clubMember.ServiceInterface
func (service *clubMemberService) Update(input clubMember.Core, id int, userId int) error {
	dataClub, errEvent := service.clubRepository.GetById(int(input.ClubID))
	if errEvent != nil {
		return helper.ServiceErrorMsg(errEvent)
	}

	if dataClub.JoinedMember >= dataClub.MemberTotal {
		return errors.New(" cannot add new member. the club is full")
	}

	input.Status = "Member"
	data, err := service.clubMemberRepository.GetById(id)
	if err != nil {
		return helper.ServiceErrorMsg(err)
	}
	dataMember, errCr := service.clubRepository.GetStatus(int(input.ClubID), userId)
	if errCr != nil {
		return errors.New("error update member. no data or you are not joined in this club")
	}
	// if data.UserID != uint(userId) {
	if dataMember.Status != "Owner" {
		return errors.New("failed update member, you are not the owner of the club")
	}
	// }
	errCreate := service.clubMemberRepository.Update(input, id)
	if errCreate != nil {
		return errors.New("failed to update data, error query")
	}

	errPut := service.clubMemberRepository.UpdateMember(int(data.ClubID))
	if errPut != nil {
		return errors.New("failed to update joined member, error query")
	}

	dataEmail := struct {
		UserName     string
		Club         string
		Status       string
		City         string
		JoinedMember uint
	}{
		UserName:     data.User.Name,
		Club:         dataClub.Name,
		Status:       dataMember.Status,
		City:         dataClub.City,
		JoinedMember: dataClub.JoinedMember,
	}
	// emailTo := dataMember.User.Email
	emailTo := data.User.Email
	log.Println("emailTo", emailTo)
	log.Println("===============================")

	// errMail := helper.SendGmailNotif(dataMember.User.Name, emailTo, dataClub.Name, dataMember.Status, dataClub.City, dataClub.JoinedMember)
	// if errMail != nil {
	// 	log.Println(errMail, "Pengiriman Email Gagal")
	// }
	errMail := thirdparty.SendEmailSMTPCheckup([]string{emailTo}, dataEmail, "notif_temp.html") //send mail
	if errMail != nil {
		log.Println(errMail, "Pengiriman Email Gagal")
	}

	return nil
}
