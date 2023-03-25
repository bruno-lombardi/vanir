package models

import "time"

type CryptoCurrency struct {
	Name     string `json:"name"`
	FullName string `json:"fullname"`
	Code     string `json:"code"`
	ImageUrl string `json:"image_url"`

	Prices map[string]PriceDetails `json:"prices"`
}

type ListTopCryptoCurrenciesQueryParams struct {
	Page       string `query:"page" validate:"required"`
	Limit      string `query:"limit" validate:"required"`
	ToCurrency string `query:"to_currency" validate:"required"`
}

type CryptoCurrencyPriceHistory struct {
	Name     string             `json:"name"`
	Code     string             `json:"code"`
	TimeFrom *time.Time         `json:"time_from"`
	TimeTo   *time.Time         `json:"time_to"`
	History  map[string][]OHLCV `json:"history"`
}

type OHLCV struct {
	Time         *time.Time `json:"time"`
	Open         float64    `json:"open"`
	High         float64    `json:"high"`
	Low          float64    `json:"low"`
	Close        float64    `json:"close"`
	CurrencyCode string     `json:"currency_code"`
}

type PriceDetails struct {
	CurrencyCode           string    `json:"currency_code"`
	Price                  float64   `json:"price"`
	Open24Hour             float64   `json:"open_24_hour"`
	High24Hour             float64   `json:"high_24_hour"`
	Low24Hour              float64   `json:"low_24_hour"`
	OpenDay                float64   `json:"open_day"`
	HighDay                float64   `json:"high_day"`
	ChangePercentage24Hour float64   `json:"change_percentage_24_hour`
	LowDay                 float64   `json:"low_day"`
	MarketCap              float64   `json:"market_cap"`
	LastUpdate             time.Time `json:"last_update"`
}
