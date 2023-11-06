package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin" // Import Gin instead of Fiber
	"github.com/moaabb/ecommerce/user_svc/domain/user"
	"github.com/moaabb/ecommerce/user_svc/infra/database/userdb"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	repository user.Repository
	l          *zap.Logger
}

func NewHandler(repo *userdb.Repository, z *zap.Logger) *UserHandler {
	return &UserHandler{
		repository: repo,
		l:          z,
	}
}

func (uh *UserHandler) GetAll(c *gin.Context) {
	uh.l.Info("Fetching Users on database")
	Users, err := uh.repository.GetAll()
	if err != nil {
		uh.l.Error("error while fetching Users", zap.Error(err))
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "could not fetch users"})
		return
	}

	c.JSON(http.StatusOK, Users)
}

func (uh *UserHandler) GetById(c *gin.Context) {
	uh.l.Info("getting User id...")
	userId, err := strconv.Atoi(c.Param("id")) // Use c.Param to get route parameters
	if err != nil {
		uh.l.Error("error getting userId", zap.Error(err))
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "could not fetch users"})
		return
	}

	uh.l.Info("Fetching User on database")
	user, err := uh.repository.GetById(uint(userId))
	if err != nil {
		uh.l.Error("error while fetching user", zap.Error(err))
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "could not fetch user"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (uh *UserHandler) Create(c *gin.Context) {
	var u user.User
	c.BindJSON(&u)

	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 12)
	if err != nil {
		uh.l.Error("error while generating user password hash", zap.Error(err))
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "could not create user"})
		return
	}

	u.Password = string(hash)
	uh.l.Info("Creating User on database")
	user, err := uh.repository.Create(u)
	if err != nil {
		uh.l.Error("error while creating user", zap.Error(err))
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "could not create user"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (uh *UserHandler) Update(c *gin.Context) {
	uh.l.Info("getting User id...")
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		uh.l.Error("error getting userId", zap.Error(err))
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "could not update users"})
		return
	}

	var u user.User
	c.BindJSON(&u)
	if u.Password == "" {
		uh.l.Info("getting user from db")
		user, err := uh.repository.GetById(uint(userId))
		if err != nil {
			uh.l.Error("error getting user from db", zap.Error(err))
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "could not update user"})
			return
		}

		u.Password = user.Password
	} else {
		hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 12)
		if err != nil {
			uh.l.Error("error while generating user password hash", zap.Error(err))
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "could not create user"})
			return
		}

		u.Password = string(hash)
	}

	uh.l.Info("updating User on database")
	user, err := uh.repository.Update(uint(userId), u)
	if err != nil {
		uh.l.Error("error while updating user", zap.Error(err))
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "could not update user"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (uh *UserHandler) Delete(c *gin.Context) {
	uh.l.Info("getting User id...")
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		uh.l.Error("error getting userId", zap.Error(err))
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "could not delete users"})
		return
	}

	uh.l.Info("deleting User on database")
	err = uh.repository.Delete(uint(userId))
	if err != nil {
		uh.l.Error("error while deleting user", zap.Error(err))
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "could not delete user"})
		return
	}

	c.Status(http.StatusNoContent)
}
