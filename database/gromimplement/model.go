package gromimplement

import (
	"gorm.io/gorm"
)

type Episode struct {
	URL         string  `gorm:"unique;not null;type:varchar(255)"`
	Name        string  `gorm:"type:varchar(255)"`
	Desc        string  `gorm:"type:varchar(1024)"`
	Series      *string `gorm:"type:varchar(255)"`
	ReleaseDate *int
	Runtime     *int
	Code        *string `gorm:"type:varchar(255)"`
	Tags        *string `gorm:"type:varchar(255)"`
	gorm.Model
}

type Movie struct {
	URL         string  `gorm:"unique;not null;type:varchar(255)"`
	Name        string  `gorm:"type:varchar(255)"`
	Desc        string  `gorm:"type:varchar(1024)"`
	Series      *string `gorm:"type:varchar(255)"`
	ReleaseDate *int
	Runtime     *int
	Code        *string `gorm:"type:varchar(255)"`
	Tags        *string `gorm:"type:varchar(255)"`
	gorm.Model
}

type Performer struct {
	URL  string  `gorm:"uniqueIndex;not null;type:varchar(255)"`
	Name string  `gorm:"type:varchar(255)"`
	Info *string `gorm:"type:varchar(255)"`
	gorm.Model
}

type Episodes2Movie struct {
	MovieID   int
	Movie     Movie
	EpisodeID int
	Episode   Episode
	gorm.Model
}

type Episode2Performers struct {
	EpisodeID   int
	Episode     Episode
	PerformerID int
	Performer   Performer
	gorm.Model
}
