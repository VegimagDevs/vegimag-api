package handlers

import (
	"fmt"
	"net/http"

	"github.com/VegimagDevs/vegimag-api/storage"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserJSON struct {
	Email    string `json:"email" binding:"required,email"`
	Username string `json:"username" binding:"required,lte=50"`
	Password string `json:"password" binding:"required"`
}

// @Summary Create an user
// @Accept json
// @Produce json
// @Param user body CreateUserJSON true "The user to create"
// @Router /users [post]
func (handlers *Handlers) CreateUser(ctx *gin.Context) {
	user := new(CreateUserJSON)
	if err := ctx.ShouldBindJSON(user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userByEmail, err := handlers.config.Storage.GetUserByEmail(user.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if userByEmail != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "this email is already used",
		})
		return
	}

	userByUsername, err := handlers.config.Storage.GetUserByUsername(user.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if userByUsername != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "this username is already used",
		})
		return
	}

	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	userId := uuid.NewV4().String()

	storableUser := &storage.User{
		Id:             userId,
		Email:          user.Email,
		Username:       user.Username,
		HashedPassword: hashedPassword,
	}

	if err := handlers.config.Storage.CreateUser(storableUser); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"userId": userId,
	})
}

// @Summary Get a new validation token
// @Accept json
// @Produce json
// @Router /users/validation-token [get]
func (handlers Handlers) GetValidationToken(ctx *gin.Context) {

}

// @Summary Validate an user
// @Accept json
// @Produce json
// @Router /users/validate [post]
func (handlers Handlers) ValidateUser(ctx *gin.Context) {

}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("error hashing the password: %w", err)
	}

	return string(bytes), nil
}

func checkPassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
