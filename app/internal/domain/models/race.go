package models

import (
	"fmt"

	"gorm.io/gorm"
)

type RaceResult struct {
	gorm.Model
	RaceID    uint   `json:"race_id" form:"race_id"`
	Race      *Race  `json:"race" form:"-"`
	Weather   string `json:"weather" form:"weather"`
	Condition string `json:"condition" form:"condition"`
	Tactics   string `json:"tactics" form:"tactics"`
	Result    uint   `json:"result" form:"result"`
	UmaID     uint   `json:"uma_id"`
}

type Race struct {
	gorm.Model
	Name       string `json:"name" form:"name"`
	Place      string `json:"place" form:"place"`
	Ground     string `json:"ground" form:"ground"`
	Length     uint   `json:"length" form:"length"`
	Course     string `json:"course" form:"course"`
	Turn       Turn   `json:"turn" form:"turn"`
}

type Turn uint

var (
	year = []string{"ジュニア級", "クラシック級", "シニア級", "ファイナルズ"}
	turn = []string{"前半", "後半"}
)

func (r Race) String() string {
	return fmt.Sprintf(
		"%s : %s %s %dm %s (%s)",
		r.Name,
		r.Place,
		r.Ground,
		r.Length,
		r.Course,
		r.Turn.String(),
	)
}

func (i Turn) String() string {
	y := i / 1000
	m := (i % 1000) / 10
	t := i % 10
	return fmt.Sprintf(
		"%s %d月%s",
		year[y-1],
		m,
		turn[t-1],
	)
}
