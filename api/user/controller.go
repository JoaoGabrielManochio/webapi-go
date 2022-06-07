package user

import (
	"net/http"
	"strconv"

	"github.com/JoaoGabrielManochio/webapi-go/database"
	"github.com/JoaoGabrielManochio/webapi-go/models"
	"github.com/gin-gonic/gin"
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
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{"error": "123", "code": "codigo do erro"})
		return
	}

	validate := validator.New()
	if err := validate.Struct(MyStruct); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		c.AbortWithStatusJSON(statusCode, map[string]string{"error": validationErrors, "code": "codigo do erro"})
		return

	u.userBusiness.PostUser(*user)

	// exemplo
	// if err := credential.IsValid(); err != nil {
	// 	c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{"error": "123", "code": "codigo do erro"})
	// 	return
	// }

	// access, statusCode, err := u.userBusiness.Login(credential)
	// if err != nil {
	// 	c.AbortWithStatusJSON(statusCode, map[string]string{"error": "123", "code": "codigo do erro"})
	// 	return
	// }

	// c.JSON(http.StatusOK, access)
	// exemplo

	// db := database.GetDatabase()

	// var user models.User

	// err := c.ShouldBindJSON(&user)

	// if err != nil {
	// 	c.JSON(400, gin.H{
	// 		"error": "cannot bind JSON: " + err.Error(),
	// 	})

	// 	return
	// }

	// err = db.Create(&user).Error

	// if err != nil {
	// 	c.JSON(400, gin.H{
	// 		"error": "cannot create user : " + err.Error(),
	// 	})

	// 	return
	// }

	c.JSON(201, user)
}

// func (c *Credential) IsValid() error {
// 	if c.DeviceInfo.DeviceCode == "" {
// 		return errors.New("device code é obrigatório")
// 	}
// 	if c.DeviceInfo.Brand == "" {
// 		return errors.New("brand é obrigatório")
// 	}
// 	if c.DeviceInfo.DeviceName == "" {
// 		return errors.New("device name é obrigatório")
// 	}
// 	if c.DeviceInfo.SystemName == "" {
// 		return errors.New("system name é obrigatório")
// 	}
// 	if c.DeviceInfo.SystemVersion == "" {
// 		return errors.New("system version é obrigatório")
// 	}
// 	return nil
// }

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
