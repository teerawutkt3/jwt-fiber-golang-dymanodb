package auth

import (
	"fiber-poc-api/common"
	"fiber-poc-api/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"time"
)

type AuthHandler struct {
	svc AuthService
}

func NewAuthHandler(svc AuthService) AuthHandler {
	return AuthHandler{svc}
}

func (h AuthHandler) LoginHandler(c *fiber.Ctx) error {
	xRequestId := utils.GetXRequestId()
	req := LoginReq{}
	err := c.BodyParser(&req)
	if err != nil {
		log.Errorf("[%s] Login invalid Request err: %+v", xRequestId, err.Error())
		return c.Status(400).JSON("invalid request")
	}
	log.Infof("[%s] username %s login date: %+v", xRequestId, req.Username, time.Now())

	token, err := h.svc.Login(req, xRequestId)
	if err != nil {
		log.Errorf("[%s] Error during login err: %+v", xRequestId, err.Error())
		if "UNAUTHORIZED" == err.Error() {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	log.Infof("[%s] login success", xRequestId)
	response := LoginRes{}
	response.Token = *token
	return c.JSON(response)
}

func (h AuthHandler) RegisterHandler(c *fiber.Ctx) error {
	xRequestId := utils.GetXRequestId()
	req := LoginReq{}
	err := c.BodyParser(&req)
	if err != nil {
		log.Errorf("[%s] Register invalid Request err: %+v", xRequestId, err.Error())
		return c.Status(400).JSON("invalid request")
	}
	log.Infof("[%s] Req: %+v", xRequestId, req)
	err = h.svc.Register(req, xRequestId)
	if err != nil {
		log.Errorf("[%s] Register Error", xRequestId)
		return c.Status(500).JSON(fiber.Map{"message": "error"})
	}
	log.Infof("[%s] Register success", xRequestId)
	response := common.Response{
		Message: "success",
	}
	return c.JSON(response)
}

func (h AuthHandler) UpdateUserHandler(c *fiber.Ctx) error {
	xRequestId := utils.GetXRequestId()
	req := LoginReq{}
	err := c.BodyParser(&req)
	if err != nil {
		log.Errorf("[%s] Register invalid Request err: %+v", xRequestId, err.Error())
		return c.Status(400).JSON("invalid request")
	}
	log.Infof("[%s] Req: %+v", xRequestId, req)
	err = h.svc.UpdateUser(req, xRequestId)
	if err != nil {
		log.Errorf("[%s] Register Error", xRequestId)
		return c.Status(500).JSON(fiber.Map{"message": "error"})
	}
	log.Infof("[%s] Register success", xRequestId)
	return c.JSON(fiber.Map{"message": "success"})
}
func (h AuthHandler) GetUserHandler(c *fiber.Ctx) error {
	xRequestId := utils.GetXRequestId()
	req := LoginReq{}
	err := c.BodyParser(&req)
	if err != nil {
		log.Errorf("[%s] Register invalid Request err: %+v", xRequestId, err.Error())
		return c.Status(400).JSON("invalid request")
	}
	log.Infof("[%s] Req: %+v", xRequestId, req)
	err = h.svc.GetUser(req, xRequestId)
	if err != nil {
		log.Errorf("[%s] Register Error", xRequestId)
		return c.Status(500).JSON(fiber.Map{"message": "error"})
	}
	log.Infof("[%s] Register success", xRequestId)
	return c.JSON(fiber.Map{"message": "success"})
}
