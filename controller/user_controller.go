package controller

import (
	"fmt"
	"net/http"

	"github.com/eulbyvan/enigma-university/model"
	"github.com/eulbyvan/enigma-university/model/dto/res"
	"github.com/eulbyvan/enigma-university/usecase"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userUseCase usecase.UserUseCase
}

func NewUserController(userUseCase usecase.UserUseCase) *UserController {
	return &UserController{userUseCase: userUseCase}
}

func (c *UserController) FindById(ctx *gin.Context) {
	userID := ctx.Query("id")

	var res res.CommonResponse

	user, err := c.userUseCase.FindById(userID)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	res.Code = http.StatusOK
	res.Status = "Success"
	res.Message = "Retrieved data successfully"
	res.Data = user

	ctx.JSON(http.StatusOK, res)
}

func (c *UserController) GetAllUsers(ctx *gin.Context) {
	_ = ctx.Query("/")

	var res res.CommonResponse

	// Memanggil fungsi GetAllUsers dari UserUseCase
	users, err := c.userUseCase.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get all users"})
		return
	}

	res.Code = http.StatusOK
	res.Status = "Success"
	res.Message = "Retrieved data successfully"
	res.Data = users

	ctx.JSON(http.StatusOK, res)
}

func (c *UserController) Create(ctx *gin.Context) {
	var user model.User

	// Mendekode data JSON dari body permintaan
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read JSON"})
		return
	}

	// Memanggil fungsi Registration dari UserUseCase
	if err := c.userUseCase.Registration(user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Mengembalikan respons berhasil
	ctx.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func (u *UserController) UpdateById(ctx *gin.Context) {
	// Mendapatkan ID dari parameter URL
	userID := ctx.Param("id")

	var updatedUser model.User

	// Mendekode data JSON dari body permintaan
	err := ctx.BindJSON(&updatedUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read JSON"})
		return
	}

	// Memanggil fungsi UpdateByID dari UserUseCase
	err = u.userUseCase.UpdateById(userID, updatedUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to update user with ID %s", userID)})
		return
	}

	// Mengembalikan respons berhasil
	ctx.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func (u *UserController) RemoveById(ctx *gin.Context) {
	// Mendapatkan ID dari parameter URL
	userID := ctx.Param("id")

	// Memanggil fungsi DeleteByID dari UserUseCase
	err := u.userUseCase.RemoveById(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to delete user with ID %s", userID)})
		return
	}

	// Mengembalikan respons berhasil
	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
