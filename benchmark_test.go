package fluenceapi

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// BenchmarkGetDatacenters benchmarks the datacenter endpoint
func BenchmarkGetDatacenters(b *testing.B) {
	// Mock response with multiple datacenters
	mockResponse := make([]DatacenterDTO, 50) // Create 50 mock datacenters
	for i := 0; i < 50; i++ {
		mockResponse[i] = DatacenterDTO{
			Id:             String("datacenter-" + string(rune(i))),
			CountryCode:    "US",
			CityCode:       "NYC",
			Index:          int64(i),
			Tier:           int64(1),
			Certifications: []string{"SOC2", "ISO27001"},
			Slug:           "us-nyc-" + string(rune(i)),
		}
	}

	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(mockResponse)
	}))
	defer server.Close()

	// Create client
	apiKey := "test-api-key"
	client, err := NewClient(&server.URL, &apiKey)
	if err != nil {
		b.Fatalf("Failed to create client: %v", err)
	}

	// Reset timer and run benchmark
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := client.GetDatacenters()
		if err != nil {
			b.Fatalf("GetDatacenters failed: %v", err)
		}
	}
}

// BenchmarkListVmsV3WithParams benchmarks the enhanced VM listing
func BenchmarkListVmsV3WithParams(b *testing.B) {
	// Mock response with multiple VMs
	mockResponse := make([]RunningInstanceV3, 20) // Create 20 mock VMs
	for i := 0; i < 20; i++ {
		mockResponse[i] = RunningInstanceV3{
			Id:              "vm-" + string(rune(i)),
			Status:          VmStatusActive,
			PricePerEpoch:   "1.23",
			Resources:       []VmResource{},
			CreatedAt:       "2024-01-01T00:00:00Z",
			NextBillingAt:   "2024-01-01T01:00:00Z",
			ReservedBalance: "10.00",
			TotalSpent:      "5.00",
			StatusChangedAt: "2024-01-01T00:00:00Z",
			TerminatedAt:    "2024-01-01T02:00:00Z",
		}
	}

	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(mockResponse)
	}))
	defer server.Close()

	// Create client
	apiKey := "test-api-key"
	client, err := NewClient(&server.URL, &apiKey)
	if err != nil {
		b.Fatalf("Failed to create client: %v", err)
	}

	// Create test parameters
	page := int64(1)
	perPage := int64(10)
	params := &ListVmsV3Params{
		Pagination: &PaginationParams{
			Page:    &page,
			PerPage: &perPage,
		},
	}

	// Reset timer and run benchmark
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := client.ListVmsV3WithParams(params)
		if err != nil {
			b.Fatalf("ListVmsV3WithParams failed: %v", err)
		}
	}
}
