package seeds

import (
	"clean/internal/product/entities"
	"log"
	"time"

	"gorm.io/gorm"
)

func SeedProduct(db *gorm.DB, categoryIds []string) {
	for _, product := range []entities.ProductEntity{
		{CategoriesId: categoryIds[0], Name: "Bois d'argent", Price: 59.00, Quantity: 50, Created_at: time.Now()},
		{CategoriesId: categoryIds[1], Name: "Bois de roses", Price: 69.00, Quantity: 20, Created_at: time.Now()},
		{CategoriesId: categoryIds[2], Name: "Bois des bois", Price: 55.00, Quantity: 45, Created_at: time.Now()},
		{CategoriesId: categoryIds[3], Name: "Bois d'or", Price: 49.00, Quantity: 15, Created_at: time.Now()},
	} {
		if err := db.Table("products").Create(&product).Error; err != nil {
			log.Fatalf("Error inserting Products : %v", err)
		}
	}
}
