package structs

import "time"

type metadata interface {
	EpisodeMetadata | MovieMetadata | map[string]interface{}
}

type UploadInterface struct {
	EpMeta  EpisodeMetadata   `json:"episodeMetadata"`
	MoMeta  MovieMetadata     `json:"movieMetdatda"`
	PerMeta map[string]string `json:"performerMetadata"`
	Media   Media             `json:"media"`
	Name    string            `json:"name"` //全网同名，存储的路径 数据库中会lower()
}

type Media struct {
	Thumbnail string   `json:"thumbnail"`
	Poster    string   `json:"poster"`
	Gallery   []string `json:"gallery"`
	Trailer   string   `json:"trailer"`
}

type EpisodeMetadata struct {
	URl         string               `json:"url" bson:"url,omitempty"` //as an unique  identifer
	Name        string               `json:"name" bson:"name,omitempty"`
	Desc        string               `json:"desc" bson:"desc"`
	Series      *string              `json:"series" bson:"series"`
	ReleaseDate *time.Time           `json:"releaseDate" bson:"releaseDate"`
	Performers  []PerformerEssential `json:"performers" bson:"performers"`
	Runtime     *int                 `json:"runtime" bson:"runtime"`
	Code        *string              `json:"code" bson:"code"` //
	Tags        []string             `json:"tags" bson:"tags"`
}
type MovieMetadata struct {
	URl         string               `json:"url" bson:"url,omitempty"` //as an unique  identifer
	Name        string               `json:"name" bson:"name,omitempty"`
	Desc        string               `json:"desc" bson:"desc"`
	Series      *string              `json:"series" bson:"series"`
	Performers  []PerformerEssential `json:"performers" bson:"performers"`
	ReleaseDate *time.Time           `json:"releaseDate" bson:"releaseDate"`
	Runtime     *int                 `json:"runtime" bson:"runtime"`
	Code        *string              `json:"code" bson:"code"`
	Tags        []string             `json:"tags" bson:"tags"`
	Fellows     []string             `json:"fellows" bson:"fellows"` //episodes url
}

type PerformerEssential struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PerformerMeta struct { //only for SQL
	URl    string
	Name   string
	Others string //JSON stringfied string
}
