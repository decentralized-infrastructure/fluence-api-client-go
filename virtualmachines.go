package fluenceapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// ListVmsV3Params represents the query parameters for ListVmsV3
type ListVmsV3Params struct {
	Status       []VirtualMachineStatus `json:"status,omitempty"`
	DcSlug       *string                `json:"dcSlug,omitempty"`
	Pagination   *PaginationParams      `json:"pagination,omitempty"`
	DealStatus   []ExtendedDealStatus   `json:"dealStatus,omitempty"`
	TerminatedAt *time.Time             `json:"terminatedAt,omitempty"`
}

// ListVmsV3 fetches all running VM instances (v3).
func (c *Client) ListVmsV3() ([]RunningInstanceV3, error) {
	return c.ListVmsV3WithParams(nil)
}

// ListVmsV3WithParams fetches VM instances with query parameters (v3).
func (c *Client) ListVmsV3WithParams(params *ListVmsV3Params) ([]RunningInstanceV3, error) {
	baseURL := fmt.Sprintf("%s/vms/v3", c.HostURL)

	if params != nil {
		queryParams := url.Values{}

		if len(params.Status) > 0 {
			statusStrings := make([]string, len(params.Status))
			for i, status := range params.Status {
				statusStrings[i] = string(status)
			}
			queryParams.Add("status", strings.Join(statusStrings, ","))
		}

		if params.DcSlug != nil {
			queryParams.Add("dcSlug", *params.DcSlug)
		}

		if params.Pagination != nil {
			if params.Pagination.Page != nil {
				queryParams.Add("pagination[page]", strconv.FormatInt(*params.Pagination.Page, 10))
			}
			if params.Pagination.PerPage != nil {
				queryParams.Add("pagination[per_page]", strconv.FormatInt(*params.Pagination.PerPage, 10))
			}
		}

		if len(params.DealStatus) > 0 {
			dealStatusStrings := make([]string, len(params.DealStatus))
			for i, status := range params.DealStatus {
				dealStatusStrings[i] = string(status)
			}
			queryParams.Add("dealStatus", strings.Join(dealStatusStrings, ","))
		}

		if params.TerminatedAt != nil {
			queryParams.Add("terminatedAt", params.TerminatedAt.Format(time.RFC3339))
		}

		if len(queryParams) > 0 {
			baseURL += "?" + queryParams.Encode()
		}
	}

	req, err := http.NewRequest("GET", baseURL, nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var vms []RunningInstanceV3
	if err := json.Unmarshal(body, &vms); err != nil {
		return nil, err
	}
	return vms, nil
}

// CreateVmV3 creates new VM(s) with the given configuration and constraints.
func (c *Client) CreateVmV3(reqBody CreateVmV3) ([]CreatedVm, error) {
	url := fmt.Sprintf("%s/vms/v3", c.HostURL)
	payload, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var created []CreatedVm
	if err := json.Unmarshal(body, &created); err != nil {
		return nil, err
	}
	return created, nil
}

// RemoveVms removes VMs by their IDs.
func (c *Client) RemoveVms(vmIds []string) (*VmsRemoved, error) {
	url := fmt.Sprintf("%s/vms/v3", c.HostURL)
	payload, err := json.Marshal(RemoveVms{VmIds: vmIds})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("DELETE", url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var removed VmsRemoved
	if err := json.Unmarshal(body, &removed); err != nil {
		return nil, err
	}
	return &removed, nil
}

// UpdateVms updates one or more VMs (patch).
func (c *Client) UpdateVms(updates []UpdateVm) error {
	url := fmt.Sprintf("%s/vms/v3", c.HostURL)
	payload, err := json.Marshal(UpdateVms{Updates: updates})
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	_, err = c.doRequest(req)
	return err
}

// EstimateDeposit estimates the deposit required for a VM configuration.
func (c *Client) EstimateDeposit(reqBody EstimateDepositRequestV3) (*EstimatedDepositV3DTO, error) {
	url := fmt.Sprintf("%s/vms/v3/estimate", c.HostURL)
	payload, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var estimate EstimatedDepositV3DTO
	if err := json.Unmarshal(body, &estimate); err != nil {
		return nil, err
	}
	return &estimate, nil
}

// GetDefaultImages fetches the list of default OS images.
func (c *Client) GetDefaultImages() ([]DefaultImageDTO, error) {
	url := fmt.Sprintf("%s/vms/v3/default_images", c.HostURL)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var images []DefaultImageDTO
	if err := json.Unmarshal(body, &images); err != nil {
		return nil, err
	}
	return images, nil
}

// GetVmStatuses fetches VM statuses for the specified VM IDs.
func (c *Client) GetVmStatuses(ids string) ([]VmStatusInfoDTO, error) {
	url := fmt.Sprintf("%s/vms/v3/status?ids=%s", c.HostURL, ids)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var statuses []VmStatusInfoDTO
	if err := json.Unmarshal(body, &statuses); err != nil {
		return nil, err
	}
	return statuses, nil
}
