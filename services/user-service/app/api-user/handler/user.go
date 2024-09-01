package handler

import (
	"user-service/business/user"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"

	dto "user-service/app/api-user/dto"
)

type UserHandler struct {
	usecase *user.UserService
}

func NewUserHandler(usecase *user.UserService) UserHandler {
	return UserHandler{
		usecase: usecase,
	}
}

// Register User
// @Summary      Register User
// @Description  This API for register user
// @Tags         users
// @Param		 request body dto.UserDTO{} true "request body"
// @Accept       json
// @Produce      json
// @Success 200 {object} dto.UserResp{}
// @Router /users/v1/register [post]
func (u *UserHandler) Register(c *fiber.Ctx) error {
	var payload dto.UserDTO

	if err := c.BodyParser(&payload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := c.JSON(payload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	validate := validator.New()
	if err := validate.Struct(payload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	user, err := u.usecase.CreateOne(c.Context(), payload)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(fiber.Map{
		"data":  user,
		"error": nil,
	})
}

// Login User
// @Summary      Login User
// @Description  This API for login user
// @Tags         users
// @Param		 request body dto.LoginDTO{} true "request body"
// @Accept       json
// @Produce      json
// @Success 200 {object} dto.LoginResp{}
// @Router /users/v1/login [post]
func (u *UserHandler) Login(c *fiber.Ctx) error {
	var payload dto.LoginDTO

	if err := c.BodyParser(&payload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := c.JSON(payload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	validate := validator.New()
	if err := validate.Struct(payload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	result, err := u.usecase.Login(c.Context(), payload)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(fiber.Map{
		"data":  result,
		"error": nil,
	})
}

// Refresh Token
// @Summary      Refresh Token
// @Description  This API for refresh token
// @Tags         users
// @Param		 request body dto.RefreshTokenDTO{} true "request body"
// @Accept       json
// @Produce      json
// @Success 200 {object} dto.LoginResp{}
// @Router /users/v1/refresh-token [post]
// @Security ApiKeyAuth
func (u *UserHandler) RefreshToken(c *fiber.Ctx) error {
	var payload = dto.RefreshTokenDTO{
		AccessToken: c.GetReqHeaders()["Authorization"][0],
	}

	if err := c.BodyParser(&payload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := c.JSON(payload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	validate := validator.New()
	if err := validate.Struct(payload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	result, err := u.usecase.RefreshToken(c.Context(), payload)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(fiber.Map{
		"data":  result,
		"error": nil,
	})
}
