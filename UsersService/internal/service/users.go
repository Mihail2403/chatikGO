package service

import (
	"errors"
	"log"
	"users-service/entity"
	"users-service/internal/repository"
)

func getImgIds(users []repository.UsersPostgresStruct) []string {
	imgIds := make([]string, 0)
	for _, user := range users {
		imgIds = append(imgIds, user.ImgID)
	}
	return imgIds
}

type UsersService struct {
	repo *repository.Repository
}

func NewUsersService(repo *repository.Repository) *UsersService {
	return &UsersService{repo: repo}
}

func (s *UsersService) Create(user *entity.User) error {
	imgId, err := s.repo.Img.Create(user.ImgBase64)
	if err != nil {
		return err
	}
	pgUsr := &repository.UsersPostgresStruct{
		Name:   user.Name,
		UserID: user.UserId,
		ImgID:  imgId,
	}
	err = s.repo.Users.Create(pgUsr)
	return err
}

func (s *UsersService) GetAll() ([]entity.User, error) {
	pgUsrs, err := s.repo.Users.GetAll()
	if err != nil {
		return nil, err
	}
	var users []entity.User
	var imgs []string
	for _, pgUsr := range pgUsrs {
		users = append(users, entity.User{
			Id:     pgUsr.ID,
			Name:   pgUsr.Name,
			UserId: pgUsr.UserID,
		})
		imgs = append(imgs, pgUsr.ImgID)
	}
	imgs, err = s.repo.Img.GetByIDArray(imgs)
	if err != nil {
		return nil, err
	}
	for i, img := range imgs {
		users[i].ImgBase64 = img
	}
	return users, nil
}

func (s *UsersService) GetByID(id int) (*entity.User, error) {
	pgUsr, err := s.repo.Users.GetByID(id)
	if err != nil {
		return nil, err
	}
	img, err := s.repo.Img.GetById(pgUsr.ImgID)
	if err != nil {
		return nil, err
	}
	return &entity.User{
		Id:        pgUsr.ID,
		Name:      pgUsr.Name,
		UserId:    pgUsr.UserID,
		ImgBase64: img,
	}, nil
}

func (s *UsersService) GetByIDArray(ids []int) ([]entity.User, error) {
	users, err := s.repo.Users.GetByIDArray(ids)
	if err != nil {
		return nil, err
	}
	imgs, err := s.repo.Img.GetByIDArray(getImgIds(users))
	if err != nil {
		return nil, err
	}
	if len(imgs) != len(users) {
		return nil, errors.New("img len not equal user len")
	}
	var usersRes []entity.User
	for i, pgUsr := range users {
		usersRes = append(usersRes, entity.User{
			Id:        pgUsr.ID,
			Name:      pgUsr.Name,
			UserId:    pgUsr.UserID,
			ImgBase64: imgs[i],
		})
	}
	return usersRes, nil
}

func (s *UsersService) Delete(id int) error {
	user, err := s.repo.Users.GetByID(id)
	if err != nil {
		log.Println(err)
		return err
	}
	err = s.repo.Img.Delete(user.ImgID)
	if err != nil {
		log.Println(err)
		return err
	}
	err = s.repo.Users.Delete(id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (s *UsersService) Update(id int, user *entity.User) error {
	pgUsr, err := s.repo.Users.GetByID(id)
	if err != nil {
		return err
	}
	if user.ImgBase64 != "" {
		err = s.repo.Img.Update(pgUsr.ImgID, user.ImgBase64)
		if err != nil {
			return err
		}
	}
	userForUpdate := &repository.UsersPostgresStruct{
		Name:  pgUsr.Name,
		ImgID: pgUsr.ImgID,
	}
	if user.Name != "" {
		userForUpdate.Name = user.Name
	}
	err = s.repo.Users.Update(id, userForUpdate)
	return err
}
