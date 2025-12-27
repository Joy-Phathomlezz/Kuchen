package repository

import (
	"github.com/Joy-Phathomlezz/Kuchen/internals/models"
	"gorm.io/gorm"
)

func CreateCake(db *gorm.DB, cake *models.Cake) error {
	// Use the provided DB instance to create a new cake record in the database
	return db.Create(cake).Error
}

func GetCakes(db *gorm.DB) ([]models.Cake, error) {
	var cakes []models.Cake
	// Use the provided DB instance to retrieve all cake records from the database
	err := db.Find(&cakes).Error
	return cakes, err
}

func GetCakeByID(db *gorm.DB, id uint) (*models.Cake, error) {
	var cake models.Cake
	// Use the provided DB instance to retrieve a single cake record by ID from the database
	err := db.First(&cake, id).Error
	return &cake, err
}

func UpdateCake(db *gorm.DB, id uint, updates map[string]interface{}) error {
	// Use the provided DB instance to update specific fields of the cake record by ID
	return db.Model(&models.Cake{}).Where("id = ?", id).Updates(updates).Error
}

func DeleteCake(db *gorm.DB, id uint) error {
	// Use the provided DB instance to delete the cake record by ID from the database
	return db.Delete(&models.Cake{}, id).Error
}
