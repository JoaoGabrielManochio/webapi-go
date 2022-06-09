package transaction_service

import (
	"encoding/json"
	"io"
	"net/http"
)

type ITransactionService interface {
	CallApi(url string) (FormatGetRequest, error)
}

// TransactionService : struct of Transaction serivce
type TransactionService struct{ http *http.Client }

type FormatGetRequest struct {
	Authorization bool `json:"authorization"`
}

// NewService : create a new Transaction service
func NewService(http *http.Client) ITransactionService {
	return &TransactionService{http}
}

// Get : make request by Get method
func (b *TransactionService) CallApi(url string) (FormatGetRequest, error) {

	var response FormatGetRequest

	resp, err := Get(url)

	if err != nil {
		return response, err
	}

	body, err := io.ReadAll(resp.Body)

	defer resp.Body.Close()

	if err != nil {
		return response, err
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return response, err
	}

	return response, err
}

func Get(url string) (*http.Response, error) {

	response, err := http.Get(url)

	return response, err

}
