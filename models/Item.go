// models/Item.go
package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model //gorm จะสร้าง ID
	Name       string
	Price      float64
}
