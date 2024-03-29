package control

import (
	"backman/database/gromimplement"
	"backman/structs"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

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

func AllSchemaHandle(w http.ResponseWriter, req *http.Request) {
	if len(schemas) == 0 {
		UpdateSchemas()
	}
	r, err := json.Marshal(schemas)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(r)
	return
}

type schemaQueryString struct {
	Name string
}

func CreateSchemaHandle(w http.ResponseWriter, req *http.Request) {
	if len(schemas) == 0 {
		UpdateSchemas()
	}
	q := req.URL.Query().Get("value")
	var err error = nil
	if q == "" {
		err = errors.New("no value inpouted")
	}
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	schemaName := strings.ToLower(q)
	if schemaName == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("name值为空！"))
		return
	}
	if slices.Contains(schemas, schemaName) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("%s这个schema已经存在！", schemaName)))
		return
	}
	err = gromimplement.CreateSchema(schemaName, GormDB)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	} else {
		w.Write([]byte("sucess!"))
		return
	}
}
