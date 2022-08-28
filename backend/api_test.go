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

func TestGetTodos(t *testing.T) {
	Convey("Get Todos", t, func() {

		repository := GetCleanTestRepository()
		service := NewService(repository)
		api := NewApi(&service)

		todo1 := models.Todo{
			ID:          GenerateUUID(8),
			Name:        "baslik1",
			IsCompleted: false,
			CreatedAt:   time.Now().UTC().Round(time.Second),
			UpdatedAt:   time.Now().UTC().Round(time.Second),
		}

		todo2 := models.Todo{
			ID:          GenerateUUID(8),
			Name:        "baslik2",
			IsCompleted: false,
			CreatedAt:   time.Now().UTC().Round(time.Second),
			UpdatedAt:   time.Now().UTC().Round(time.Second),
		}

		repository.CreateTodo(todo1)
		repository.CreateTodo(todo2)

		Convey("When the get request sent ", func() {
			app := SetupApp(&api)
			req, _ := http.NewRequest(http.MethodGet, "/todos", nil)

			resp, err := app.Test(req)
			So(err, ShouldBeNil)

			Convey("Then status code should be 200", func() {
				So(resp.StatusCode, ShouldEqual, fiber.StatusOK)
			})
			Convey("Then the reques should return all todos", func() {
				actualResult := []models.Todo{}
				actualResponseBody, _ := ioutil.ReadAll(resp.Body)
				err := json.Unmarshal(actualResponseBody, &actualResult)
				So(err, ShouldBeNil)

				So(actualResult, ShouldHaveLength, 2)
				So(actualResult[0].ID, ShouldEqual, todo1.ID)
				So(actualResult[0].Name, ShouldEqual, todo1.Name)
				So(actualResult[0].IsCompleted, ShouldEqual, todo1.IsCompleted)
				So(actualResult[0].UpdatedAt, ShouldEqual, todo1.UpdatedAt)
				So(actualResult[0].CreatedAt, ShouldEqual, todo1.CreatedAt)
				So(actualResult[1].ID, ShouldEqual, todo2.ID)
				So(actualResult[1].Name, ShouldEqual, todo2.Name)
				So(actualResult[1].IsCompleted, ShouldEqual, todo2.IsCompleted)
				So(actualResult[1].UpdatedAt, ShouldEqual, todo2.UpdatedAt)
				So(actualResult[1].CreatedAt, ShouldEqual, todo2.CreatedAt)

			})

		})

	})

}

// Tekil Todo Get Etme Method Get
func TestGetTodo(t *testing.T) {
	Convey("Get todo", t, func() {

		repository := GetCleanTestRepository()
		service := NewService(repository)
		api := NewApi(&service)

		todo1 := models.Todo{
			ID:          GenerateUUID(8),
			Name:        "baslik2",
			IsCompleted: false,
			CreatedAt:   time.Now().UTC().Round(time.Second),
			UpdatedAt:   time.Now().UTC().Round(time.Second),
		}
		repository.CreateTodo(todo1)

		Convey("When the get request sent ", func() {
			app := SetupApp(&api)
			req, _ := http.NewRequest(http.MethodGet, "/todo/"+todo1.ID, nil)
			resp, err := app.Test(req, 30000)

			So(err, ShouldBeNil)

			Convey("Then status code should be 200", func() {
				So(resp.StatusCode, ShouldEqual, fiber.StatusOK)
			})

			Convey("Then product should be returned", func() {
				actualResult := models.Todo{}
				actualRespBody, _ := ioutil.ReadAll(resp.Body)
				err := json.Unmarshal(actualRespBody, &actualResult)

				So(err, ShouldBeNil)

				So(actualResult.ID, ShouldEqual, todo1.ID)
				So(actualResult.Name, ShouldEqual, todo1.Name)
				So(actualResult.IsCompleted, ShouldEqual, todo1.IsCompleted)
				So(actualResult.UpdatedAt, ShouldEqual, todo1.UpdatedAt)
				So(actualResult.CreatedAt, ShouldEqual, todo1.CreatedAt)
			})
		})
	})

}

// Tekil Todo Update Etme Method Put
func TestUpdateTodo(t *testing.T) {
	Convey("Update todo", t, func() {

		repository := GetCleanTestRepository()
		service := NewService(repository)
		api := NewApi(&service)

		todo1 := models.Todo{
			ID:          GenerateUUID(8),
			Name:        "baslik2",
			IsCompleted: false,
			CreatedAt:   time.Now().UTC().Round(time.Second),
			UpdatedAt:   time.Now().UTC().Round(time.Second),
		}
		repository.CreateTodo(todo1)

		Convey("when the put request sent", func() {
			app := SetupApp(&api)

			todo2 := models.TodoDTO{
				Name:        "batu",
				IsCompleted: false,
				CreatedAt:   time.Now().UTC().Round(time.Second),
			}
			reqBody, err := json.Marshal(todo2)

			So(err, ShouldBeNil)
			req, _ := http.NewRequest(http.MethodPut, "/todos/"+todo1.ID, bytes.NewReader(reqBody))
			req.Header.Add("Content-Type", "application/json")

			resp, err := app.Test(req, 30000)

			So(err, ShouldBeNil)

			Convey("then status should be 200", func() {
				So(resp.StatusCode, ShouldEqual, fiber.StatusOK)
			})
			Convey("Then product should be updated", func() {
				actualResult := models.Todo{}

				respBody, _ := ioutil.ReadAll(resp.Body)

				err = json.Unmarshal(respBody, &actualResult)
				So(err, ShouldBeNil)
				So(actualResult.ID, ShouldEqual, todo1.ID)
				So(actualResult.IsCompleted, ShouldEqual, false)
				So(actualResult.CreatedAt, ShouldEqual, todo2.CreatedAt)

			})
		})
	})
}

// Tekil Todo Delete Etme Method Delete
func TestDeleteTodo(t *testing.T) {
	Convey("Delete todo that user wants", t, func() {

		repository := GetCleanTestRepository()
		service := NewService(repository)
		api := NewApi(&service)

		todo1 := models.Todo{
			ID:          GenerateUUID(8),
			Name:        "baslik2",
			IsCompleted: false,
			CreatedAt:   time.Now().UTC().Round(time.Second),
			UpdatedAt:   time.Now().UTC().Round(time.Second),
		}

		repository.CreateTodo(todo1)

		Convey("When the delete request sent ", func() {
			app := SetupApp(&api)

			req, _ := http.NewRequest(http.MethodDelete, "/todos/"+todo1.ID, nil)
			resp, err := app.Test(req, 30000)
			So(err, ShouldBeNil)

			Convey("Then status code should be 204", func() {
				So(resp.StatusCode, ShouldEqual, fiber.StatusNoContent)
			})

			Convey("Then phone should be deleted", func() {
				todos, _ := repository.GetTodos()
				So(todos, ShouldHaveLength, 0)
				So(todos, ShouldResemble, []models.Todo{})
			})
		})
	})
}

// Tekil Todo Ekleme Etme Method Post
func TestAddTodo(t *testing.T) {
	Convey("Add todo", t, func() {
		repository := GetCleanTestRepository()
		service := NewService(repository)
		api := NewApi(&service)

		todo3 := models.Todo{
			/* ID:          GenerateUUID(8), */
			Name:        "baslik2",
			IsCompleted: false,
			CreatedAt:   time.Now().UTC().Round(time.Second),
			UpdatedAt:   time.Now().UTC().Round(time.Second),
		}

		Convey("when the post request send", func() {

			reqBody, err := json.Marshal(todo3)

			req, _ := http.NewRequest(http.MethodPost, "/todos", bytes.NewReader(reqBody))
			req.Header.Add("Content-Type", "application/json")
			req.Header.Set("Content-Length", strconv.Itoa(len(reqBody)))

			app := SetupApp(&api)
			resp, err := app.Test(req, 30000)
			So(err, ShouldBeNil)

			Convey("Then status code should be 201", func() {
				So(resp.StatusCode, ShouldEqual, fiber.StatusCreated)
			})

			Convey("Then Added todo should return", func() {
				actualResult, err := repository.GetTodo(todo3.ID)

				So(err, ShouldBeNil)
				/* So(actualResult, ShouldHaveLength, 1) */
				So(actualResult, ShouldNotBeNil)
				/* So(actualResult.ID, ShouldEqual, todo3.ID) */
				So(actualResult.Name, ShouldEqual, todo3.Name)
				So(actualResult.IsCompleted, ShouldEqual, todo3.IsCompleted)

			})
		})
	})
}
