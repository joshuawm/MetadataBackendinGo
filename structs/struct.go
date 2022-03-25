package structs

type metadata interface {
	EpisodeMetadata | MovieMetadata | map[string]interface{}
}

type UploadInterface struct {
	EpMeta  EpisodeMetadata        `json:"epmeta"`
	MoMeta  MovieMetadata          `json:"moviemeta"`
	PerMeta map[string]interface{} `json:"permeta"`
	Media   Media                  `json:"media"`
	Name    string                 `json:"name"` //全网同名，存储的路径
}

type Media struct {
	Thumbnail string   `json:"thumbnail`
	Poster    string   `json:"thumbnail"`
	Gallery   []string `json:"gallery"`
	Trailer   string   `json:"trailer"`
}

type EpisodeMetadata struct {
	URl        string   `json:"url" bson:"url,omitempty"` //as an unique  identifer
	Name       string   `json:"name" bson:"name,omitempty"`
	Desc       string   `json:"desc" bson:"desc"`
	Series     string   `json:"series" bson:"series"`
	Performers []string `json:"performers" bson:"performers"`
	Runtime    int64    `json:"runtime" bson:"runtime"`
	Tags       []string `json:"tags" bson:"tags"`
}
type MovieMetadata struct {
	URl        string   `json:"url" bson:"url,omitempty"` //as an unique  identifer
	Name       string   `json:"name" bson:"name,omitempty"`
	Desc       string   `json:"desc" bson:"desc"`
	Series     string   `json:"series" bson:"series"`
	Performers []string `json:"performers" bson:"performers"`
	Runtime    int64    `json:"runtime" bson:"runtime"`
	Tags       []string `json:"tags" bson:"tags"`
	Fellows    []string `json:"fellows" bson:"fellows"` //episodes url
}

type PerformerMeta struct { //only for SQL
	URl    string
	Name   string
	Others string //JSON stringfied string
}
