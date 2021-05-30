package models

import (
	"fmt"

	"gorm.io/gorm"
)

type RaceResult struct {
	gorm.Model
	RaceID    uint   `json:"race_id" form:"race_id"`
	Race      *Race  `json:"-" form:"-"`
	Weather   string `json:"weather" form:"weather"`
	Condition string `json:"condition" form:"condition"`
	Tactics   string `json:"tactics" form:"tactics"`
	Result    uint   `json:"result" form:"result"`
	UmaID     uint   `json:"-"`
}

type Race struct {
	gorm.Model `json:"-"`
	Name       string `json:"name" form:"name"`
	Place      string `json:"place" form:"place"`
	Ground     string `json:"ground" form:"ground"`
	Length     uint   `json:"length" form:"length"`
	Course     string `json:"course" form:"course"`
	Turn       uint   `json:"turn" form:"turn"`
}

var (
	year = []string{"ジュニア級", "クラシック級", "シニア級", "ファイナルズ"}
	turn = []string{"前半", "後半"}
)

func (r Race) String() string {
	y := r.Turn / 1000
	m := (r.Turn % 1000) / 10
	t := r.Turn % 10
	return fmt.Sprintf(
		"%s : %s %s %dm %s (%s %d月%s)",
		r.Name,
		r.Place,
		r.Ground,
		r.Length,
		r.Course,
		year[y-1],
		m,
		turn[t-1],
	)
}
