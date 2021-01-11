package middleware

import (
	"errors"
	"net/http"
	"strconv"
	"todo-api/api/auth"
	"todo-api/api/models"
	"todo-api/api/responses"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func SetMiddlewareAuthentication(next http.HandlerFunc, DB *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.ValidateToken(r)
		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, err)
			return
		}
		userId, err := auth.ExtractTokenID(r)
		user := models.User{}
		_, err = user.FindById(DB, int64(userId))
		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, err)
			return
		}

		next(w, r)
	}
}

func TodoOwner(next http.HandlerFunc, DB *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		err := auth.ValidateToken(r)
		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, err)
			return
		}

		vars := mux.Vars(r)
		id, err := strconv.ParseInt(vars["id"], 10, 64)
		userID, err := auth.ExtractTokenID(r)
		todo := models.Todo{}
		todoInfo, err := todo.Find(DB, uint(id))

		if todoInfo.UserID != userID {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("Todo Doesn't belongs to User"))
			return
		}
		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, err)
			return
		}
		next(w, r)
	}
}
