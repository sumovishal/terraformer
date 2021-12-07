package sumologic

import (
	"errors"
	"os"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type SumologicProvider struct { //nolint
	terraformutils.Provider
	AccessKey string
	AccessID  string
	BaseUrl   string
}

func (p *SumologicProvider) Init(args []string) error {

	accessKey := os.Getenv("SUMOLOGIC_ACCESSKEY")
	if accessKey == "" {
		return errors.New("set SUMOLOGIC_ACCESSKEY env var")
	}
	p.AccessKey = accessKey

	accessID := os.Getenv("SUMOLOGIC_ACCESSID")
	if accessID == "" {
		return errors.New("set SUMOLOGIC_ACCESSID env var")
	}
	p.AccessID = accessID

	baseUrl := os.Getenv("SUMOLOGIC_BASE_URL")
	if baseUrl == "" {
		return errors.New("set SUMOLOGIC_BASE_URL env var")
	}
	p.BaseUrl = baseUrl

	return nil
}

// GetName return string of provider name for Sumologic
func (p *SumologicProvider) GetName() string {
	return "sumologic"
}

// GetProviderData return map of provider data for Sumologic
func (p *SumologicProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{}
}

// GetResourceConnections return map of resource connections for Sumologic
func (SumologicProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{}
}

// GetSupportedService return map of support service for Sumologic
func (p *SumologicProvider) GetSupportedService() map[string]terraformutils.ServiceGenerator {
	return map[string]terraformutils.ServiceGenerator{
		"user":    &UserGenerator{},
		"monitor": &MonitorGenerator{},
	}
}

// InitService ...
func (p *SumologicProvider) InitService(serviceName string, verbose bool) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New("Sumologic: " + serviceName + " not supported service")
	}
	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetProviderName(p.GetName())
	p.Service.SetVerbose(verbose)
	p.Service.SetArgs(map[string]interface{}{
		"AccessKey": p.AccessKey,
		"AccessID":  p.AccessID,
		"BaseUrl":   p.BaseUrl,
	})

	return nil
}
