package api

import (
	"JurusanKu/src/config"
	db "JurusanKu/src/database/sqlc"
	"strings"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type register struct {
	Email string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
	Role string `form:"role" binding:"required"`
	Name string `form:"name" binding:"required"`
}

func (server *Server) Register(ctx *gin.Context) {
	// get request body
	var req register
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// register user
	args := db.RegisterUserParams{
		Email: req.Email,
		Password: string(hashedPassword),
		Role: req.Role,
		Name: req.Name,
	}

	newUser, err := server.store.RegisterUser(ctx, args)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates")  {
			ctx.JSON(http.StatusConflict, gin.H{"error": "username already exists"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, newUser.CreatedAt)
}

type login struct {
	Email string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func (server *Server) Login(ctx *gin.Context) {
	// get request body
	var req login
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// get user
	user, err := server.store.GetUserByEmail(ctx, req.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "invalid username"})
		return
	}

	// compare password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid password"})
		return
	}

	// generate token
	token, payload, err := config.CreateToken(user.Email, user.Role, server.config.JWT_EXP, server.config.JWT_KEY)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	ctx.SetCookie("token", token, 3600, "/", "localhost", false, true)
	log.Println(ctx.Cookie("token"))
	ctx.JSON(http.StatusOK, payload)
}

type getProfileRequest struct {
	Email string `form:"email" binding:"required"`
}

func (server *Server) getProfile(ctx *gin.Context) {
	var req getProfileRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := server.store.GetUserByEmail(ctx, req.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

type updateProfileRequest struct {
	Email string `form:"email" binding:"required"`
	Name string `form:"name" binding:"required"`
	Nick string `form:"nick" binding:"required"`
}

func (server *Server) updateProfile(ctx *gin.Context) {
	var req updateProfileRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	args := db.UpdateProfileParams{
		Name: req.Name,
		Nick: req.Nick,
		Email: req.Email,
	}

	user, err := server.store.UpdateProfile(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update profile"})
		return
	}
	ctx.JSON(http.StatusOK, user)
}