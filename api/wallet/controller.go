package wallet

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/JoaoGabrielManochio/webapi-go/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// IWalletController : interface of Wallet business
type IWalletController interface {
	CreateWallet(c *gin.Context)
	GetWallet(c *gin.Context)
	GetWallets(c *gin.Context)
	UpdateWallet(c *gin.Context)
	DeleteWallet(c *gin.Context)
}

// WalletController : struct of Wallet Controller
type WalletController struct {
	walletBusiness IWalletBusiness
}

// NewWalletController : create a new Wallet Controller
func NewWalletController(walletBusiness IWalletBusiness) IWalletController {
	return &WalletController{walletBusiness}
}

// GetWallet godoc
// @Summary Get a wallet
// @id getWallet
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} models.Wallet
// @Router /wallet/{{id}} [GET]
func (u *WalletController) GetWallet(c *gin.Context) {

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

	statusCode, wallet, err := u.walletBusiness.GetWallet(newId)

	if err != nil {
		c.AbortWithStatusJSON(statusCode, map[string]string{"error": err.Error(), "code": strconv.Itoa(statusCode)})
		return
	}

	c.JSON(http.StatusOK, wallet)
}

// CreateWallet godoc
// @Summary Create a new wallet
// @id createWallet
// @Produce  json
// @Param name path string true "name"
// @Param user_id path int true "user_id"
// @Param value path number true "value"
// @Success 200 {object} models.Wallet
// @Router /wallet [POST]
func (u *WalletController) CreateWallet(c *gin.Context) {

	wallet := &models.Wallet{}
	if err := c.ShouldBindJSON(wallet); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{"error": "check the types of the fields informed in the POST", "code": strconv.Itoa(http.StatusBadRequest)})
		return
	}

	validate := validator.New()
	if err := validate.Struct(wallet); err != nil {
		validationErrors := err.(validator.ValidationErrors)

		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{"error": validationErrors.Error(), "code": strconv.Itoa(http.StatusBadRequest)})
		return
	}

	err := isValid(wallet)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{"error": err.Error(), "code": strconv.Itoa(http.StatusBadRequest)})
		return
	}

	statusCode, wallet, err := u.walletBusiness.PostWallet(*wallet)

	if err != nil {
		c.AbortWithStatusJSON(statusCode, map[string]string{"error": err.Error(), "code": strconv.Itoa(statusCode)})
		return
	}

	c.JSON(http.StatusCreated, wallet)
}

// GetWallets godoc
// @Summary Get all wallets
// @id getWallets
// @Produce  json
// @Success 200 {object} models.Wallet
// @Router /wallet [GET]
func (u *WalletController) GetWallets(c *gin.Context) {

	statusCode, wallet, err := u.walletBusiness.GetWallets()

	if err != nil {
		c.AbortWithStatusJSON(statusCode, map[string]string{"error": err.Error(), "code": strconv.Itoa(statusCode)})
		return
	}

	c.JSON(http.StatusOK, wallet)
}

// UpdateWallet godoc
// @Summary Update a wallet
// @id updateWallet
// @Produce  json
// @Param id path int true "id"
// @Param name path string true "name"
// @Param user_id path int true "user_id"
// @Param value path number true "value"
// @Success 200 {object} models.Wallet
// @Router /wallet [PUT]
func (u *WalletController) UpdateWallet(c *gin.Context) {

	wallet := &models.Wallet{}
	if err := c.ShouldBindJSON(wallet); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{"error": "check the types of the fields informed in the POST", "code": strconv.Itoa(http.StatusBadRequest)})
		return
	}

	validate := validator.New()
	if err := validate.Struct(wallet); err != nil {
		validationErrors := err.(validator.ValidationErrors)

		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{"error": validationErrors.Error(), "code": strconv.Itoa(http.StatusBadRequest)})
		return
	}

	if wallet.Id <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{"error": "ID canot be 0 or less then 0", "code": strconv.Itoa(http.StatusBadRequest)})
	}

	err := isValid(wallet)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{"error": err.Error(), "code": strconv.Itoa(http.StatusBadRequest)})
		return
	}

	statusCode, wallet, err := u.walletBusiness.UpdateWallet(*wallet)

	if err != nil {
		c.AbortWithStatusJSON(statusCode, map[string]string{"error": err.Error(), "code": strconv.Itoa(statusCode)})
		return
	}

	c.JSON(http.StatusCreated, wallet)
}

// DeleteWallet godoc
// @Summary Delete a wallet
// @id deleteWallet
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} models.Wallet
// @Router /wallet [DELETE]
func (u *WalletController) DeleteWallet(c *gin.Context) {

	id := c.Param("id")

	newId, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{"error": "ID has to be integer", "code": strconv.Itoa(http.StatusBadRequest)})
		return
	}

	if newId <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{"error": "ID canot be 0 or less then 0,", "code": strconv.Itoa(http.StatusBadRequest)})
		return
	}

	statusCode, err := u.walletBusiness.DeleteWallet(newId)

	if err != nil {
		c.AbortWithStatusJSON(statusCode, map[string]string{"error": err.Error(), "code": strconv.Itoa(statusCode)})
		return
	}

	c.JSON(http.StatusOK, "wallet deleted")
}

//isValid : valid required fields
func isValid(w *models.Wallet) error {

	if w.Name == "" {
		return errors.New("name canot be empty")
	}

	if w.UserId <= 0 {
		return errors.New("user_id canot be empty or less than zero")
	}

	if w.Value < 0 {
		return errors.New("value canot be less than zero")
	}

	return nil
}
