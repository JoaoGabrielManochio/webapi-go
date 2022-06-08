package user

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/JoaoGabrielManochio/webapi-go/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// IUserController : interface of user business
type IUserController interface {
	CreateUser(c *gin.Context)
	GetUser(c *gin.Context)
	GetUsers(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

// UserController : struct of user Controller
type UserController struct {
	userBusiness IUserBusiness
}

// NewUserController : create a new user Controller
func NewUserController(userBusiness IUserBusiness) IUserController {
	return &UserController{userBusiness}
}

func (u *UserController) GetUser(c *gin.Context) {

	id := c.Param("id")

	newId, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{"error": "ID has to be integer", "code": strconv.Itoa(http.StatusBadRequest)})
		return
	}

	if newId <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{"error": "ID canot be 0,", "code": strconv.Itoa(http.StatusBadRequest)})
		return
	}

	statusCode, user, err := u.userBusiness.GetUser(newId)

	if err != nil {
		c.AbortWithStatusJSON(statusCode, map[string]string{"error": err.Error(), "code": strconv.Itoa(statusCode)})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (u *UserController) CreateUser(c *gin.Context) {

	user := &models.User{}
	if err := c.ShouldBindJSON(user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{"error": "check the types of the fields informed in the POST", "code": strconv.Itoa(http.StatusBadRequest)})
		return
	}

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		validationErrors := err.(validator.ValidationErrors)

		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{"error": validationErrors.Error(), "code": strconv.Itoa(http.StatusBadRequest)})
		return
	}

	err := isValid(user)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{"error": err.Error(), "code": strconv.Itoa(http.StatusBadRequest)})
		return
	}

	statusCode, user, err := u.userBusiness.PostUser(*user)

	if err != nil {
		c.AbortWithStatusJSON(statusCode, map[string]string{"error": err.Error(), "code": strconv.Itoa(statusCode)})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (u *UserController) GetUsers(c *gin.Context) {

	// -> verificar adicionar validação de paginação
	statusCode, user, err := u.userBusiness.GetUsers()

	if err != nil {
		c.AbortWithStatusJSON(statusCode, map[string]string{"error": err.Error(), "code": strconv.Itoa(statusCode)})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (u *UserController) UpdateUser(c *gin.Context) {

	user := &models.User{}
	if err := c.ShouldBindJSON(user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{"error": "check the types of the fields informed in the POST", "code": strconv.Itoa(http.StatusBadRequest)})
		return
	}

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		validationErrors := err.(validator.ValidationErrors)

		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{"error": validationErrors.Error(), "code": strconv.Itoa(http.StatusBadRequest)})
		return
	}

	if user.Id <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{"error": "ID canot be 0", "code": strconv.Itoa(http.StatusBadRequest)})
	}

	statusCode, user, err := u.userBusiness.UpdateUser(*user)

	if err != nil {
		c.AbortWithStatusJSON(statusCode, map[string]string{"error": err.Error(), "code": strconv.Itoa(statusCode)})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (u *UserController) DeleteUser(c *gin.Context) {

	id := c.Param("id")

	newId, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{"error": "ID has to be integer", "code": strconv.Itoa(http.StatusBadRequest)})
		return
	}

	if newId <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{"error": "ID canot be 0,", "code": strconv.Itoa(http.StatusBadRequest)})
		return
	}

	statusCode, err := u.userBusiness.DeleteUser(newId)

	if err != nil {
		c.AbortWithStatusJSON(statusCode, map[string]string{"error": err.Error(), "code": strconv.Itoa(statusCode)})
		return
	}

	c.JSON(http.StatusOK, "user deleted")
}

//isValid : valid required fields
func isValid(u *models.User) error {

	if u.Name == "" {
		return errors.New("name canot be empty")
	}

	if u.Email == "" {
		return errors.New("email canot be empty")
	}

	if u.CPFCNPJ == "" {
		return errors.New("document canot be empty")
	}

	return nil
}
