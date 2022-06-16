package controllers

import (
	"crypto/sha256"
	"encoding/base64"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/udayangaac/shipments-service/api/schema"
	"github.com/udayangaac/shipments-service/config"
	"github.com/udayangaac/shipments-service/repo"
	"github.com/udayangaac/shipments-service/repo/entity"
)

type UserController struct {
	UserRepo repo.UserRepo
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var req schema.CreateUserReq
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": true,
			"msg":   err.Error(),
		})
		return
	}
	if err = c.UserRepo.Save(ctx.Request.Context(), &entity.User{
		Name:     req.Name,
		Password: req.Password,
		Email:    encryptPassword(req.Email),
	}); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": true,
			"msg":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"error": false,
		"msg":   "successfully created the user",
	})
}

func (c *UserController) LoginUser(ctx *gin.Context) {
	var (
		req         schema.LoginReq
		user        entity.User
		tokenString string
	)
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": true,
			"msg":   err.Error(),
		})
		return
	}

	user, err = c.UserRepo.FindByEmail(ctx.Request.Context(), req.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": true,
			"msg":   err.Error(),
		})
		return
	}

	expirationTime := time.Now().Add(time.Second * time.Duration(config.ServerConf.Jwt.Duration))
	claims := jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err = token.SignedString(config.ServerConf.Jwt.Key)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"msg":   err.Error(),
		})
	}

	if user.Password == encryptPassword(user.Password) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"“UID”:":          user.ID,
			"“access_token”:": tokenString,
		})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"msg":   "invalid credentials",
		})
		return
	}
}

func encryptPassword(pwd string) string {
	h := sha256.Sum256([]byte(pwd))
	return base64.StdEncoding.EncodeToString(h[:])
}
