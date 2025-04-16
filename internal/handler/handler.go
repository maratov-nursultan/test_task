package handler

import (
	"database/sql"
	"errors"
	"github.com/labstack/echo/v4"
	user "github.com/maratov-nursultan/profile/internal/manager"
	"github.com/maratov-nursultan/profile/internal/model"
	"strconv"
)

type Handler struct {
	userManager user.ManagerSDK
}

func NewHandler(userManager user.ManagerSDK) *Handler {
	return &Handler{
		userManager: userManager,
	}
}

func (h *Handler) CheckIin(c echo.Context) error {
	iin := c.Param("iin")
	if iin == "" {
		return model.IinEmpty
	}

	if len(iin) != 12 {
		return model.ErrIinInvalid
	}

	_, err := strconv.Atoi(iin)
	if err != nil {
		return model.ErrIinInvalid
	}

	resp, err := h.userManager.CheckIin(iin)
	if err != nil {
		return err
	}

	return c.JSON(200, resp)
}

func (h *Handler) CreateUser(c echo.Context) error {
	req := new(model.InfoRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(500, model.NewInfo(false, err.Error()))
	}

	if req.Iin == "" {
		return c.JSON(500, model.NewInfo(false, ""))
	}

	if len(req.Iin) != 12 {
		return c.JSON(500, model.NewInfo(false, ""))
	}

	_, err := strconv.Atoi(req.Iin)
	if err != nil {
		return c.JSON(500, model.NewInfo(false, ""))
	}

	_, err = h.userManager.CheckIin(req.Iin)
	if err != nil {
		return err
	}

	err = h.userManager.CreateUser(c.Request().Context(), req)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.JSON(404, model.NewInfo(false, err.Error()))
		}
		return c.JSON(500, model.NewInfo(false, err.Error()))
	}

	return c.JSON(200, model.NewInfo(true, ""))
}

func (h *Handler) ListUserByName(c echo.Context) error {
	name := c.Param("name")
	if name == "" {
		return model.NameEmpty
	}

	resp, err := h.userManager.ListUserByName(c.Request().Context(), name)
	if err != nil {
		return err
	}

	return c.JSON(200, resp)
}

func (h *Handler) GetUserByIin(c echo.Context) error {
	iin := c.Param("iin")
	if iin == "" {
		return model.IinEmpty
	}

	if len(iin) != 12 {
		return model.ErrIinInvalid
	}

	_, err := strconv.Atoi(iin)
	if err != nil {
		return model.ErrIinInvalid
	}

	_, err = h.userManager.CheckIin(iin)
	if err != nil {
		return err
	}

	resp, err := h.userManager.GetUserByIin(c.Request().Context(), iin)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.JSON(404, model.NewInfo(false, err.Error()))
		}
		return c.JSON(500, model.NewInfo(false, err.Error()))
	}

	return c.JSON(200, resp)
}
