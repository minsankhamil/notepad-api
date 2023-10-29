package controllers

import (
	"log"
	"net/http"
	"notepad-api/configs"
	"notepad-api/middleware"
	"notepad-api/models"
	"notepad-api/repositories"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

var validate = validator.New()

func CreateUser(c echo.Context) error {
	var user models.User
	c.Bind(&user)

	// Validate the user data
	if err := validate.Struct(user); err != nil {
		log.Println(err.Error())
		return c.JSON(400, models.BaseResponse{
			Status:  false,
			Message: "Data masukan user tidak valid",
			Data:    nil,
		})
	}

	//encrypt the user password data
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	//add user data to database
	err = repositories.AddUser(&user)
	if err != nil {
		return c.JSON(500, models.BaseResponse{
			Status:  false,
			Message: "Gagal menambahkan user",
			Data:    nil,
		})
	}
	return c.JSON(200, models.BaseResponse{
		Status:  true,
		Message: "Sukses menambahkan user",
		Data:    user,
	})

}

// show the profile of user
func ProfileUser(c echo.Context) error {
	var oneuser models.User
	id := c.Param("id")

	//validate route is correct, :id must number
	IdUser, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return c.JSON(400, models.BaseResponse{
			Status:  false,
			Message: "Input harus berupa bilangan bulat",
			Data:    nil,
		})
	}

	//search the user profile
	err = repositories.GetOneUser(uint(IdUser), &oneuser)
	if err != nil {
		return c.JSON(500, models.BaseResponse{
			Status:  false,
			Message: "Gagal menemukan user yang dicari",
			Data:    nil,
		})
	}
	return c.JSON(200, models.BaseResponse{
		Status:  true,
		Message: "Sukses mendapatkan user yang dicari",
		Data:    oneuser,
	})

}

func DeleteOneUser(c echo.Context) error {
	var oneuser models.User
	id := c.Param("id")

	//validate route is correct, :id must number
	IdUser, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return c.JSON(400, models.BaseResponse{
			Status:  false,
			Message: "Input harus berupa bilangan bulat",
			Data:    nil,
		})
	}

	//delete user data
	err = repositories.DeleteOneUser(uint(IdUser), &oneuser)
	if err != nil {
		return c.JSON(500, models.BaseResponse{
			Status:  false,
			Message: "Gagal menghapus user yang dicari",
			Data:    nil,
		})
	}
	return c.JSON(200, models.BaseResponse{
		Status:  true,
		Message: "Berhasil menghapus user yang dicari",
		Data:    oneuser,
	})
}

func Login(c echo.Context) error {
	var user models.User
	c.Bind(&user)
	inputPassword := user.Password

	if err := configs.DB.Where("email = ?", user.Email).First(&user).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, "Email or password is incorrect")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(inputPassword)); err != nil {
		return c.JSON(http.StatusUnauthorized, "Email or password is incorrect")
	}

	tokenString, err := middleware.GenerateToken(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to generate token")
	}
	return c.JSON(http.StatusOK, map[string]string{
		"token": tokenString,
	})
}
