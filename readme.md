# Document

## API Usage
所有的接口将仿照RESTFul的方式进行设计，也就是/api/catalogue/function
### **Upload**
将爬取的数据上传到服务端进行处理的接口，将包含数据入库（ NOSQL(MongoDb) SQL(MariaDB) ），媒体数据上传到S3服务商(backblaze ...etc) 和Onedrive进行保存 \

>`POST` &nbsp; `/api/v1/upload`: \
```go
type UploadInterface struct {
	EpMeta  EpisodeMetadata   `json:"epmeta"`
	MoMeta  MovieMetadata     `json:"moviemeta"`
	PerMeta map[string]string `json:"permeta"`
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
	Series      string               `json:"series" bson:"series"`
	ReleaseDate int                  `json:"releaseDate" bson:"releaseDate"`
	Performers  []PerformerEssential `json:"performers" bson:"performers"`
	Runtime     int                  `json:"runtime" bson:"runtime"`
	Code        string               `json:"code" bson:"code"` //
	Tags        []string             `json:"tags" bson:"tags"`
}
type MovieMetadata struct {
	URl         string               `json:"url" bson:"url,omitempty"` //as an unique  identifer
	Name        string               `json:"name" bson:"name,omitempty"`
	Desc        string               `json:"desc" bson:"desc"`
	Series      string               `json:"series" bson:"series"`
	Performers  []PerformerEssential `json:"performers" bson:"performers"`
	ReleaseDate int                  `json:"releaseDate" bson:"releaseDate"`
	Runtime     int                  `json:"runtime" bson:"runtime"`
	Code        string               `json:"code" bson:"code"`
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


```

### **Redis**
Redis操作接口
#### Redis BloomFilter 
URL去重检测 \
接口：\
`GET` &nbsp; `/api/v1/redis/bf/exist?value={string}` \
使用bloomfilter检测value传入的string是否存在
`POST` &nbsp; `/api/v1/redis/bf/put?value={string}` \
使用bloomfilter放入当前的value，若存在会报错并且返回status为false

##### Response
```go
type RedisResponse struct {
	Result bool   `json:"result"`
	Err    string `json:"err"`
}
```
 

## 后台架构
### SQL数据库

