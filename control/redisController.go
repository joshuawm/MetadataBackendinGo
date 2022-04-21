package control

import "github.com/gofiber/fiber/v2"

////////////RedisRoom////////////////////
type RedisParam struct {
	Value string `query:"value"`
}
type RedisResponse struct {
	Result bool   `json:"result"`
	Err    string `json:"err"`
}

func RedisHandlePut(c *fiber.Ctx) error {
	param := new(RedisParam)
	c.QueryParser(param)
	result, e := BFilterMan.BloomFilterAdd(param.Value)
	var r RedisResponse
	if e == nil {
		r = RedisResponse{result, ""}
	} else {
		r = RedisResponse{result, e.Error()}
	}
	return c.JSON(r)
}

//exist
func RedisHandleGet(c *fiber.Ctx) error {
	param := new(RedisParam)
	c.QueryParser(param)
	result, e := BFilterMan.BloomFilterExist(param.Value)
	var r RedisResponse
	if e == nil {
		r = RedisResponse{result, ""}
	} else {
		r = RedisResponse{result, e.Error()}
	}
	return c.JSON(r)
}
