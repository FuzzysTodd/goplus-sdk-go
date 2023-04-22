package nft

import (
	"encoding/json"
	"time"

	"k8s.io/utils/pointer"

	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/gen/client"
	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/gen/client/nft_controller"
)

type Config struct {
	// timeout seconds
	Timeout int
}

type NFTSecurity struct {
	AccessToken string
	Config      *Config
}

func NewNFTSecurity(accessToken string, config *Config) *NFTSecurity {
	return &NFTSecurity{
		AccessToken: accessToken,
		Config:      config,
	}
}

type Result struct {
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Result  map[string]any `json:"result"`
}

func (s *NFTSecurity) Run(chainId, address string) (*Result, error) {
	params := nft_controller.NewGetNftInfoUsingGET1Params()
	params.SetChainID(chainId)
	params.SetContractAddresses(address)
	if s.Config != nil && s.Config.Timeout != 0 {
		params.SetTimeout(time.Duration(s.Config.Timeout))
	}
	if s.AccessToken != "" {
		params.SetAuthorization(pointer.String(s.AccessToken))
	}

	response, err := client.Default.NftController.GetNftInfoUsingGET1(params)
	if err != nil {
		return nil, err
	}

	tmp, err := json.Marshal(response.Payload)
	if err != nil {
		return nil, err
	}

	res := Result{}
	err = json.Unmarshal(tmp, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
