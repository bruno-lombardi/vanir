package clients

import (
	"fmt"
	"vanir/internal/pkg/config"

	"github.com/imroc/req/v3"
)

type CryptoCompareHttpClient struct {
	*req.Client
}

func NewCryptoCompareHttpClient() *CryptoCompareHttpClient {
	conf := config.GetConfig()
	client := req.C().
		SetCommonHeader("Accept", "application/json").
		SetBaseURL(conf.HttpClients.CryptoCompareBaseURL).
		SetCommonHeader("Authorization", fmt.Sprintf("Apikey %s", conf.HttpClients.CryptoCompareAPIKey)).
		EnableDumpEachRequest()

	return &CryptoCompareHttpClient{
		Client: client,
	}
}
