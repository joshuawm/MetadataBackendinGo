package control

import (
	"backman/database/gromimplement"
	"backman/structs"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/slices"
)

var schemas = make([]string, 0)

func GormHandler(RawData structs.UploadInterface) error {
	if RawData.EpMeta.URl != "" {
		// GormDB.Exec()
		return gromimplement.InsertEpisode(GormDB, RawData.EpMeta, RawData.Name)
	} else if RawData.MoMeta.URl != "" {
		return gromimplement.InsertMovie(GormDB, RawData.MoMeta, RawData.Name)
	} else if RawData.PerMeta["URL"] != "" {
		u := RawData.PerMeta["URL"]
		n := RawData.PerMeta["Name"]
		delete(RawData.PerMeta, "Name")
		delete(RawData.PerMeta, "URL")
		info, err := json.Marshal(RawData.PerMeta)
		if err != nil {
			log.Fatal(err)
		}
		d := structs.PerformerMeta{URl: u, Name: n, Others: string(info)}
		return gromimplement.InsertPerformer(GormDB, d, RawData.Name)
	} else {
		return errors.New("判断失败，没有符合条件的数据传入到此")
	}
}

/////////////Schema/////////////
func UpdateSchemas() {
	GormDB.Raw("SELECT nspname FROM pg_catalog.pg_namespace;").Scan(&schemas)
}

func AllSchemaHandle(c *fiber.Ctx) error {
	if len(schemas) == 0 {
		UpdateSchemas()
	}
	r, err := json.Marshal(schemas)
	if err != nil {
		c.SendStatus(http.StatusBadRequest)
		return c.Send([]byte(err.Error()))
	}
	return c.Send(r)
}

type schemaQueryString struct {
	Name string
}

func CreateSchemaHandle(c *fiber.Ctx) error {
	if len(schemas) == 0 {
		UpdateSchemas()
	}
	q := schemaQueryString{}
	err := c.QueryParser(&q)
	if err != nil {
		c.SendStatus(http.StatusBadRequest)
		return c.Send([]byte(err.Error()))
	}
	schemaName := q.Name
	if schemaName == "" {
		c.SendStatus(http.StatusBadRequest)
		return c.Send([]byte("name值为空！"))
	}
	if slices.Contains(schemas, schemaName) {
		c.SendStatus(http.StatusBadRequest)
		return c.Send([]byte(fmt.Sprintf("%s这个schema已经存在！", schemaName)))
	}
	err = gromimplement.CreateSchema(schemaName, GormDB)
	if err != nil {
		c.SendStatus(http.StatusBadRequest)
		return c.Send([]byte(err.Error()))
	} else {
		return c.Send([]byte("sucess!"))
	}
}
