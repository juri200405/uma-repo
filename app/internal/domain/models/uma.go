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

	Speed uint `json:"speed" form:"speed"`
	Stamina uint `json:"stamina" form:"stamina"`
	Power uint `json:"power" form:"power"`
	Spirit uint `json:"spirit" form:"spirit"`
	Intelligence uint `json:"intelligence" form:"intelligence"`

	Turf uint `json:"turf" form:"turf"`
	Dirt uint `json:"dirt" form:"dirt"`

	Short uint `json:"short" form:"short"`
	Mile uint `json:"mile" form:"mile"`
	Middle uint `json:"middle" form:"middle"`
	Long uint `json:"long" form:"long"`


	Front uint `json:"front" form:"front"` // 逃げ
	Stalker uint `json:"stalker" form:"stalker"` // 先行
	Sotp uint `json:"sotp" form:"sotp"` // 差し
	Offer uint `json:"offer" form:"offer"` // 追込

	Point uint `json:"point" form:"point"`

	BlueFactorID uint `json:"blue_factor_id" form:"blue_factor_id"`
	BlueFactor *Factor `json:"blue_factor" form:"blue_factor"`
	RedFactorID uint `json:"red_factor_id" form:"red_factor_id"`
	RedFactor *Factor `json:"red_factor" form:"red_factor"`
	WhiteFactors []Factor `json:"white_factors" form:"white_factors" gorm:"many2many:uma_wfactor"`

	RaceResults []RaceResult `json:"races" form:"races"`
}
