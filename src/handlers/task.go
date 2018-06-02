package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/FernandoCagale/serverless-go/src/error"
	"github.com/FernandoCagale/serverless-go/src/models"
	"github.com/FernandoCagale/serverless-go/src/render"
	"github.com/FernandoCagale/serverless-go/src/utils"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func FindAll(w http.ResponseWriter, r *http.Request) {
	db, err := utils.GetConnection(r)
	if err != nil {
		render.ResponseError(w, error.AddInternalServerError(err.Error()))
		return
	}

	tasks := []models.Task{}
	if err := db.Find(&tasks).Error; err != nil {
		render.ResponseError(w, error.AddInternalServerError(err.Error()))
		return
	}

	render.Response(w, tasks, http.StatusOK)
}

func FindById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		render.ResponseError(w, error.AddBadRequestError("Invalid task ID"))
		return
	}

	db, err := utils.GetConnection(r)
	if err != nil {
		render.ResponseError(w, error.AddInternalServerError(err.Error()))
		return
	}

	var task models.Task
	if err := db.Find(&task, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			render.ResponseError(w, error.AddNotFoundError("Task not found"))
		default:
			render.ResponseError(w, error.AddInternalServerError(err.Error()))
		}
		return
	}

	render.Response(w, task, http.StatusOK)
}

func UpdateById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		render.ResponseError(w, error.AddBadRequestError("Invalid task ID"))
		return
	}

	var task models.Task
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&task); err != nil {
		render.ResponseError(w, error.AddBadRequestError("Invalid request payload"))
		return
	}

	defer r.Body.Close()

	db, err := utils.GetConnection(r)
	if err != nil {
		render.ResponseError(w, error.AddInternalServerError(err.Error()))
		return
	}

	if err := db.Find(&models.Task{}, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			render.ResponseError(w, error.AddNotFoundError("Task not found"))
		default:
			render.ResponseError(w, error.AddInternalServerError(err.Error()))
		}
		return
	}

	if err := db.Save(&task).Error; err != nil {
		render.ResponseError(w, error.AddInternalServerError(err.Error()))
		return
	}

	render.Response(w, task, http.StatusOK)
}

func DeleteById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		render.ResponseError(w, error.AddBadRequestError("Invalid task ID"))
		return
	}

	db, err := utils.GetConnection(r)
	if err != nil {
		render.ResponseError(w, error.AddInternalServerError(err.Error()))
		return
	}

	var task models.Task
	if err := db.Find(&task, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			render.ResponseError(w, error.AddNotFoundError("Task not found"))
		default:
			render.ResponseError(w, error.AddInternalServerError(err.Error()))
		}
		return
	}

	if err := db.Delete(&task).Error; err != nil {
		render.ResponseError(w, error.AddInternalServerError(err.Error()))
		return
	}

	render.Response(w, map[string]string{"deleted": "true"}, http.StatusOK)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&task); err != nil {
		render.ResponseError(w, error.AddBadRequestError("Invalid request payload"))
		return
	}

	defer r.Body.Close()

	db, err := utils.GetConnection(r)
	if err != nil {
		render.ResponseError(w, error.AddInternalServerError(err.Error()))
		return
	}

	if err := db.Save(&task).Error; err != nil {
		render.ResponseError(w, error.AddInternalServerError(err.Error()))
		return
	}

	render.Response(w, task, http.StatusCreated)
}
