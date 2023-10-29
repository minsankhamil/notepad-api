package repositories

import (
	"notepad-api/configs"
	"notepad-api/models"
)

func AddNote(note *models.Note) error {
	err := configs.DB.Create(&note).Error
	return err
}

func GetOneNote(idNote uint, onenote *models.Note) error {
	err := configs.DB.First(&onenote, idNote).Error
	return err
}

func GetAllNotes(notes *[]models.Note) error {
	err := configs.DB.Find(notes).Error
	return err
}

func DeleteOneNote(idNote uint, onenote *models.Note) error {
	err := configs.DB.Delete(&onenote, idNote).Error
	return err
}
