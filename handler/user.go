package handler

import (
	"net/http"
	"startup-api/helper"
	"startup-api/user"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	// tangkap input dari user
	// map input dari user ke struct RegisterUserInput
	// struct di atas kita passing sebagai parameter service

	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}
		// response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", err.Error())
		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	// h.userService.RegisterUser(input)
	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(newUser, "tokentokentoken")
	jsonResponse := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)
	// c.JSON(http.StatusOK, nil)
	c.JSON(http.StatusOK, jsonResponse)
}

func (h *userHandler) Login(c *gin.Context) {
	// user memasukan input
	// input ditangkap handler
	// maping dari input user ke struct
	// input struct passing service
	// diservice mencari dengan bantuan repository user dengan email x
	// mencocokan password
	var input user.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedinUser,err := h.userService.Login(input)

	if err != nil {
		errorMessage := gin.H{"errors":err.Error()}

		response := helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := user.FormatUser(loggedinUser, "tokentokentoken")

	response := helper.APIResponse("Successfuly loggedin", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) CheckEmailAvailability(c *gin.Context) {
	// data input email dari user
	// input email di mapping ke struct input
	// struct input di passing ke service
	// service akan memanggil repository = email sudah ada atau belum
	// repository = db
	var input user.CheckEmailInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Email checking Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	IsEmailAvailable, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		errorMessage := gin.H{"errors": "Server error"}
		response := helper.APIResponse("Email checking Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	data := gin.H{
		"is_available": IsEmailAvailable,
	}
	
	metaMessage := "Email has ben registered"

	if IsEmailAvailable {
		metaMessage = "Email is available"
	}
	response := helper.APIResponse(metaMessage, http.StatusOK, "success", data)
	c.JSON(http.StatusUnprocessableEntity, response)


}

func (h *userHandler) UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{"is_upload":false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)
	
		c.JSON(http.StatusBadRequest, response)
		return
	}
	path := "images/" + file.Filename
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_upload":false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)
	
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// harusnya dapat dari jwt
	userID := 2
	_, err = h.userService.SaveAvatar(userID, path)
	if err != nil {
		data := gin.H{"is_upload": false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)
	
		c.JSON(http.StatusBadRequest, response)
		return
	}

		data := gin.H{"is_upload": true}
		response := helper.APIResponse("Avatar succesfuly uploaded", http.StatusOK, "success", data)
	
		c.JSON(http.StatusOK, response)
}