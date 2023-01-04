package service

import (
	"errors"
	"ikuzports/features/chat"
	"ikuzports/features/club"
	"ikuzports/features/clubActivity"
	"ikuzports/features/clubMember"
	"ikuzports/features/galery"
	"ikuzports/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetById(t *testing.T) {
	repo := new(mocks.ClubRepo)
	repo2 := new(mocks.ClubMemberRepo)
	t.Run("Success get club", func(t *testing.T) {
		inputRepo := club.Core{ID: 1, Name: "club", Address: "disana", City: "disini", CategoryID: 1, Description: "joss", Logo: "jpg", JoinedMember: 1, MemberTotal: 2, Rule: "no rule", Requirement: "no req"}

		repo.On("GetById", 1).Return(inputRepo, nil).Once()
		srv := New(repo, repo2)
		response, err := srv.GetById(1)

		assert.NoError(t, err)
		assert.Equal(t, inputRepo.Name, response.Name)
		repo.AssertExpectations(t)
	})

	t.Run("Failed get club by id", func(t *testing.T) {
		inputRepo := club.Core{ID: 1, Name: "club", Address: "disana", City: "disini", CategoryID: 1, Description: "joss", Logo: "jpg", JoinedMember: 1, MemberTotal: 2, Rule: "no rule", Requirement: "no req"}

		repo.On("GetById", 1).Return(club.Core{}, errors.New("failed")).Once()
		srv := New(repo, repo2)
		response, err := srv.GetById(1)

		assert.NotNil(t, err)
		assert.NotEqual(t, inputRepo.Name, response.Name)
		repo.AssertExpectations(t)
	})
}

func TestGetMembers(t *testing.T) {
	repo := new(mocks.ClubRepo)
	repo2 := new(mocks.ClubMemberRepo)
	t.Run("Success get all clubs member", func(t *testing.T) {
		inputRepo := []clubMember.Core{{ID: 1, UserID: 1, ClubID: 1, Status: "member"}}

		repo.On("GetMembers", 1).Return(inputRepo, nil).Once()
		srv := New(repo, repo2)
		response, err := srv.GetMembers(1)

		assert.NoError(t, err)
		assert.Equal(t, inputRepo[0].ID, response[0].ID)
		repo.AssertExpectations(t)
	})

	t.Run("Failed get all members", func(t *testing.T) {
		repo.On("GetMembers", 1).Return(nil, errors.New("failed")).Once()

		srv := New(repo, repo2)
		response, err := srv.GetMembers(1)

		assert.NotNil(t, err)
		assert.Nil(t, response)
		repo.AssertExpectations(t)
	})
}

func TestGetChats(t *testing.T) {
	repo := new(mocks.ClubRepo)
	repo2 := new(mocks.ClubMemberRepo)
	t.Run("Success get all clubs chat", func(t *testing.T) {
		inputRepo := []chat.Core{{ID: 1, UserID: 1, ClubID: 1, Message: "member"}}

		repo.On("GetChats", 1).Return(inputRepo, nil).Once()
		srv := New(repo, repo2)
		response, err := srv.GetChats(1)

		assert.NoError(t, err)
		assert.Equal(t, inputRepo[0].ID, response[0].ID)
		repo.AssertExpectations(t)
	})

	t.Run("Failed get all chats", func(t *testing.T) {
		repo.On("GetChats", 1).Return(nil, errors.New("failed")).Once()

		srv := New(repo, repo2)
		response, err := srv.GetChats(1)
		assert.NotNil(t, err)
		assert.Nil(t, response)
		repo.AssertExpectations(t)
	})
}

func TestGetGalleries(t *testing.T) {
	repo := new(mocks.ClubRepo)
	repo2 := new(mocks.ClubMemberRepo)
	t.Run("Success get all clubs galleries", func(t *testing.T) {
		inputRepo := []galery.Core{{ID: 1, Url: "url", ClubID: 1, Caption: "member"}}

		repo.On("GetGaleries", 1).Return(inputRepo, nil).Once()
		srv := New(repo, repo2)
		response, err := srv.GetGaleries(1)

		assert.NoError(t, err)
		assert.Equal(t, inputRepo[0].ID, response[0].ID)
		repo.AssertExpectations(t)
	})

	t.Run("Failed get all galeries", func(t *testing.T) {
		repo.On("GetGaleries", 1).Return(nil, errors.New("failed")).Once()

		srv := New(repo, repo2)
		response, err := srv.GetGaleries(1)
		assert.NotNil(t, err)
		assert.Nil(t, response)
		repo.AssertExpectations(t)
	})
}

func TestGetActivities(t *testing.T) {
	repo := new(mocks.ClubRepo)
	repo2 := new(mocks.ClubMemberRepo)
	t.Run("Success get all clubs activities", func(t *testing.T) {
		inputRepo := []clubActivity.Core{{ID: 1, Name: "main bola", ClubID: 1, StartTime: "15:00", EndTime: "16:00", Day: "kamis", Location: "lapangan", ActivityDetail: "main aja"}}

		repo.On("GetActivities", 1).Return(inputRepo, nil).Once()
		srv := New(repo, repo2)
		response, err := srv.GetActivities(1)

		assert.NoError(t, err)
		assert.Equal(t, inputRepo[0].ID, response[0].ID)
		repo.AssertExpectations(t)
	})

	t.Run("Failed get all activities", func(t *testing.T) {
		repo.On("GetActivities", 1).Return(nil, errors.New("failed")).Once()

		srv := New(repo, repo2)
		response, err := srv.GetActivities(1)
		assert.NotNil(t, err)
		assert.Nil(t, response)
		repo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	repo := new(mocks.ClubRepo)
	repo2 := new(mocks.ClubMemberRepo)
	id := 1
	userId := 2
	t.Run("Success delete clubs", func(t *testing.T) {
		data := club.Status{ID: 1, UserID: 2, ClubID: 3, Status: "Owner"}

		repo.On("GetStatus", 1, 2).Return(data, nil).Once()
		repo.On("Delete", 1).Return(nil).Once()
		repo2.On("DeleteMember", 1).Return(nil).Once()

		srv := New(repo, repo2)
		err := srv.Delete(id, userId)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("failed delete clubs", func(t *testing.T) {
		data := club.Status{ID: 1, UserID: 2, ClubID: 3, Status: "member"}

		repo.On("GetStatus", 1, 2).Return(data, nil).Once()

		srv := New(repo, repo2)
		err := srv.Delete(id, userId)
		assert.NotNil(t, err)
		// assert.NotEqual(t, inputRepo.Name, response.Name)
		repo.AssertExpectations(t)
	})

	t.Run("failed delete clubs", func(t *testing.T) {
		sts := club.Status{}
		repo.On("GetStatus", id, userId).Return(sts, errors.New("error delete club. no data")).Once()

		srv := New(repo, repo2)
		err := srv.Delete(id, userId)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("failed delete club", func(t *testing.T) {
		data := club.Status{ID: 1, UserID: 2, ClubID: 3, Status: "Owner"}

		repo.On("GetStatus", 1, 2).Return(data, nil).Once()
		repo.On("Delete", 1).Return(errors.New("failed")).Once()

		srv := New(repo, repo2)
		err := srv.Delete(id, userId)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("failed delete club", func(t *testing.T) {
		data := club.Status{ID: 1, UserID: 2, ClubID: 3, Status: "Owner"}

		repo.On("GetStatus", 1, 2).Return(data, nil).Once()
		repo.On("Delete", 1).Return(nil).Once()
		repo2.On("DeleteMember", 1).Return(errors.New("failed")).Once()

		srv := New(repo, repo2)
		err := srv.Delete(id, userId)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	repo := new(mocks.ClubRepo)
	repo2 := new(mocks.ClubMemberRepo)
	id := 1
	userId := 2
	t.Run("Success update clubs", func(t *testing.T) {
		inputRepo := club.Core{ID: 1, Name: "club", Address: "disana", City: "disini", CategoryID: 1, Description: "joss", Logo: "jpg", JoinedMember: 1, MemberTotal: 2, Rule: "no rule", Requirement: "no req"}
		data := club.Status{ID: 1, UserID: 2, ClubID: 3, Status: "Owner"}

		repo.On("GetStatus", 1, 2).Return(data, nil).Once()
		repo.On("Update", inputRepo, 1).Return(nil).Once()

		srv := New(repo, repo2)
		err := srv.Update(inputRepo, id, userId)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("failed updates clubs", func(t *testing.T) {
		sts := club.Status{}
		inputRepo := club.Core{ID: 1, Name: "club", Address: "disana", City: "disini", CategoryID: 1, Description: "joss", Logo: "jpg", JoinedMember: 1, MemberTotal: 2, Rule: "no rule", Requirement: "no req"}

		repo.On("GetStatus", id, userId).Return(sts, errors.New("error delete club. no data")).Once()

		srv := New(repo, repo2)
		err := srv.Update(inputRepo, id, userId)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("failed updates clubs", func(t *testing.T) {
		inputRepo := club.Core{ID: 1, Name: "club", Address: "disana", City: "disini", CategoryID: 1, Description: "joss", Logo: "jpg", JoinedMember: 1, MemberTotal: 2, Rule: "no rule", Requirement: "no req"}

		data := club.Status{ID: 1, UserID: 2, ClubID: 3, Status: "member"}

		repo.On("GetStatus", id, userId).Return(data, nil).Once()

		srv := New(repo, repo2)
		err := srv.Update(inputRepo, id, userId)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Update clubs", func(t *testing.T) {
		inputRepo := club.Core{ID: 1, Name: "club", Address: "disana", City: "disini", CategoryID: 1, Description: "joss", Logo: "jpg", JoinedMember: 1, MemberTotal: 2, Rule: "no rule", Requirement: "no req"}
		data := club.Status{ID: 1, UserID: 2, ClubID: 3, Status: "Owner"}

		repo.On("GetStatus", 1, 2).Return(data, nil).Once()
		repo.On("Update", inputRepo, 1).Return(errors.New("failed")).Once()

		srv := New(repo, repo2)
		err := srv.Update(inputRepo, id, userId)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestCreate(t *testing.T) {
	repo := new(mocks.ClubRepo)
	repo2 := new(mocks.ClubMemberRepo)

	t.Run("succes create club", func(t *testing.T) {
		inputRepo := club.Core{ID: 1, Name: "club", Address: "disana", City: "disini", CategoryID: 1, Description: "joss", Logo: "jpg", JoinedMember: 1, MemberTotal: 2, Rule: "no rule", Requirement: "no req"}

		dataMember := clubMember.Core{UserID: 1, ClubID: 1, Status: "Owner"}

		repo.On("Create", inputRepo, 1).Return(nil).Once()
		repo.On("GetLastID").Return(1, nil).Once()
		repo2.On("Create", dataMember).Return(nil).Once()
		repo.On("UpdateMember", 1).Return(1, nil).Once()

		srv := New(repo, repo2)
		err := srv.Create(inputRepo, 1)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Create club, name empty", func(t *testing.T) {
		inputRepo := club.Core{ID: 1, Address: "disana", City: "disini", CategoryID: 1, Description: "joss", Logo: "jpg", JoinedMember: 1, MemberTotal: 2, Rule: "no rule", Requirement: "no req"}

		srv := New(repo, repo2)
		err := srv.Create(inputRepo, 1)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Create club", func(t *testing.T) {
		inputRepo := club.Core{ID: 1, Name: "jojo", Address: "disana", City: "disini", CategoryID: 1, Description: "joss", Logo: "jpg", JoinedMember: 1, MemberTotal: 2, Rule: "no rule", Requirement: "no req"}

		repo.On("Create", inputRepo, 1).Return(errors.New("failed")).Once()

		srv := New(repo, repo2)
		err := srv.Create(inputRepo, 1)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed get last id", func(t *testing.T) {
		inputRepo := club.Core{ID: 1, Name: "jojo", Address: "disana", City: "disini", CategoryID: 1, Description: "joss", Logo: "jpg", JoinedMember: 1, MemberTotal: 2, Rule: "no rule", Requirement: "no req"}

		repo.On("Create", inputRepo, 1).Return(nil).Once()
		repo.On("GetLastID").Return(0, errors.New("failed")).Once()

		srv := New(repo, repo2)
		err := srv.Create(inputRepo, 1)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed create data member", func(t *testing.T) {
		inputRepo := club.Core{ID: 1, Name: "jojo", Address: "disana", City: "disini", CategoryID: 1, Description: "joss", Logo: "jpg", JoinedMember: 1, MemberTotal: 2, Rule: "no rule", Requirement: "no req"}

		dataMember := clubMember.Core{UserID: 1, ClubID: 1, Status: "Owner"}

		repo.On("Create", inputRepo, 1).Return(nil).Once()
		repo.On("GetLastID").Return(1, nil).Once()
		repo2.On("Create", dataMember).Return(errors.New("failed")).Once()

		srv := New(repo, repo2)
		err := srv.Create(inputRepo, 1)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed update joined member", func(t *testing.T) {
		inputRepo := club.Core{ID: 1, Name: "jojo", Address: "disana", City: "disini", CategoryID: 1, Description: "joss", Logo: "jpg", JoinedMember: 1, MemberTotal: 2, Rule: "no rule", Requirement: "no req"}

		dataMember := clubMember.Core{UserID: 1, ClubID: 1, Status: "Owner"}

		repo.On("Create", inputRepo, 1).Return(nil).Once()
		repo.On("GetLastID").Return(1, nil).Once()
		repo2.On("Create", dataMember).Return(nil).Once()
		repo.On("UpdateMember", 1).Return(0, errors.New("failed")).Once()

		srv := New(repo, repo2)
		err := srv.Create(inputRepo, 1)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestGetAll(t *testing.T) {
	repo := new(mocks.ClubRepo)
	repo2 := new(mocks.ClubMemberRepo)
	// limit := 1
	// offset := 2

	t.Run("Succes get all clubs", func(t *testing.T) {
		data := []club.Core{{ID: 1, Name: "club", Address: "disana", City: "disini", CategoryID: 1, Description: "joss", Logo: "jpg", JoinedMember: 1, MemberTotal: 2, Rule: "no rule", Requirement: "no req"}}

		repo.On("GetAllWithSearch", "name", "city", 1, 9, 9).Return(data, 9, nil).Once()

		srv := New(repo, repo2)
		res, _, err := srv.GetAll("name", "city", 1, 2)
		assert.NoError(t, err)
		assert.Equal(t, data[0].Name, res[0].Name)
		repo.AssertExpectations(t)
	})

	t.Run("Succes get all clubs", func(t *testing.T) {
		data := []club.Core{{ID: 1, Address: "disana", Description: "joss", Logo: "jpg", JoinedMember: 1, MemberTotal: 2, Rule: "no rule", Requirement: "no req"}}

		repo.On("GetAll", 9, 9).Return(data, 10, nil).Once()

		srv := New(repo, repo2)
		res, _, err := srv.GetAll("", "", 0, 2)
		assert.NoError(t, err)
		assert.Equal(t, data[0].Name, res[0].Name)
		repo.AssertExpectations(t)
	})

	t.Run("Failed get all clubs", func(t *testing.T) {
		repo.On("GetAllWithSearch", "name", "city", 1, 9, 9).Return(nil, 0, errors.New("faield")).Once()

		srv := New(repo, repo2)
		_, _, err := srv.GetAll("name", "city", 1, 2)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

}
