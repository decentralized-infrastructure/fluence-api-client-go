package fluenceapi

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestGetDatacenters tests the new datacenter endpoint
func TestGetDatacenters(t *testing.T) {
	// Mock response data matching the expected DatacenterDTO structure
	mockResponse := []DatacenterDTO{
		{
			Id:             String("datacenter-1"),
			CountryCode:    "US",
			CityCode:       "NYC",
			Index:          1,
			Tier:           1,
			Certifications: []string{"SOC2", "ISO27001"},
			Slug:           "us-nyc-1",
		},
		{
			Id:             String("datacenter-2"),
			CountryCode:    "DE",
			CityCode:       "BER",
			Index:          1,
			Tier:           2,
			Certifications: []string{"ISO27001"},
			Slug:           "de-ber-1",
		},
	}

	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Validate the request
		if r.URL.Path != "/v1/marketplace/datacenters" {
			t.Errorf("Expected path /v1/marketplace/datacenters, got %s", r.URL.Path)
		}
		if r.Method != "GET" {
			t.Errorf("Expected GET method, got %s", r.Method)
		}

		// Check authorization header
		auth := r.Header.Get("Authorization")
		if auth == "" {
			t.Error("Expected Authorization header")
		}
		if auth != "Bearer test-api-key" {
			t.Errorf("Expected 'Bearer test-api-key', got '%s'", auth)
		}

		// Return mock response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(mockResponse)
	}))
	defer server.Close()

	// Create client with test server URL
	apiKey := "test-api-key"
	client, err := NewClient(&server.URL, &apiKey)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	// Test the GetDatacenters method
	datacenters, err := client.GetDatacenters()
	if err != nil {
		t.Fatalf("GetDatacenters failed: %v", err)
	}

	// Validate the response
	if len(datacenters) != 2 {
		t.Errorf("Expected 2 datacenters, got %d", len(datacenters))
	}

	// Validate first datacenter
	dc1 := datacenters[0]
	if dc1.CountryCode != "US" {
		t.Errorf("Expected CountryCode 'US', got '%s'", dc1.CountryCode)
	}
	if dc1.CityCode != "NYC" {
		t.Errorf("Expected CityCode 'NYC', got '%s'", dc1.CityCode)
	}
	if dc1.Slug != "us-nyc-1" {
		t.Errorf("Expected Slug 'us-nyc-1', got '%s'", dc1.Slug)
	}
	if len(dc1.Certifications) != 2 {
		t.Errorf("Expected 2 certifications, got %d", len(dc1.Certifications))
	}

	// Validate second datacenter
	dc2 := datacenters[1]
	if dc2.CountryCode != "DE" {
		t.Errorf("Expected CountryCode 'DE', got '%s'", dc2.CountryCode)
	}
	if dc2.Tier != 2 {
		t.Errorf("Expected Tier 2, got %d", dc2.Tier)
	}
}

// TestListVmsV3WithParams tests the enhanced VM listing with parameters
func TestListVmsV3WithParams(t *testing.T) {
	// Mock response data
	mockResponse := []RunningInstanceV3{
		{
			Id:              "vm-123",
			Status:          "Active",
			PricePerEpoch:   "1.23",
			Resources:       []VmResource{},
			CreatedAt:       "2024-01-01T00:00:00Z",
			NextBillingAt:   "2024-01-01T01:00:00Z",
			ReservedBalance: "10.00",
			TotalSpent:      "5.00",
			StatusChangedAt: "2024-01-01T00:00:00Z",
			TerminatedAt:    "2024-01-01T02:00:00Z",
		},
	}

	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Validate the request path
		if r.URL.Path != "/vms/v3" {
			t.Errorf("Expected path /vms/v3, got %s", r.URL.Path)
		}
		if r.Method != "GET" {
			t.Errorf("Expected GET method, got %s", r.Method)
		}

		// Check for query parameters
		query := r.URL.Query()
		if pagination := query.Get("pagination[page]"); pagination != "1" {
			t.Errorf("Expected pagination[page]=1, got %s", pagination)
		}
		if perPage := query.Get("pagination[per_page]"); perPage != "5" {
			t.Errorf("Expected pagination[per_page]=5, got %s", perPage)
		}
		if dcSlug := query.Get("dcSlug"); dcSlug != "us-nyc-1" {
			t.Errorf("Expected dcSlug=us-nyc-1, got %s", dcSlug)
		}

		// Return mock response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(mockResponse)
	}))
	defer server.Close()

	// Create client with test server URL
	apiKey := "test-api-key"
	client, err := NewClient(&server.URL, &apiKey)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	// Create test parameters
	page := int64(1)
	perPage := int64(5)
	dcSlug := "us-nyc-1"
	params := &ListVmsV3Params{
		DcSlug: &dcSlug,
		Pagination: &PaginationParams{
			Page:    &page,
			PerPage: &perPage,
		},
	}

	// Test the ListVmsV3WithParams method
	vms, err := client.ListVmsV3WithParams(params)
	if err != nil {
		t.Fatalf("ListVmsV3WithParams failed: %v", err)
	}

	// Validate the response
	if len(vms) != 1 {
		t.Errorf("Expected 1 VM, got %d", len(vms))
	}

	vm := vms[0]
	if vm.Id != "vm-123" {
		t.Errorf("Expected VM ID 'vm-123', got '%s'", vm.Id)
	}
	if vm.Status != "Active" {
		t.Errorf("Expected Status 'Active', got '%s'", vm.Status)
	}
	if vm.TerminatedAt != "2024-01-01T02:00:00Z" {
		t.Errorf("Expected TerminatedAt '2024-01-01T02:00:00Z', got '%s'", vm.TerminatedAt)
	}
}

// TestDatacenterDTOFields tests that all required fields are present in DatacenterDTO
func TestDatacenterDTOFields(t *testing.T) {
	// Test JSON marshaling/unmarshaling
	jsonData := `{
		"id": "test-id",
		"countryCode": "US",
		"cityCode": "NYC",
		"index": 1,
		"tier": 2,
		"certifications": ["SOC2", "ISO27001"],
		"slug": "us-nyc-1"
	}`

	var dc DatacenterDTO
	err := json.Unmarshal([]byte(jsonData), &dc)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	// Validate all fields
	if string(dc.Id) != "test-id" {
		t.Errorf("Expected Id 'test-id', got '%s'", dc.Id)
	}
	if dc.CountryCode != "US" {
		t.Errorf("Expected CountryCode 'US', got '%s'", dc.CountryCode)
	}
	if dc.CityCode != "NYC" {
		t.Errorf("Expected CityCode 'NYC', got '%s'", dc.CityCode)
	}
	if dc.Index != 1 {
		t.Errorf("Expected Index 1, got %d", dc.Index)
	}
	if dc.Tier != 2 {
		t.Errorf("Expected Tier 2, got %d", dc.Tier)
	}
	if len(dc.Certifications) != 2 {
		t.Errorf("Expected 2 certifications, got %d", len(dc.Certifications))
	}
	if dc.Slug != "us-nyc-1" {
		t.Errorf("Expected Slug 'us-nyc-1', got '%s'", dc.Slug)
	}

	// Test marshaling back to JSON
	_, err = json.Marshal(dc)
	if err != nil {
		t.Fatalf("Failed to marshal to JSON: %v", err)
	}
}

// TestVmStatusEnums tests that all VM status constants are correctly defined
func TestVmStatusEnums(t *testing.T) {
	// Test all VM status constants
	statuses := []string{
		VmStatusNew,
		VmStatusLaunching,
		VmStatusActive,
		VmStatusSmallBalance,
		VmStatusInsufficientFunds,
		VmStatusTerminated,
		VmStatusStopped,
		VmStatusFailed,
	}

	expectedStatuses := []string{
		"New",
		"Launching",
		"Active",
		"SmallBalance",
		"InsufficientFunds",
		"Terminated",
		"Stopped",
		"Failed",
	}

	if len(statuses) != len(expectedStatuses) {
		t.Errorf("Expected %d status constants, got %d", len(expectedStatuses), len(statuses))
	}

	for i, status := range statuses {
		if status != expectedStatuses[i] {
			t.Errorf("Expected status constant %d to be '%s', got '%s'", i, expectedStatuses[i], status)
		}
	}
}
