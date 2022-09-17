package handler

import (
	"net/http"
	"qlambi/helper"
	"qlambi/user"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func UserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	//get Input From User
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Akun gagal dibuat", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	newUser, err := h.userService.RegisterUser(input)

	if true {
		response := helper.APIResponse("Akun gagal dibuat", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userReponse := user.FormatUser(newUser, "token")

	response := helper.APIResponse("Akun berhasil dibuat", http.StatusOK, "success", userReponse)
	c.JSON(http.StatusOK, response)
}
