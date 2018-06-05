package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	errors "github.com/FernandoCagale/serverless-go/src/error"
	"github.com/FernandoCagale/serverless-go/src/models"
	"github.com/FernandoCagale/serverless-go/src/render"
	"github.com/FernandoCagale/serverless-go/src/utils"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

func Login(w http.ResponseWriter, r *http.Request) {
	auth := new(models.Auth)

	db, err := utils.GetConnection(r)
	if err != nil {
		render.ResponseError(w, errors.AddInternalServerError(err.Error()))
		return
	}

	var authBind models.Auth
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&authBind); err != nil {
		render.ResponseError(w, errors.AddBadRequestError("Invalid request payload"))
		return
	}

	defer r.Body.Close()

	if err := db.Where("username = ?", authBind.Username).First(&auth).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			render.ResponseError(w, errors.AddUnauthorizedError("Username or password invalid"))
		default:
			render.ResponseError(w, errors.AddInternalServerError(err.Error()))
		}
		return
	}

	if valid := auth.ValidatePassword(authBind.Password); !valid {
		render.ResponseError(w, errors.AddUnauthorizedError("Username or password invalid"))
		return
	}

	token, err := createJwtToken(auth)
	if err != nil {
		render.ResponseError(w, errors.AddInternalServerError(err.Error()))
		return
	}

	render.Response(w, map[string]string{"token": token}, http.StatusOK)
}

func createJwtToken(auth *models.Auth) (string, error) {
	claims := models.JwtClaims{
		auth.Username,
		jwt.StandardClaims{
			Id:        strconv.Itoa(auth.ID),
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	return rawToken.SignedString([]byte("secret"))
}
