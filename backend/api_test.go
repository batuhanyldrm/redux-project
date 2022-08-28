package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"
	"time"

	"example.com/greetings/models"
	"github.com/gofiber/fiber/v2"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCreateTodo(t *testing.T) {

	Convey("Get Todos", t, func() {
		repository := NewRepository()
		service := NewService(repository)
		api := NewApi(&service)

		todo1 := models.TodoDTO{
			Name:        "Faruk",
			IsCompleted: true,
			CreatedAt:   time.Now().UTC().Round(time.Second),
		}

		Convey("When the post request send", func() {
			reqBody, err := json.Marshal(todo1)

			req, _ := http.NewRequest(http.MethodPost, "/todos", bytes.NewReader(reqBody))
			req.Header.Add("Content-Type", "application/json")
			req.Header.Set("Content-Length", strconv.Itoa(len(reqBody)))

			app := SetupApp(&api)
			resp, _ := app.Test(req, 3000)
			So(err, ShouldBeNil)

			Convey("Then status code Should be 201", func() {
				So(resp.StatusCode, ShouldEqual, fiber.StatusCreated)
			})

			Convey("Then post todo should return", func() {
				actualResult := models.Todo{}
				respBody, _ := ioutil.ReadAll(resp.Body)

				err = json.Unmarshal(respBody, &actualResult)
				So(err, ShouldBeNil)
				So(actualResult, ShouldNotBeNil)

				So(actualResult.Name, ShouldEqual, todo1.Name)
				So(actualResult.IsCompleted, ShouldEqual, todo1.IsCompleted)
				So(actualResult.CreatedAt, ShouldEqual, todo1.CreatedAt)
			})
		})
	})
}
