package gromimplement

import (
	"backman/structs"
	"encoding/json"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func CreateSchema(name string, db *gorm.DB) error {
	r := db.Exec(fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS  %s", name))
	if r.Error != nil {
		return r.Error
	}
	err := db.Table(fmt.Sprintf("%s.episodes", name)).AutoMigrate(&Episode{})
	if err != nil {
		return err
	}
	err = db.Table(fmt.Sprintf("%s.movies", name)).AutoMigrate(&Movie{})
	if err != nil {
		return err
	}
	err = db.Table(fmt.Sprintf("%s.performers", name)).AutoMigrate(&Performer{})
	if err != nil {
		return err
	}
	err = db.Table(fmt.Sprintf("%s.episodes2movie", name)).AutoMigrate(&Episodes2Movie{})
	if err != nil {
		return err
	}
	err = db.Table(fmt.Sprintf("%s.episode2performers", name)).AutoMigrate(&Episode2Performers{})
	if err != nil {
		return err
	}
	return nil
}

func InsertEpisode(db *gorm.DB, data structs.EpisodeMetadata, name string) error {
	//preprocess

	return db.Transaction(func(tx *gorm.DB) error {
		tags, err := json.Marshal(data.Tags)
		if err != nil {
			return err
		}
		episode := Episode{URL: data.URl, Name: data.Name, Desc: data.Desc, Series: cString(data.Series), ReleaseDate: cInt(data.ReleaseDate), Runtime: cInt(data.Runtime), Tags: cString(string(tags)), Code: cString(data.Code)}
		if err != nil {

			return err
		}
		result := tx.Table(fmt.Sprintf("%s.episodes", name)).Create(&episode)
		if result.Error != nil {
			return result.Error
		}
		//Get Performer ID
		for _, p := range data.Performers {
			per := Performer{Name: p.Name, URL: p.URL}
			result := db.Table(fmt.Sprintf("%s.performers", name)).First(&per)
			if result.RowsAffected > 1 {
				return errors.New("这是一条不靠谱的error，检测到RowAffected数值大于1，具体请在源代码中搜索这个报错，无关请删除这个报错")
			}
			result = db.Table(fmt.Sprintf("%s.episode2performers", name)).Create(&Episode2Performers{EpisodeID: int(episode.ID), PerformerID: int(per.ID)})
			if result.Error != nil {
				return result.Error
			}
		}
		return nil
	})
}

func InsertMovie(db *gorm.DB, data structs.MovieMetadata, name string) error {
	return db.Transaction(func(tx *gorm.DB) error {
		tags, err := json.Marshal(data.Tags)
		if err != nil {
			return err
		}
		movie := Movie{URL: data.URl, Name: data.Name, Desc: data.Desc, Series: cString(data.Series), ReleaseDate: cInt(data.ReleaseDate), Runtime: cInt(data.Runtime), Tags: cString(string(tags)), Code: cString(data.Code)}
		result := db.Table(fmt.Sprintf("%s.movies", name)).Create(&movie)
		if result.Error != nil {
			return result.Error
		}
		for _, epURL := range data.Fellows {
			Episode := Episode{URL: epURL}
			result := db.Table(fmt.Sprintf("%s.episodes", name)).First(&Episode)
			if result.Error != nil {
				return result.Error
			}
			result = db.Table(fmt.Sprintf("%s.episodes2movie", name)).Create(&Episodes2Movie{EpisodeID: int(Episode.ID), MovieID: int(movie.ID)})
			if result.Error != nil {
				return result.Error
			}
		}
		return nil
	})
}

func InsertPerformer(db *gorm.DB, data structs.PerformerMeta, name string) error {
	return db.Transaction(func(tx *gorm.DB) error {
		per := Performer{URL: data.URl, Name: data.Name, Info: cString(data.Others)}
		result := db.Table(fmt.Sprintf("%s.performers", name)).Create(&per)
		if result.Error != nil {
			return result.Error
		}
		return nil
	})
}

func cString(v string) *string {
	if v == "" || v == "[]" {
		return nil
	} else {
		return &v
	}
}

func cInt(v int) *int {
	if v == 0 {
		return nil
	} else {
		return &v
	}
}

func main() {
	fmt.Println("gorm main exec!")
}
