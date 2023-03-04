package service

import (
	"github.com/prayogatriady/todolist-app/model/entity"
	"github.com/prayogatriady/todolist-app/model/web"
	"github.com/prayogatriady/todolist-app/repository"
)

type ListServInterface interface {
	CreateTodolist(userID int64, listRequest web.ListRequest) (web.ListResponse, error)
	GetLists(userID int64) ([]web.ListResponse, error)
	EditList(listID int64, userID int64, listRequest web.ListRequest) (web.ListResponse, error)
	DeleteList(listID int64, userID int64) error
}

type ListServ struct {
	ListRepo repository.ListRepoInterface
}

func NewListServ(listRepo repository.ListRepoInterface) ListServInterface {
	return &ListServ{
		ListRepo: listRepo,
	}
}

func (us *ListServ) CreateTodolist(userID int64, listRequest web.ListRequest) (web.ListResponse, error) {
	var listEntity entity.List
	listEntity = entity.List{
		Title:       listRequest.Title,
		Description: listRequest.Description,
		UserID:      userID,
	}

	listEntity, err := us.ListRepo.CreateTodolist(listEntity)
	if err != nil {
		return web.ListResponse{}, err
	}

	var listResponse web.ListResponse
	listResponse = web.ListResponse{
		ID:          listEntity.ID,
		Title:       listEntity.Title,
		Description: listEntity.Description,
		UserID:      listEntity.UserID,
	}

	return listResponse, nil
}

func (us *ListServ) GetLists(userID int64) ([]web.ListResponse, error) {
	listsEntity, err := us.ListRepo.GetListsByUserID(userID)
	if err != nil {
		return []web.ListResponse{}, err
	}

	var listsResponse []web.ListResponse

	for _, list := range listsEntity {
		l := web.ListResponse{
			ID:          list.ID,
			Title:       list.Title,
			Description: list.Description,
			UserID:      list.UserID,
		}
		listsResponse = append(listsResponse, l)
	}

	return listsResponse, nil
}

func (us *ListServ) EditList(listID int64, userID int64, listRequest web.ListRequest) (web.ListResponse, error) {
	var listEntity entity.List
	listEntity = entity.List{
		Title:       listRequest.Title,
		Description: listRequest.Description,
		UserID:      userID,
	}
	listEntity, err := us.ListRepo.UpdateList(listID, userID, listEntity)
	if err != nil {
		return web.ListResponse{}, err
	}

	var listResponse web.ListResponse
	listResponse = web.ListResponse{
		ID:          listEntity.ID,
		Title:       listEntity.Title,
		Description: listEntity.Description,
		UserID:      listEntity.UserID,
	}

	return listResponse, nil
}

func (us *ListServ) DeleteList(listID int64, userID int64) error {
	if err := us.ListRepo.DeleteList(listID, userID); err != nil {
		return err
	}
	return nil
}
