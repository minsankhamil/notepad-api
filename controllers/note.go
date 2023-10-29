package controllers

import (
	"log"
	"notepad-api/models"
	"notepad-api/repositories"
	"strconv"

	"github.com/labstack/echo/v4"
)

func AddNewNote(c echo.Context) error {
	var note models.Note
	c.Bind(&note)

	err := repositories.AddNote(&note)
	if err != nil {
		log.Println(err.Error())
		return c.JSON(500, models.BaseResponse{
			Status:  false,
			Message: "Gagal menambahkan note",
			Data:    nil,
		})
	}
	return c.JSON(200, models.BaseResponse{
		Status:  true,
		Message: "Sukses menambahkan note",
		Data:    note,
	})

}

func ReadOneNote(c echo.Context) error {
	var onenote models.Note
	id := c.Param("id")
	IdNote, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return c.JSON(400, models.BaseResponse{
			Status:  false,
			Message: "Input harus berupa bilangan bulat",
			Data:    nil,
		})
	}

	err = repositories.GetOneNote(uint(IdNote), &onenote)
	if err != nil {
		return c.JSON(500, models.BaseResponse{
			Status:  false,
			Message: "Gagal menemukan note yang dicari",
			Data:    nil,
		})
	}
	return c.JSON(200, models.BaseResponse{
		Status:  true,
		Message: "Sukses mendapatkan note yang dicari",
		Data:    onenote,
	})

}

func ReadAllNote(c echo.Context) error {
	var notes []models.Note
	err := repositories.GetAllNotes(&notes)
	if err != nil {
		return c.JSON(500, models.BaseResponse{
			Status:  false,
			Message: "Gagal mendapatkan semua data note",
			Data:    nil,
		})
	}
	return c.JSON(200, models.BaseResponse{
		Status:  true,
		Message: "Sukses mendapatkan semua data note",
		Data:    notes,
	})
}

func DeleteOneNote(c echo.Context) error {
	var onenote models.Note
	id := c.Param("id")
	IdNote, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return c.JSON(400, models.BaseResponse{
			Status:  false,
			Message: "Input harus berupa bilangan bulat",
			Data:    nil,
		})
	}

	err = repositories.DeleteOneNote(uint(IdNote), &onenote)
	if err != nil {
		return c.JSON(500, models.BaseResponse{
			Status:  false,
			Message: "Gagal delete note yang dicari",
			Data:    nil,
		})
	}
	return c.JSON(200, models.BaseResponse{
		Status:  true,
		Message: "Berhasil delete note yang dicari",
		Data:    onenote,
	})
}
