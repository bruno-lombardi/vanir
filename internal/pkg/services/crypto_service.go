package services

import (
	"math"
	"strconv"
	"strings"
	"sync"
	"time"
	"vanir/internal/pkg/data/http/clients"
	"vanir/internal/pkg/data/models"
	"vanir/internal/pkg/data/repositories"
)

type CryptoService interface {
	ListTopCryptoCurrencies(params *models.ListTopCryptoCurrenciesQueryParams) (*models.ListCryptoCurrenciesResponse, error)
	ListUserFavoriteCryptoCurrencies(params *models.ListUserFavoriteCryptoCurrenciesQueryParams) (*models.ListCryptoCurrenciesResponse, error)
}

type CryptoServiceImpl struct {
	client              *clients.CryptoCompareHttpClient
	favoritesRepository *repositories.FavoritesRepository
}

var cryptoService *CryptoServiceImpl
var cryptoServiceOnce sync.Once

func GetCryptoService(client *clients.CryptoCompareHttpClient, favoritesRepository repositories.FavoritesRepository) CryptoService {
	cryptoServiceOnce.Do(func() {
		cryptoService = NewCryptoServiceImpl(client, favoritesRepository)
	})
	return cryptoService
}

func NewCryptoServiceImpl(client *clients.CryptoCompareHttpClient, favoritesRepository repositories.FavoritesRepository) *CryptoServiceImpl {
	return &CryptoServiceImpl{
		client:              client,
		favoritesRepository: &favoritesRepository,
	}
}

func (s *CryptoServiceImpl) ListTopCryptoCurrencies(params *models.ListTopCryptoCurrenciesQueryParams) (*models.ListCryptoCurrenciesResponse, error) {
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

	cryptoCurrencies := []models.CryptoCurrency{}
	page, err := strconv.Atoi(params.Page)
	if err != nil {
		return nil, err
	}

	limit, err := strconv.Atoi(params.Limit)
	if err != nil {
		return nil, err
	}
	count := response.MetaData.Count

	if len(response.Data) <= 0 {
		return &models.ListCryptoCurrenciesResponse{
			Data: cryptoCurrencies,
			Paginated: models.Paginated{
				TotalPages: 0,
				Count:      0,
				Page:       page,
				Limit:      limit,
			},
		}, nil
	}

	cryptoCurrencies = mapTopListResponse(response, params, cryptoCurrencies)

	return &models.ListCryptoCurrenciesResponse{
		Data: cryptoCurrencies,
		Paginated: models.Paginated{
			TotalPages: int(math.Ceil(float64(count) / float64(limit))),
			Count:      response.MetaData.Count,
			Page:       page,
			Limit:      limit,
		},
	}, nil
}

func (s *CryptoServiceImpl) ListUserFavoriteCryptoCurrencies(params *models.ListUserFavoriteCryptoCurrenciesQueryParams) (*models.ListCryptoCurrenciesResponse, error) {
	page, err := strconv.Atoi(params.Page)
	if err != nil {
		return nil, err
	}

	limit, err := strconv.Atoi(params.Limit)
	if err != nil {
		return nil, err
	}

	favorites, count, err := (*s.favoritesRepository).FindAllByUserID(params.UserID, page, limit)
	if err != nil {
		return nil, err
	}

	if len(favorites) <= 0 {
		return &models.ListCryptoCurrenciesResponse{
			Data: []models.CryptoCurrency{},

			Paginated: models.Paginated{
				Page:       page,
				Limit:      limit,
				Count:      int(count),
				TotalPages: 0,
			},
		}, nil
	}

	var favoriteReferences []string
	for _, favorite := range favorites {
		favoriteReferences = append(favoriteReferences, favorite.Reference)
	}
	userFavorites := strings.Join(favoriteReferences, ",")

	response := &clients.MultipleSymbolsResponse{}
	err = s.client.Get("data/pricemultifull").
		SetQueryParam("fsyms", userFavorites).
		SetQueryParam("tsyms", params.ToCurrency).
		Do(nil).
		Into(&response)
	if err != nil {
		return nil, err
	}

	favoriteCryptos := []models.CryptoCurrency{}

	for _, raw := range response.Raw {
		for _, details := range raw {
			crypto := models.CryptoCurrency{
				Name:     details.Fromsymbol,
				Code:     details.Fromsymbol,
				ImageUrl: details.Imageurl,
				Prices:   map[string]models.PriceDetails{},
			}
			crypto.Prices[params.ToCurrency] = mapRawPriceDetails(details, params.ToCurrency)
			favoriteCryptos = append(favoriteCryptos, crypto)
		}
	}

	return &models.ListCryptoCurrenciesResponse{
		Data: favoriteCryptos,

		Paginated: models.Paginated{
			Page:       page,
			Limit:      limit,
			Count:      int(count),
			TotalPages: int(math.Ceil(float64(count) / float64(limit))),
		},
	}, nil
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
		CurrencyCode:        currencyCode,
		Price:               rawPriceDetails.Price,
		OpenDay:             rawPriceDetails.Openday,
		HighDay:             rawPriceDetails.Highday,
		LowDay:              rawPriceDetails.Lowday,
		MarketCap:           rawPriceDetails.Mktcap,
		ChangePercentageDay: rawPriceDetails.Changepctday,
		LastUpdate:          time.Unix(int64(rawPriceDetails.Lastupdate), 0),
	}
}
