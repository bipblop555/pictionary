package seeds

import (
	"clean/internal/category/entities"
	"log"

	"gorm.io/gorm"
)

func SeedCategory(db *gorm.DB) []string {
	var categoryIds []string

	for _, category := range []entities.CategorieEntity{
		{Name: "Hommes"},
		{Name: "Femmes"},
		{Name: "Fruité"},
		{Name: "Forts"},
	} {
		if err := db.Table("categories").Create(&category).Error; err != nil {
			log.Fatalf("Error inserting categories : %v", err)
		}
		categoryIds = append(categoryIds, category.Id)
	}
	return categoryIds
}
