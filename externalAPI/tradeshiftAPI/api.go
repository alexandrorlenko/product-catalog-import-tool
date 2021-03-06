package tradeshiftAPI

import (
	"fmt"
	"go.uber.org/dig"
	"net/url"
	"ts/config"
	"ts/externalAPI/rest"
)

type TradeshiftAPI struct {
	TennantId string
	Client    rest.RestClientInterface
}

type Deps struct {
	dig.In
	Connection rest.RestClientInterface
	Config     *config.Config
}

func NewTradeshiftAPI(deps Deps) *TradeshiftAPI {
	return &TradeshiftAPI{
		Client: deps.Connection,
	}
}

func (t *TradeshiftAPI) GetIdentifier() (map[string]interface{}, error) {
	method := "/product-engine/supplier/supplier/v1/properties/identifier"
	resp, err := t.Client.Get(method)
	if err != nil {
		return nil, err
	}
	r, err := rest.ParseResponse(resp)
	return r, err
}

func (t *TradeshiftAPI) SetIdentifier(identifier string) error {
	method := "/product-engine/supplier/supplier/v1/properties/identifier"
	data := map[string]interface{}{
		"autoGenerated": false,
		"name":          identifier,
	}
	_, err := t.Client.Post(method, data, nil)
	return err
}

func (t *TradeshiftAPI) UploadFile(filePath string) (map[string]interface{}, error) {
	method := "/product-engine/supplier/supplier/v1/files"

	resp, err := t.Client.PostFile(method, filePath)
	r, err := rest.ParseResponse(resp)
	return r, err
}

func (t *TradeshiftAPI) RunImportAction(fileID string) (string, error) {
	method := fmt.Sprintf("/product-engine/supplier/supplier/v1/product-import/files/%v/actions/import-products", url.QueryEscape(fileID))
	resp, err := t.Client.Post(method,
		nil,
		[]rest.UrlAttributes{
			{
				Key:   "currency",
				Value: "USD",
			},
			{
				Key:   "fileLocale",
				Value: "en_US",
			},
		})
	if err != nil {
		return "", err
	}
	r, err := rest.ParseResponseToString(resp)
	return r, err
}

func (t *TradeshiftAPI) GetActionResult(actionID string) (map[string]interface{}, error) {
	method := fmt.Sprintf("/product-engine/supplier/supplier/v1/actions/%v", url.QueryEscape(actionID))
	resp, err := t.Client.Get(method)
	r, err := rest.ParseResponse(resp)
	return r, err
}

func (t *TradeshiftAPI) GetImportResult(actionID string) (string, error) {
	method := fmt.Sprintf("/product-engine/supplier/supplier/v1/actions/%v/reports/import-product-report/download", url.QueryEscape(actionID))
	resp, err := t.Client.Get(method)
	r, err := rest.ParseResponseToString(resp)
	return r, err
}
