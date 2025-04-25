package models

import "gorm.io/gorm"

type Crop struct {
	gorm.Model
	Name               string `json:"name"`
	Description        string `json:"description"`
	Image              string `json:"image"`
	Duration           string `json:"duration"`
	Type               string `json:"type"`
	SoilType           string `json:"soil_type"`
	TempRange          string `json:"temp_range"`
	IsRainfallRequired bool   `json:"is_rainfall_required"`
	PhRange            string `json:"ph_range"`
	MarketValue        int    `json:"market_value"`
}

type Product struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

type NewsArticle struct {
	gorm.Model
	Title            string `json:"title"`
	ShortDescription string `json:"short_description"`
	Content          string `json:"content"`
	Tags             string `json:"tags"`
	Image            string `json:"image"`
	Date             string `json:"date"`
}

type HistoryItem struct {
	gorm.Model
	ProductID uint   `json:"product_id"`
	Quantity  int    `json:"quantity"`
	Type      string `json:"type"`
}

type User struct {
	gorm.Model
	Name           string `gorm:"size:255;not null;unique" json:"name"`
	Email          string `gorm:"size:255;not null;unique" json:"email"`
	HashedPassword string `gorm:"size:255;not null;"`
}

type UserRegister struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ProductSell struct {
	gorm.Model
	ProductID uint   `json:"product_id"`
	Quantity  int    `json:"quantity"`
	Type      string `json:"type"`
	SaleDate  string `json:"sale_date"`
	Address   string `json:"address"`
}
