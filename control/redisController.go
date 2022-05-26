package control

import (
	"encoding/json"
	"net/http"
)

////////////RedisRoom////////////////////

type RedisResponse struct {
	Result bool   `json:"result"`
	Err    string `json:"err"`
}

func RedisHandlePut(w http.ResponseWriter, req *http.Request) {
	param := req.URL.Query().Get("value")
	var r RedisResponse
	if param == "" {
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte("no value inputed"))
		return
	}
	result, e := BFilterMan.BloomFilterAdd(param)
	if e == nil {
		r = RedisResponse{result, ""}
	} else {
		r = RedisResponse{result, e.Error()}
	}
	c, e := json.Marshal(r)
	if e != nil {
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte(e.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(c)
	return
}

//exist
func RedisHandleGet(w http.ResponseWriter, req *http.Request) {
	param := req.URL.Query().Get("value")
	var r RedisResponse
	if param == "" {
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte("no value inputed"))
		return
	}
	result, e := BFilterMan.BloomFilterExist(param)
	if e == nil {
		r = RedisResponse{result, ""}
	} else {
		r = RedisResponse{result, e.Error()}
	}
	c, e := json.Marshal(r)
	if e != nil {
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte(e.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(c)
	return
}
