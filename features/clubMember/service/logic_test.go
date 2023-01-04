package service

import (
	"errors"
	"ikuzports/features/club"
	"ikuzports/features/clubMember"
	"ikuzports/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	repo := new(mocks.ClubMemberRepo)
	repo2 := new(mocks.ClubRepo)

	t.Run("Failed create member", func(t *testing.T) {
		input := clubMember.Core{UserID: 1, ClubID: 2, Status: "Requested"}
		data := clubMember.Core{UserID: 1, ClubID: 2, Status: "Member"}
		repo.On("FindMember", 2, 1).Return(data, nil).Once()

		srv := New(repo, repo2)
		err := srv.Create(input)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed create member", func(t *testing.T) {
		input := clubMember.Core{UserID: 1, ClubID: 2, Status: "Requested"}
		data := clubMember.Core{UserID: 1, ClubID: 2, Status: "Requested"}
		repo.On("FindMember", 2, 1).Return(data, nil).Once()

		srv := New(repo, repo2)
		err := srv.Create(input)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("failed create member", func(t *testing.T) {
		input := clubMember.Core{UserID: 1, ClubID: 2, Status: "Requested"}

		repo.On("FindMember", 2, 1).Return(clubMember.Core{}, errors.New("failed"))
		repo.On("Create", input).Return(errors.New("failed")).Once()

		srv := New(repo, repo2)
		err := srv.Create(input)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("success create member", func(t *testing.T) {
		input := clubMember.Core{UserID: 1, ClubID: 2, Status: "Requested"}

		repo.On("FindMember", 2, 1).Return(clubMember.Core{}, errors.New("failed"))
		repo.On("Create", input).Return(nil).Once()

		srv := New(repo, repo2)
		err := srv.Create(input)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})
}

func TestGetAll(t *testing.T) {
	repo := new(mocks.ClubMemberRepo)
	repo2 := new(mocks.ClubRepo)

	t.Run("Succes get all member", func(t *testing.T) {
		input := []clubMember.Core{{UserID: 1, ClubID: 2, Status: "Requested"}}

		repo.On("GetAll").Return(input, nil).Once()

		srv := New(repo, repo2)
		res, err := srv.GetAll("")
		assert.NoError(t, err)
		assert.Equal(t, input[0].ClubID, res[0].ClubID)

		repo.AssertExpectations(t)
	})
	t.Run("Succes get all member", func(t *testing.T) {
		input := []clubMember.Core{{UserID: 1, ClubID: 2, Status: "Requested"}}

		repo.On("GetAllByStatus", "status").Return(input, nil).Once()

		srv := New(repo, repo2)
		res, err := srv.GetAll("status")
		assert.NoError(t, err)
		assert.Equal(t, input[0].ClubID, res[0].ClubID)

		repo.AssertExpectations(t)
	})

	t.Run("Failed get all member", func(t *testing.T) {

		repo.On("GetAll").Return(nil, errors.New("failed")).Once()

		srv := New(repo, repo2)
		_, err := srv.GetAll("")
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestGetById(t *testing.T) {
	repo := new(mocks.ClubMemberRepo)
	repo2 := new(mocks.ClubRepo)

	t.Run("Success get member", func(t *testing.T) {
		input := clubMember.Core{UserID: 1, ClubID: 2, Status: "Requested"}

		repo.On("GetById", 1).Return(input, nil).Once()
		srv := New(repo, repo2)
		response, err := srv.GetById(1)

		assert.NoError(t, err)
		assert.Equal(t, input.ClubID, response.ClubID)
		repo.AssertExpectations(t)
	})

	t.Run("Failed get member by id", func(t *testing.T) {
		inputRepo := clubMember.Core{UserID: 1, ClubID: 2, Status: "Requested"}

		repo.On("GetById", 1).Return(clubMember.Core{}, errors.New("failed")).Once()
		srv := New(repo, repo2)
		response, err := srv.GetById(1)

		assert.NotNil(t, err)
		assert.NotEqual(t, inputRepo.ClubID, response.ClubID)
		repo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	repo2 := new(mocks.ClubRepo)
	repo := new(mocks.ClubMemberRepo)
	id := 1
	userId := 2
	t.Run("Success delete members", func(t *testing.T) {
		data := clubMember.Core{ID: 1, UserID: 2, ClubID: 1, Status: "Owner"}
		data2 := club.Status{ID: 1, UserID: 2, ClubID: 1, Status: "Owner"}

		repo.On("GetById", 1).Return(data, nil).Once()
		repo2.On("GetStatus", 1, 2).Return(data2, nil).Once()
		repo.On("Delete", 1).Return(nil).Once()
		repo.On("UpdateMember", 1).Return(nil).Once()

		srv := New(repo, repo2)
		err := srv.Delete(id, userId)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("failed delete members", func(t *testing.T) {

		repo.On("GetById", 1).Return(clubMember.Core{}, errors.New("failed")).Once()

		srv := New(repo, repo2)
		err := srv.Delete(id, userId)
		assert.NotNil(t, err)
		// assert.NotEqual(t, inputRepo.Name, response.Name)
		repo.AssertExpectations(t)
	})

	t.Run("failed delete members", func(t *testing.T) {
		data := clubMember.Core{ID: 1, UserID: 2, ClubID: 1, Status: "Owner"}
		repo.On("GetById", 1).Return(data, nil).Once()
		repo2.On("GetStatus", id, userId).Return(club.Status{}, errors.New("error delete club. no data")).Once()

		srv := New(repo, repo2)
		err := srv.Delete(id, userId)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("failed delete members", func(t *testing.T) {
		data := club.Status{ID: 1, UserID: 2, ClubID: 1, Status: "Member"}
		data2 := clubMember.Core{ID: 1, UserID: 2, ClubID: 1, Status: "Owner"}

		repo.On("GetById", 1).Return(data2, nil).Once()
		repo2.On("GetStatus", id, userId).Return(data, nil).Once()

		srv := New(repo, repo2)
		err := srv.Delete(id, userId)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("failed delete member", func(t *testing.T) {
		data := club.Status{ID: 1, UserID: 2, ClubID: 3, Status: "Owner"}
		data2 := clubMember.Core{ID: 1, UserID: 2, ClubID: 1, Status: "Owner"}

		repo.On("GetById", 1).Return(data2, nil).Once()
		repo2.On("GetStatus", 1, 2).Return(data, nil).Once()
		repo.On("Delete", 1).Return(errors.New("failed")).Once()

		srv := New(repo, repo2)
		err := srv.Delete(id, userId)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("failed delete member", func(t *testing.T) {
		data := club.Status{ID: 1, UserID: 2, ClubID: 3, Status: "Owner"}
		data2 := clubMember.Core{ID: 1, UserID: 2, ClubID: 1, Status: "Owner"}

		repo.On("GetById", 1).Return(data2, nil).Once()
		repo2.On("GetStatus", 1, 2).Return(data, nil).Once()
		repo.On("Delete", 1).Return(nil).Once()
		repo.On("UpdateMember", 1).Return(errors.New("failed")).Once()

		srv := New(repo, repo2)
		err := srv.Delete(id, userId)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	repo := new(mocks.ClubMemberRepo)
	repo2 := new(mocks.ClubRepo)

	t.Run("Succes Update members", func(t *testing.T) {
		data := clubMember.Core{ID: 1, UserID: 2, User: clubMember.User{Name: "adaa", Email: "Joni@gmail.com"}, ClubID: 1, Status: "Member"}
		dataStatus := club.Status{ID: 1, UserID: 2, ClubID: 1, Status: "Owner"}
		dataClub := club.Core{ID: 1, Name: "club", Address: "disana", City: "disini", CategoryID: 1, Description: "joss", Logo: "jpg", JoinedMember: 1, MemberTotal: 2, Rule: "no rule", Requirement: "no req"}

		repo2.On("GetById", 1).Return(dataClub, nil).Once()
		repo.On("GetById", 1).Return(data, nil).Once()
		repo2.On("GetStatus", 1, 2).Return(dataStatus, nil).Once()
		repo.On("Update", data, 1).Return(nil).Once()
		repo.On("UpdateMember", 1).Return(nil).Once()

		srv := New(repo, repo2)
		err := srv.Update(data, 1, 2)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("failed update members", func(t *testing.T) {
		data := clubMember.Core{ID: 1, UserID: 2, User: clubMember.User{Name: "adaa", Email: "Joni@gmail.com"}, ClubID: 1, Status: "Member"}
		repo2.On("GetById", 1).Return(club.Core{}, errors.New("failed")).Once()

		srv := New(repo, repo2)
		err := srv.Update(data, 1, 2)
		assert.NotNil(t, err)
		// assert.NotEqual(t, inputRepo.Name, response.Name)
		repo.AssertExpectations(t)
	})

	t.Run("failed update members", func(t *testing.T) {
		data := clubMember.Core{ID: 1, UserID: 2, User: clubMember.User{Name: "adaa", Email: "Joni@gmail.com"}, ClubID: 1, Status: "Member"}

		dataClub := club.Core{ID: 1, Name: "club", Address: "disana", City: "disini", CategoryID: 1, Description: "joss", Logo: "jpg", JoinedMember: 2, MemberTotal: 1, Rule: "no rule", Requirement: "no req"}

		repo2.On("GetById", 1).Return(dataClub, nil).Once()

		srv := New(repo, repo2)
		err := srv.Update(data, 1, 2)
		assert.NotNil(t, err)
		// assert.NotEqual(t, inputRepo.Name, response.Name)
		repo.AssertExpectations(t)
	})

	t.Run("failed update members", func(t *testing.T) {
		data := clubMember.Core{ID: 1, UserID: 2, User: clubMember.User{Name: "adaa", Email: "Joni@gmail.com"}, ClubID: 1, Status: "Member"}

		dataClub := club.Core{ID: 1, Name: "club", Address: "disana", City: "disini", CategoryID: 1, Description: "joss", Logo: "jpg", JoinedMember: 2, MemberTotal: 4, Rule: "no rule", Requirement: "no req"}

		repo2.On("GetById", 1).Return(dataClub, nil).Once()
		repo.On("GetById", 1).Return(clubMember.Core{}, errors.New("failed")).Once()

		srv := New(repo, repo2)
		err := srv.Update(data, 1, 2)
		assert.NotNil(t, err)
		// assert.NotEqual(t, inputRepo.Name, response.Name)
		repo.AssertExpectations(t)
	})

	t.Run("failed update members", func(t *testing.T) {
		data := clubMember.Core{ID: 1, UserID: 2, User: clubMember.User{Name: "adaa", Email: "Joni@gmail.com"}, ClubID: 1, Status: "Member"}

		dataClub := club.Core{ID: 1, Name: "club", Address: "disana", City: "disini", CategoryID: 1, Description: "joss", Logo: "jpg", JoinedMember: 2, MemberTotal: 4, Rule: "no rule", Requirement: "no req"}

		repo2.On("GetById", 1).Return(dataClub, nil).Once()
		repo.On("GetById", 1).Return(data, nil).Once()
		repo2.On("GetStatus", 1, 2).Return(club.Status{}, errors.New("failed")).Once()

		srv := New(repo, repo2)
		err := srv.Update(data, 1, 2)
		assert.NotNil(t, err)
		// assert.NotEqual(t, inputRepo.Name, response.Name)
		repo.AssertExpectations(t)
	})

	t.Run("failed update members", func(t *testing.T) {
		data := clubMember.Core{ID: 1, UserID: 2, User: clubMember.User{Name: "adaa", Email: "Joni@gmail.com"}, ClubID: 1, Status: "Member"}

		dataStatus := club.Status{ID: 1, UserID: 2, ClubID: 1, Status: "Member"}

		dataClub := club.Core{ID: 1, Name: "club", Address: "disana", City: "disini", CategoryID: 1, Description: "joss", Logo: "jpg", JoinedMember: 2, MemberTotal: 4, Rule: "no rule", Requirement: "no req"}

		repo2.On("GetById", 1).Return(dataClub, nil).Once()
		repo.On("GetById", 1).Return(data, nil).Once()
		repo2.On("GetStatus", 1, 2).Return(dataStatus, nil).Once()

		srv := New(repo, repo2)
		err := srv.Update(data, 1, 2)
		assert.NotNil(t, err)
		// assert.NotEqual(t, inputRepo.Name, response.Name)
		repo.AssertExpectations(t)
	})

	t.Run("failed update members", func(t *testing.T) {
		data := clubMember.Core{ID: 1, UserID: 2, User: clubMember.User{Name: "adaa", Email: "Joni@gmail.com"}, ClubID: 1, Status: "Member"}

		dataStatus := club.Status{ID: 1, UserID: 2, ClubID: 1, Status: "Owner"}

		dataClub := club.Core{ID: 1, Name: "club", Address: "disana", City: "disini", CategoryID: 1, Description: "joss", Logo: "jpg", JoinedMember: 2, MemberTotal: 4, Rule: "no rule", Requirement: "no req"}

		repo2.On("GetById", 1).Return(dataClub, nil).Once()
		repo.On("GetById", 1).Return(data, nil).Once()
		repo2.On("GetStatus", 1, 2).Return(dataStatus, nil).Once()
		repo.On("Update", data, 1).Return(errors.New("failed")).Once()

		srv := New(repo, repo2)
		err := srv.Update(data, 1, 2)
		assert.NotNil(t, err)
		// assert.NotEqual(t, inputRepo.Name, response.Name)
		repo.AssertExpectations(t)
	})
	t.Run("failed update members", func(t *testing.T) {
		data := clubMember.Core{ID: 1, UserID: 2, User: clubMember.User{Name: "adaa", Email: "Joni@gmail.com"}, ClubID: 1, Status: "Member"}

		dataStatus := club.Status{ID: 1, UserID: 2, ClubID: 1, Status: "Owner"}

		dataClub := club.Core{ID: 1, Name: "club", Address: "disana", City: "disini", CategoryID: 1, Description: "joss", Logo: "jpg", JoinedMember: 2, MemberTotal: 4, Rule: "no rule", Requirement: "no req"}

		repo2.On("GetById", 1).Return(dataClub, nil).Once()
		repo.On("GetById", 1).Return(data, nil).Once()
		repo2.On("GetStatus", 1, 2).Return(dataStatus, nil).Once()
		repo.On("Update", data, 1).Return(nil).Once()
		repo.On("UpdateMember", 1).Return(errors.New("failed")).Once()

		srv := New(repo, repo2)
		err := srv.Update(data, 1, 2)
		assert.NotNil(t, err)
		// assert.NotEqual(t, inputRepo.Name, response.Name)
		repo.AssertExpectations(t)
	})
}
