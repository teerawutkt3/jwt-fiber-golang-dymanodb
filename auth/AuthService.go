package auth

import (
	"errors"
	"fiber-poc-api/database/entity"
	"fiber-poc-api/database/repository"
	"fiber-poc-api/utils"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"time"
)

type AuthService struct {
	repo repository.UserRepository
}

func NewAuthService(repo repository.UserRepository) AuthService {
	return AuthService{repo}
}

func (svc AuthService) Login(req LoginReq, xRequestId string) (*string, error) {

	user, err := svc.repo.GetUserByUsername(req.Username, xRequestId)
	if err != nil {
		log.Infof("[%s] error during GetUserByUsername err: %+v", xRequestId, err.Error())
		return nil, err
	}
	err = utils.ValidatePassword(user.Password, req.Password)
	if err != nil {
		log.Infof("[%s] ValidatePassword not pass, UNAUTHORIZED", xRequestId)
		return nil, errors.New("UNAUTHORIZED")
	}

	expTime := viper.GetDuration("jwt.expire")
	claims := jwt.MapClaims{
		"username": user.Username,
		"admin":    true,
		"exp":      time.Now().Add(time.Hour * expTime).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(viper.GetString("jwt.secret")))
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (svc AuthService) Register(req LoginReq, xRequestId string) error {

	// Validate
	userExit, err := svc.repo.GetUserByUsername(req.Username, xRequestId)
	if err != nil {
		log.Infof("[%s] error during GetUserByUsername err: %+v", xRequestId, err.Error())
		return err
	}
	if userExit != nil {
		msg := "username exits"
		log.Errorf("[%s] error register %s %s", xRequestId, req.Username, msg)
		return errors.New(msg)
	}
	// Create User
	createdDate := utils.FormatDate(time.Now(), "2006-02-01 15:04:05")
	password, _ := utils.HashPassword(req.Password)
	user := &entity.User{
		Username:    req.Username,
		Password:    password,
		IsDeleted:   "N",
		CreatedDate: createdDate,
		UpdatedDate: createdDate,
	}
	err = svc.repo.CreateUser(user)
	if err != nil {
		log.Errorf("[%s] CreateUser error: %+v", xRequestId, err.Error())
		return err
	}
	return nil
}

func (svc AuthService) UpdateUser(req LoginReq, xRequestId string) error {

	user := &entity.User{
		Username:    req.Username,
		Password:    req.Password,
		IsDeleted:   "N",
		UpdatedDate: utils.FormatDate(time.Now(), "2006-02-01 15:04:05"),
	}
	err := svc.repo.UpdateUser(user)
	if err != nil {
		log.Errorf("[%s] CreateUser error: %+v", xRequestId, err.Error())
		return err
	}
	return nil
}

func (svc AuthService) GetUser(req LoginReq, xRequestId string) error {

	log.Infof("trestdddd")
	_, err := svc.repo.GetUserByUsername("",xRequestId)
	if err != nil {
		log.Errorf("[%s] CreateUser error: %+v", xRequestId, err.Error())
		return err
	}
	return nil
}
