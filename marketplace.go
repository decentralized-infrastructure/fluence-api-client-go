package fluenceapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// GetBasicConfigurations fetches a list of basic configurations from the marketplace.
func (c *Client) GetBasicConfigurations() ([]string, error) {
	url := fmt.Sprintf("%s/marketplace/basic_configurations", c.HostURL)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var configs []string
	if err := json.Unmarshal(body, &configs); err != nil {
		return nil, err
	}
	return configs, nil
}

// GetAvailableCountries fetches a list of available countries on the marketplace.
func (c *Client) GetAvailableCountries() ([]string, error) {
	url := fmt.Sprintf("%s/marketplace/countries", c.HostURL)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var countries []string
	if err := json.Unmarshal(body, &countries); err != nil {
		return nil, err
	}
	return countries, nil
}

// GetAvailableHardware fetches a list of available hardware on the marketplace.
func (c *Client) GetAvailableHardware() (*AvailableHardware, error) {
	url := fmt.Sprintf("%s/marketplace/hardware", c.HostURL)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var hardware AvailableHardware
	if err := json.Unmarshal(body, &hardware); err != nil {
		return nil, err
	}
	return &hardware, nil
}

// GetMarketplaceOffers fetches a list of offers that match the given constraints.
func (c *Client) GetMarketplaceOffers(constraints *OfferConstraints) ([]MarketOffering, error) {
	url := fmt.Sprintf("%s/marketplace/offers", c.HostURL)
	payload := map[string]interface{}{}
	if constraints != nil {
		payload = map[string]interface{}{
			"additionalResources":      constraints.AdditionalResources,
			"basicConfiguration":       constraints.BasicConfiguration,
			"datacenter":               constraints.Datacenter,
			"hardware":                 constraints.Hardware,
			"maxTotalPricePerEpochUsd": constraints.MaxTotalPricePerEpochUsd,
		}
	}
	bodyBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var offers []MarketOffering
	if err := json.Unmarshal(body, &offers); err != nil {
		return nil, err
	}
	return offers, nil
}

// GetDatacenters fetches a list of registered datacenters.
func (c *Client) GetDatacenters() ([]DatacenterDTO, error) {
	url := fmt.Sprintf("%s/v1/marketplace/datacenters", c.HostURL)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var datacenters []DatacenterDTO
	if err := json.Unmarshal(body, &datacenters); err != nil {
		return nil, err
	}
	return datacenters, nil
}
