package user

import (
	"net/http"
	"strconv"

	"github.com/JoaoGabrielManochio/webapi-go/database"
	"github.com/JoaoGabrielManochio/webapi-go/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// IUserController : interface of user business
type IUserController interface {
	// PostUser(user model.UserClaims) (int, *model.BanklyDocumentResponse, error)
	CreateUser(c *gin.Context)
}

// UserController : struct of user Controller
type UserController struct {
	userBusiness IUserBusiness
}

// NewUserController : create a new user Controller
func NewUserController(userBusiness IUserBusiness) IUserController {
	return &UserController{userBusiness}
}

func ShowUser(c *gin.Context) {
	id := c.Param("id")

	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})

		return
	}

	db := database.GetDatabase()

	var user models.User

	err = db.First(&user, newid).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find user : " + err.Error(),
		})

		return
	}

	c.JSON(200, user)
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

	statusCode, user, err := u.userBusiness.PostUser(*user)

	if err != nil {
		c.AbortWithStatusJSON(statusCode, map[string]string{"error": err.Error(), "code": strconv.Itoa(statusCode)})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func ShowUsers(c *gin.Context) {

	// response, err := dbRepository.FindAll()

	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error": "cannot list user : " + err.Error(),
	// 	})

	// 	return
	// }

	// c.JSON(http.StatusOK, response)
}

func UpdateUser(c *gin.Context) {
	db := database.GetDatabase()

	var user models.User

	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})

		return
	}

	err = db.Save(&user).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot update user : " + err.Error(),
		})

		return
	}

	c.JSON(200, user)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})

		return
	}

	db := database.GetDatabase()

	err = db.Delete(&models.User{}, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot delete user: " + err.Error(),
		})

		return
	}

	c.Status(204)
}
