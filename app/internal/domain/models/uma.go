package models

import (
	"gorm.io/gorm"
)

type Uma struct {
	gorm.Model

	Name string `json:"name" form:"name"`

	BlueInheritanceID uint `json:"blue_inheritance_id" form:"blue_inheritance_id"`
	BlueInheritance *Uma `json:"blue_inheritance" form:"blue_inheritance"`
	RedInheritanceID uint `json:"red_inheritance_id" form:"red_inheritance_id"`
	RedInheritance *Uma `json:"red_inheritance" form:"red_inheritance"`

	Speed int `json:"speed" form:"speed"`
	Stamina int `json:"stamina" form:"stamina"`
	Power int `json:"power" form:"power"`
	Spirit int `json:"spirit" form:"spirit"`
	Intelligence int `json:"intelligence" form:"intelligence"`

	Turf int `json:"turf" form:"turf"`
	Dirt int `json:"dirt" form:"dirt"`

	Short int `json:"short" form:"short"`
	Mile int `json:"mile" form:"mile"`
	Middle int `json:"middle" form:"middle"`
	Long int `json:"long" form:"long"`


	Front int `json:"front" form:"front"` // 逃げ
	Stalker int `json:"stalker" form:"stalker"` // 先行
	Sotp int `json:"sotp" form:"sotp"` // 差し
	Offer int `json:"offer" form:"offer"` // 追込

	Point int `json:"point" form:"point"`

	BleuFactorID int `json:"blue_factor_id" form:"blue_factor_id"`
	BleuFactor *Factor `json:"blue_factor" form:"blue_factor"`
	RedFactorID int `json:"red_factor_id" form:"red_factor_id"`
	RedFactor *Factor `json:"red_factor" form:"red_factor"`
	WhiteFactors []Factor `json:"white_factors" form:"white_factors" gorm:"many2many:uma_wfactor"`
}
