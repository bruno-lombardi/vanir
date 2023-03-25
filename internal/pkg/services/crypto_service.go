package services

import (
	"sync"
	"time"
	"vanir/internal/pkg/data/http/clients"
	"vanir/internal/pkg/data/models"
)

type CryptoService interface {
	ListTopCryptoCurrencies(params *models.ListTopCryptoCurrenciesQueryParams) ([]models.CryptoCurrency, error)
}

type CryptoServiceImpl struct {
	client *clients.CryptoCompareHttpClient
}

var cryptoService *CryptoServiceImpl
var cryptoServiceOnce sync.Once

func GetCryptoService() CryptoService {
	cryptoServiceOnce.Do(func() {
		cryptoService = NewCryptoServiceImpl(clients.NewCryptoCompareHttpClient())
	})
	return cryptoService
}

func NewCryptoServiceImpl(client *clients.CryptoCompareHttpClient) *CryptoServiceImpl {
	return &CryptoServiceImpl{
		client: client,
	}
}

func (s *CryptoServiceImpl) ListTopCryptoCurrencies(params *models.ListTopCryptoCurrenciesQueryParams) ([]models.CryptoCurrency, error) {
	response := &clients.TopListResponse{}
	err := s.client.Get("data/top/totaltoptiervolfull").
		SetQueryParam("limit", params.Limit).
		SetQueryParam("page", params.Page).
		SetQueryParam("tsym", params.ToCurrency).
		SetQueryParam("ascending", "true").
		Do(nil).
		Into(&response)

	if err != nil {
		return nil, err
	}

	topCryptos := []models.CryptoCurrency{}

	if len(response.Data) <= 0 {
		return topCryptos, nil
	}

	topCryptos = mapTopListResponse(response, params, topCryptos)

	return topCryptos, nil
}

func mapTopListResponse(response *clients.TopListResponse, params *models.ListTopCryptoCurrenciesQueryParams, topCryptos []models.CryptoCurrency) []models.CryptoCurrency {
	for _, item := range response.Data {
		crypto := models.CryptoCurrency{
			Name:     item.CoinInfo.Name,
			FullName: item.CoinInfo.FullName,
			Code:     item.CoinInfo.Name,
			ImageUrl: item.CoinInfo.ImageURL,
			Prices:   map[string]models.PriceDetails{},
		}
		rawPriceDetails := item.Raw[params.ToCurrency]
		crypto.Prices[params.ToCurrency] = mapRawPriceDetails(rawPriceDetails, params.ToCurrency)
		topCryptos = append(topCryptos, crypto)
	}
	return topCryptos
}

func mapRawPriceDetails(rawPriceDetails clients.PriceDetailsRaw, currencyCode string) models.PriceDetails {
	return models.PriceDetails{
		CurrencyCode:           currencyCode,
		Price:                  rawPriceDetails.Price,
		Open24Hour:             rawPriceDetails.Open24Hour,
		High24Hour:             rawPriceDetails.High24Hour,
		Low24Hour:              rawPriceDetails.Low24Hour,
		OpenDay:                rawPriceDetails.Openday,
		HighDay:                rawPriceDetails.Highday,
		LowDay:                 rawPriceDetails.Lowday,
		MarketCap:              rawPriceDetails.Mktcap,
		ChangePercentage24Hour: rawPriceDetails.Changepct24Hour,
		LastUpdate:             time.Unix(int64(rawPriceDetails.Lastupdate), 0),
	}
}
