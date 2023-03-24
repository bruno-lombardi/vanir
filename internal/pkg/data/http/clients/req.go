package clients

import "github.com/imroc/req/v3"

type ReqHttpClient struct {
	*req.Client
}

func NewReqHttpClient(baseUrl string) *ReqHttpClient {
	client := req.C().
		SetCommonHeader("Accept", "application/vnd+json").
		SetBaseURL(baseUrl).
		EnableDumpEachRequest()
	return &ReqHttpClient{
		Client: client,
	}
}
