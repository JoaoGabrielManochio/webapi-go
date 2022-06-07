package user

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/JoaoGabrielManochio/webapi-go/models"
	"github.com/JoaoGabrielManochio/webapi-go/repository/user"
	"gorm.io/gorm"
)

// IUserBusiness : interface of user business
type IUserBusiness interface {
	PostUser(user models.User) (int, *models.User, error)
	// GetUsers(userId string) (int, *model.DocumentStatusInfo, error)
}

// UserBusiness : struct of user business
type UserBusiness struct {
	UserRepository user.IUserRepository
}

// NewUserBusiness : create a new user business
func NewUserBusiness(UserRepository user.IUserRepository) IUserBusiness {
	return &UserBusiness{UserRepository}
}

// PostUser : post user
func (a *UserBusiness) PostUser(user models.User) (int, *models.User, error) {

	document, err := a.UserRepository.GetDocument(&model.UserDocument{
		UserId:             user.UserId,
		UserDocumentTypeId: *documentTypeId,
		UserDocumentSideId: documentSideId,
		IsActive:           true,
	})
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return http.StatusInternalServerError, nil, err
	}

	if document != nil && document.Id != 0 {
		documentStatus := strings.ToLower(document.UserDocumentStatus.Name)
		if documentStatus == Constants.Q2Bank.KYCStatus.InProgress || documentStatus == Constants.Q2Bank.KYCStatus.Approved {
			return http.StatusBadRequest, nil, errors.New("tipo do documento já está vinculado")
		}

		if err := a.UserRepository.UpdateStatus(document.Id, document.UserDocumentStatusId, false); err != nil {
			return http.StatusInternalServerError, nil, err
		}
	}

	merchantDetail, statusCode, err := a.ChrisService.GetMerchant(user.MerchantId)
	if err != nil {
		return statusCode, nil, err
	}

	documentNumber := merchantDetail.LegalRepresentative.DocumentNumber
	statusCode, banklyDocument, err := a.BanklyDocumentService.Post(documentNumber, documentType, documentSide, fileName, contentType, file)
	if err != nil {
		return statusCode, nil, err
	}

	err = a.UserRepository.Create(&model.UserDocument{
		UserId:               user.UserId,
		UserDocumentTypeId:   *documentTypeId,
		UserDocumentSideId:   documentSideId,
		UserDocumentStatusId: Constants.Q2Bank.UserDocument.Status.InProgress,
		Token:                banklyDocument.Token,
		IsActive:             true,
		CreatedAt:            time.Now(),
	})

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusAccepted, &banklyDocument, nil
}

/*

// GetDocumentStatus : call bankly service and get document status with document type
func (a *UserBusiness) GetDocumentStatus(userId string) (int, *model.DocumentStatusInfo, error) {
	documents, err := a.UserRepository.Get(&model.UserDocument{
		UserId:   userId,
		IsActive: true,
	})

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	documentsStatus := getDocumentsStatus(documents)
	return http.StatusOK, documentsStatus, nil
}
*/
