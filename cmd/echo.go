package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

const key = `somekey`

type RespObject struct {
	Message string `json:"message"`
}

type Response struct {
	Data interface{}
	Code int
}

func main() {
	e := echo.New()
	e.GET("/hello", handler, middle)
	//curl localhost:3000/hello
	//{"message":"Hello World"}
	e.Logger.Fatal(e.Start(":3000"))
}

func handler(c echo.Context) error {
	resp := RespObject{
		Message: "Hello",
	}
	setToContext(c, Response{
		Data: resp,
		Code: http.StatusOK,
	})

	return nil
}

func middle(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err != nil {
			log.Fatal(err)
		}

		r := getFromContext(c)
		if r.Code != http.StatusOK {
			log.Fatal(fmt.Printf("response code is %d", r.Code))
		}

		data, ok := r.Data.(RespObject)
		if !ok {
			log.Fatal("data is not RespObject")
		}

		data.Message = data.Message + ` World`

		return c.JSON(r.Code, data)
	}
}

func setToContext(c echo.Context, response Response) {
	c.Set(key, response)
}

func getFromContext(c echo.Context) Response {
	return c.Get(key).(Response)
}
