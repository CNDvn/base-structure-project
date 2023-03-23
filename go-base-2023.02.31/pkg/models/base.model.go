package models

type Base struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt uint `gorm:"autoCreateTime:milli; UNSIGNED"`
	UpdatedAt uint `gorm:"autoUpdateTime:milli; UNSIGNED"`
	DeletedAt uint `gorm:"UNSIGNED"`
}
