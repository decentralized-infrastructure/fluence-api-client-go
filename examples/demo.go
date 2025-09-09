// Demo application for the Fluence API Go client
// This demonstrates all available endpoints and functionality
//
// Usage:
//
//	export FLUENCE_API_KEY=your-api-key
//	go run examples/demo.go
package main

import (
	"fmt"
	"os"

	fluenceapi "github.com/decentralized-infrastructure/fluence-api-client-go"
)

func main() {
	apiKey := os.Getenv("FLUENCE_API_KEY")
	if apiKey == "" {
		fmt.Println("Please set the FLUENCE_API_KEY environment variable.")
		return
	}

	client, err := fluenceapi.NewClient(nil, &apiKey)
	if err != nil {
		fmt.Printf("Error creating client: %v\n", err)
		return
	}

	// List SSH keys
	sshKeys, err := client.ListSshKeys()
	if err != nil {
		fmt.Printf("Error listing SSH keys: %v\n", err)
	} else {
		fmt.Printf("SSH Keys: %+v\n", sshKeys)
	}

	// List running VMs
	vms, err := client.ListVmsV3()
	if err != nil {
		fmt.Printf("Error listing VMs: %v\n", err)
	} else {
		fmt.Printf("Running VMs: %+v\n", vms)
	}

	// Test new endpoints
	// Get default OS images
	images, err := client.GetDefaultImages()
	if err != nil {
		fmt.Printf("Error getting default images: %v\n", err)
	} else {
		fmt.Printf("Default Images: %+v\n", images)
	}

	// Get basic configurations
	configs, err := client.GetBasicConfigurations()
	if err != nil {
		fmt.Printf("Error getting basic configurations: %v\n", err)
	} else {
		fmt.Printf("Basic Configurations: %+v\n", configs)
	}

	// Get available countries
	countries, err := client.GetAvailableCountries()
	if err != nil {
		fmt.Printf("Error getting available countries: %v\n", err)
	} else {
		fmt.Printf("Available Countries: %+v\n", countries)
	}

	// Get available hardware
	hardware, err := client.GetAvailableHardware()
	if err != nil {
		fmt.Printf("Error getting available hardware: %v\n", err)
	} else {
		fmt.Printf("Available Hardware: %+v\n", hardware)
	}

	// Get datacenters (new endpoint)
	fmt.Println("\n=== Testing New Datacenter Endpoint ===")
	datacenters, err := client.GetDatacenters()
	if err != nil {
		fmt.Printf("Error getting datacenters: %v\n", err)
	} else {
		fmt.Printf("Found %d datacenters\n", len(datacenters))
		for i, dc := range datacenters {
			if i < 3 { // Show first 3 datacenters as examples
				fmt.Printf("  Datacenter %d: %s-%s (Tier %d, Slug: %s)\n",
					i+1, dc.CountryCode, dc.CityCode, dc.Tier, dc.Slug)
			}
		}
		if len(datacenters) > 3 {
			fmt.Printf("  ... and %d more datacenters\n", len(datacenters)-3)
		}
	}

	// Test enhanced VM listing with parameters
	fmt.Println("\n=== Testing Enhanced VM Listing ===")
	// Test with pagination
	pagination := &fluenceapi.PaginationParams{
		Page:    &[]int64{1}[0],
		PerPage: &[]int64{5}[0],
	}
	params := &fluenceapi.ListVmsV3Params{
		Pagination: pagination,
	}

	vmsWithParams, err := client.ListVmsV3WithParams(params)
	if err != nil {
		fmt.Printf("Error listing VMs with parameters: %v\n", err)
	} else {
		fmt.Printf("Found %d VMs (with pagination limit of 5)\n", len(vmsWithParams))
		if len(vmsWithParams) > 0 {
			fmt.Printf("  First VM: %s (Status: %s)\n", vmsWithParams[0].Id, vmsWithParams[0].Status)
		}
	}
}
