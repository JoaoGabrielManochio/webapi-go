package transaction

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/JoaoGabrielManochio/webapi-go/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// ITransactionController : interface of Transaction business
type ITransactionController interface {
	CreateTransaction(c *gin.Context)
	GetTransaction(c *gin.Context)
	GetTransactions(c *gin.Context)
}

// TransactionController : struct of Transaction Controller
type TransactionController struct {
	transactionBusiness ITransactionBusiness
}

// NewTransactionController : create a new Transaction Controller
func NewTransactionController(transactionBusiness ITransactionBusiness) ITransactionController {
	return &TransactionController{transactionBusiness}
}

// GetTransaction godoc
// @Summary Show a transaction
// @ID get-int
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} Transaction
// @Router /transaction/{{id}} [GET]
func (u *TransactionController) GetTransaction(c *gin.Context) {

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

	statusCode, transaction, err := u.transactionBusiness.GetTransaction(newId)

	if err != nil {
		c.AbortWithStatusJSON(statusCode, map[string]string{"error": err.Error(), "code": strconv.Itoa(statusCode)})
		return
	}

	c.JSON(http.StatusOK, transaction)
}

// GetTransactions godoc
// @Summary Show all transaction
// @Produce  json
// @Success 200 {object} Transaction
// @Router /transaction [GET]
func (u *TransactionController) GetTransactions(c *gin.Context) {

	statusCode, transaction, err := u.transactionBusiness.GetTransactions()

	if err != nil {
		c.AbortWithStatusJSON(statusCode, map[string]string{"error": err.Error(), "code": strconv.Itoa(statusCode)})
		return
	}

	c.JSON(http.StatusOK, transaction)
}

// CreateTransaction godoc
// @Summary Create a new transaction
// @value get-float
// @payer_id get-int
// @payer_receive_id get-int
// @Produce  json
// @Param value path int true "value"
// @Param payer_id path int true "payer_id"
// @Param payer_receive_id path int true "payer_receive_id"
// @Success 200 {object} Transaction
// @Router /transaction [POST]
func (u *TransactionController) CreateTransaction(c *gin.Context) {

	transaction := &models.Transaction{}
	if err := c.ShouldBindJSON(transaction); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{"error": "check the types of the fields informed in the POST", "code": strconv.Itoa(http.StatusBadRequest)})
		return
	}

	validate := validator.New()
	if err := validate.Struct(transaction); err != nil {
		validationErrors := err.(validator.ValidationErrors)

		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{"error": validationErrors.Error(), "code": strconv.Itoa(http.StatusBadRequest)})
		return
	}

	err := isValid(transaction)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{"error": err.Error(), "code": strconv.Itoa(http.StatusBadRequest)})
		return
	}

	statusCode, transaction, err := u.transactionBusiness.PostTransaction(*transaction)

	if err != nil {
		c.AbortWithStatusJSON(statusCode, map[string]string{"error": err.Error(), "code": strconv.Itoa(statusCode)})
		return
	}

	c.JSON(http.StatusCreated, transaction)
}

//isValid : valid required fields
func isValid(t *models.Transaction) error {

	if t.Value <= 0 {
		return errors.New("value canot be 0 or less then 0")
	}

	if t.Payer_id <= 0 {
		return errors.New("payer_id canot be empty or less than zero")
	}

	if t.Payer_receive_id <= 0 {
		return errors.New("payer_recive_id canot be 0 or less than zero")
	}

	return nil
}
