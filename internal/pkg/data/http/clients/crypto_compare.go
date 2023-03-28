package clients

import (
	"fmt"
	"net/http"
	"vanir/internal/pkg/config"
	"vanir/internal/pkg/protocols"

	"github.com/imroc/req/v3"
)

type CryptoCompareHttpClient struct {
	*req.Client
}

type TopListResponse struct {
	Data     []TopListCurrencyItem   `json:"Data"`
	MetaData TopListResponseMetadata `json:"MetaData"`
}

type MultipleSymbolsResponse struct {
	Raw map[string]map[string]PriceDetailsRaw `json:"Raw"`
}

type TopListResponseMetadata struct {
	Count int `json:"Count"`
}

type PriceDetailsRaw struct {
	Type                    string  `json:"TYPE"`
	Market                  string  `json:"MARKET"`
	Fromsymbol              string  `json:"FROMSYMBOL"`
	Tosymbol                string  `json:"TOSYMBOL"`
	Flags                   string  `json:"FLAGS"`
	Price                   float64 `json:"PRICE"`
	Lastupdate              int     `json:"LASTUPDATE"`
	Median                  float64 `json:"MEDIAN"`
	Lastvolume              float64 `json:"LASTVOLUME"`
	Lastvolumeto            float64 `json:"LASTVOLUMETO"`
	Lasttradeid             string  `json:"LASTTRADEID"`
	Volumeday               float64 `json:"VOLUMEDAY"`
	Volumedayto             float64 `json:"VOLUMEDAYTO"`
	Volume24Hour            float64 `json:"VOLUME24HOUR"`
	Volume24Hourto          float64 `json:"VOLUME24HOURTO"`
	Openday                 float64 `json:"OPENDAY"`
	Highday                 float64 `json:"HIGHDAY"`
	Lowday                  float64 `json:"LOWDAY"`
	Open24Hour              float64 `json:"OPEN24HOUR"`
	High24Hour              float64 `json:"HIGH24HOUR"`
	Low24Hour               float64 `json:"LOW24HOUR"`
	Lastmarket              string  `json:"LASTMARKET"`
	Volumehour              float64 `json:"VOLUMEHOUR"`
	Volumehourto            float64 `json:"VOLUMEHOURTO"`
	Openhour                float64 `json:"OPENHOUR"`
	Highhour                float64 `json:"HIGHHOUR"`
	Lowhour                 float64 `json:"LOWHOUR"`
	Toptiervolume24Hour     float64 `json:"TOPTIERVOLUME24HOUR"`
	Toptiervolume24Hourto   float64 `json:"TOPTIERVOLUME24HOURTO"`
	Change24Hour            float64 `json:"CHANGE24HOUR"`
	Changepct24Hour         float64 `json:"CHANGEPCT24HOUR"`
	Changeday               float64 `json:"CHANGEDAY"`
	Changepctday            float64 `json:"CHANGEPCTDAY"`
	Changehour              float64 `json:"CHANGEHOUR"`
	Changepcthour           float64 `json:"CHANGEPCTHOUR"`
	Conversiontype          string  `json:"CONVERSIONTYPE"`
	Conversionsymbol        string  `json:"CONVERSIONSYMBOL"`
	Supply                  float64 `json:"SUPPLY"`
	Mktcap                  float64 `json:"MKTCAP"`
	Mktcappenalty           int     `json:"MKTCAPPENALTY"`
	Circulatingsupply       float64 `json:"CIRCULATINGSUPPLY"`
	Circulatingsupplymktcap float64 `json:"CIRCULATINGSUPPLYMKTCAP"`
	Totalvolume24H          float64 `json:"TOTALVOLUME24H"`
	Totalvolume24Hto        float64 `json:"TOTALVOLUME24HTO"`
	Totaltoptiervolume24H   float64 `json:"TOTALTOPTIERVOLUME24H"`
	Totaltoptiervolume24Hto float64 `json:"TOTALTOPTIERVOLUME24HTO"`
	Imageurl                string  `json:"IMAGEURL"`
}

type TopListCurrencyItem struct {
	CoinInfo struct {
		Name         string `json:"Name"`
		FullName     string `json:"FullName"`
		ImageURL     string `json:"ImageUrl"`
		URL          string `json:"Url"`
		Algorithm    string `json:"Algorithm"`
		ProofType    string `json:"ProofType"`
		Type         int    `json:"Type"`
		DocumentType string `json:"DocumentType"`
	} `json:"CoinInfo"`
	Raw map[string]PriceDetailsRaw `json:"RAW"`
}

func NewCryptoCompareHttpClient() *CryptoCompareHttpClient {
	conf := config.GetConfig()
	client := req.C().
		SetCommonHeader("Accept", "application/json").
		SetBaseURL(conf.HttpClients.CryptoCompareBaseURL).
		SetCommonHeader("Authorization", fmt.Sprintf("Apikey %s", conf.HttpClients.CryptoCompareAPIKey)).
		EnableDumpEachRequest().
		SetCommonErrorResult(&protocols.CryptoCompareError{}).
		OnAfterResponse(func(client *req.Client, resp *req.Response) error {
			if resp.Err != nil {
				return nil
			}
			if apiErr, ok := resp.ErrorResult().(*protocols.CryptoCompareError); ok {
				resp.Err = apiErr
				return nil
			}
			if !resp.IsSuccessState() {
				return fmt.Errorf("bad response, raw dump:\n%s", resp.Dump())
			}
			// HTTP status is 2xx, but crypto compare returns error
			if resp.StatusCode == http.StatusOK {
				err := &protocols.CryptoCompareError{}
				resp.Unmarshal(err)
				if err.Response == "Error" {
					return err
				}
			}
			return nil
		})

	return &CryptoCompareHttpClient{
		Client: client,
	}
}
