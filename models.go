package fluenceapi

// String is a type alias for string, used for OpenAPI compatibility
type String string

// ChainAddress is a type alias for string, used for blockchain addresses
type ChainAddress string

// VmId is a type alias for string, representing UUID format for VM IDs
type VmId string

// ErrorBody represents the error response format
type ErrorBody struct {
	Error string `json:"error"`
}

// AddSshKey represents the request body for adding an SSH key
type AddSshKey struct {
	Name      string `json:"name"`
	PublicKey string `json:"publicKey"`
}

// SshKey represents an SSH key object
type SshKey struct {
	Fingerprint string  `json:"fingerprint"`
	Algorithm   string  `json:"algorithm"`
	Comment     string  `json:"comment"`
	PublicKey   string  `json:"public_key"` // Note: uses underscore in API
	Active      bool    `json:"active"`
	CreatedAt   string  `json:"createdAt"`
	Name        *string `json:"name,omitempty"` // Optional field
}

// RemoveSshKey represents the request body for removing an SSH key
type RemoveSshKey struct {
	Fingerprint string `json:"fingerprint"`
}

// DatacenterConstraint represents a constraint on datacenter countries
type DatacenterConstraint struct {
	Countries []string `json:"countries"`
}

// AdditionalStorage represents additional storage resources
type AdditionalStorage struct {
	Supply uint64 `json:"supply"`
	Units  string `json:"units"`
	Type   string `json:"type"`
}

// AdditionalResources represents additional resources (currently only storage)
type AdditionalResources struct {
	Storage []AdditionalStorage `json:"storage,omitempty"`
}

// CpuHardware represents CPU hardware constraints
type CpuHardware struct {
	Architecture string `json:"architecture"`
	Manufacturer string `json:"manufacturer"`
}

// MemoryHardware represents memory hardware constraints
type MemoryHardware struct {
	Type       string `json:"type"`
	Generation string `json:"generation"`
}

// StorageHardware represents storage hardware constraints
type StorageHardware struct {
	Type string `json:"type"`
}

// HardwareConstraints represents hardware constraints for offers
type HardwareConstraints struct {
	Cpu     []CpuHardware     `json:"cpu,omitempty"`
	Memory  []MemoryHardware  `json:"memory,omitempty"`
	Storage []StorageHardware `json:"storage,omitempty"`
}

// AvailableHardware represents the response for available hardware on the marketplace
type AvailableHardware struct {
	Cpu     []CpuHardware     `json:"cpu"`
	Memory  []MemoryHardware  `json:"memory"`
	Storage []StorageHardware `json:"storage"`
}

// OfferConstraints represents constraints for marketplace offers
type OfferConstraints struct {
	AdditionalResources      *AdditionalResources  `json:"additionalResources,omitempty"`
	BasicConfiguration       *string               `json:"basicConfiguration,omitempty"`
	Datacenter               *DatacenterConstraint `json:"datacenter,omitempty"`
	Hardware                 *HardwareConstraints  `json:"hardware,omitempty"`
	MaxTotalPricePerEpochUsd *string               `json:"maxTotalPricePerEpochUsd,omitempty"`
}

// ConfigurationPrice represents a configuration and its price
type ConfigurationPrice struct {
	Slug  string `json:"slug"`
	Price string `json:"price"`
}

// Datacenter represents a datacenter object
type Datacenter struct {
	CountryCode    string   `json:"countryCode"`
	CityCode       string   `json:"cityCode"`
	CityIndex      uint32   `json:"cityIndex"`
	Tier           uint32   `json:"tier"`
	Certifications []string `json:"certifications"`
}

// DatacenterDTO represents a datacenter data transfer object
type DatacenterDTO struct {
	Id             String   `json:"id"`
	CountryCode    string   `json:"countryCode"`
	CityCode       string   `json:"cityCode"`
	Index          int64    `json:"index"`
	Tier           int64    `json:"tier"`
	Certifications []string `json:"certifications"`
	Slug           string   `json:"slug"`
}

// ResourceMetadata is a generic map for metadata (use interface{} for flexibility)
type ResourceMetadata map[string]interface{}

// Resource represents a resource object
type Resource struct {
	Type     string           `json:"type"`
	Metadata ResourceMetadata `json:"metadata"`
	Price    string           `json:"price"`
}

// Supply represents a supply object
type Supply struct {
	Supply uint64 `json:"supply"`
	Units  string `json:"units"`
}

// AdditionalSupply represents additional supply for a resource
type AdditionalSupply struct {
	Resource
	Supply
	PerVmLimit *uint64 `json:"perVmLimit,omitempty"`
}

// ServerOffering represents a server offering
type ServerOffering struct {
	AvailableBasicInstances uint64             `json:"availableBasicInstances"`
	AdditionalResources     []AdditionalSupply `json:"additionalResources"`
}

// MarketOffering represents a market offering
type MarketOffering struct {
	Configuration       ConfigurationPrice `json:"configuration"`
	Resources           []Resource         `json:"resources"`
	Datacenter          Datacenter         `json:"datacenter"`
	Servers             []ServerOffering   `json:"servers"`
	MaxAdditionalSupply []AdditionalSupply `json:"maxAdditionalSupply"`
}

// VmConfiguration represents a VM configuration
type VmConfiguration struct {
	Hostname  *string     `json:"hostname,omitempty"`
	Name      string      `json:"name"`
	OpenPorts []OpenPorts `json:"openPorts"`
	OsImage   string      `json:"osImage"`
	SshKeys   []string    `json:"sshKeys"`
}

// OpenPorts represents an open port specification
type OpenPorts struct {
	Port     uint16 `json:"port"`
	Protocol string `json:"protocol"`
}

// CreateVmV3 represents the request body for creating VMs
type CreateVmV3 struct {
	Constraints     *OfferConstraints `json:"constraints,omitempty"`
	Instances       int               `json:"instances"`
	RequestId       *string           `json:"requestId,omitempty"`
	VmConfiguration VmConfiguration   `json:"vmConfiguration"`
}

// CreatedVm represents a created VM response
type CreatedVm struct {
	VmId   string `json:"vmId"`
	VmName string `json:"vmName"`
}

// RemoveVms represents the request body for removing VMs
type RemoveVms struct {
	VmIds []string `json:"vmIds"`
}

// VmsRemoved represents the response for removed VMs
type VmsRemoved struct {
	RemovedIds   []string `json:"removed_ids"`
	Transactions []string `json:"transactions"`
}

// PortSpec represents a port specification for a running instance
type PortSpec struct {
	Port     uint16 `json:"port"`
	Protocol string `json:"protocol"`
}

// VmResource represents a VM resource
type VmResource struct {
	Supply
	Type     string           `json:"type"`
	Details  interface{}      `json:"details"`
	Metadata ResourceMetadata `json:"metadata"`
}

// RunningInstanceV3 represents a running VM instance
type RunningInstanceV3 struct {
	Id              string       `json:"id"`
	Status          string       `json:"status"`
	PricePerEpoch   string       `json:"pricePerEpoch"`
	Resources       []VmResource `json:"resources"`
	CreatedAt       string       `json:"createdAt"`
	NextBillingAt   string       `json:"nextBillingAt"`
	ReservedBalance string       `json:"reservedBalance"`
	TotalSpent      string       `json:"totalSpent"`
	StatusChangedAt string       `json:"statusChangedAt"`
	TerminatedAt    string       `json:"terminatedAt"`
	Datacenter      *Datacenter  `json:"datacenter,omitempty"`
	OsImage         *string      `json:"osImage,omitempty"`
	Ports           *[]PortSpec  `json:"ports,omitempty"`
	PublicIp        *string      `json:"publicIp,omitempty"`
	VmName          *string      `json:"vmName,omitempty"`
}

// UpdateVm represents a patch update for a VM
type UpdateVm struct {
	Id        string       `json:"id"`
	OpenPorts *[]OpenPorts `json:"openPorts,omitempty"`
	VmName    *string      `json:"vmName,omitempty"`
}

// UpdateVms represents a patch update for multiple VMs
type UpdateVms struct {
	Updates []UpdateVm `json:"updates"`
}

// DefaultImageDTO represents a default OS image
type DefaultImageDTO struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Distribution string `json:"distribution"`
	Slug         string `json:"slug"`
	DownloadUrl  string `json:"downloadUrl"`
	Username     string `json:"username"`
	CreatedAt    string `json:"createdAt"`
	UpdatedAt    string `json:"updatedAt"`
}

// EstimateDepositRequestV3 represents the request for deposit estimation
type EstimateDepositRequestV3 struct {
	Constraints *OfferConstraints `json:"constraints,omitempty"`
	Instances   int               `json:"instances"`
}

// EstimatedDepositV3DTO represents the estimated deposit response
type EstimatedDepositV3DTO struct {
	DepositAmountUsdc  string `json:"depositAmountUsdc"`
	DepositEpochs      uint64 `json:"depositEpochs"`
	TotalPricePerEpoch string `json:"totalPricePerEpoch"`
	MaxPricePerEpoch   string `json:"maxPricePerEpoch"`
	Instances          int    `json:"instances"`
}

// VmStatusInfoDTO represents VM status information
type VmStatusInfoDTO struct {
	Id              VmId    `json:"id"`
	Status          string  `json:"status"`
	StatusChangedAt string  `json:"statusChangedAt"`
	PublicIp        *string `json:"publicIp,omitempty"`
}

// Enum constants for various types
const (
	// VmStatus enum values
	VmStatusNew               = "New"
	VmStatusLaunching         = "Launching"
	VmStatusActive            = "Active"
	VmStatusSmallBalance      = "SmallBalance"
	VmStatusInsufficientFunds = "InsufficientFunds"
	VmStatusTerminated        = "Terminated"
	VmStatusStopped           = "Stopped"
	VmStatusFailed            = "Failed"

	// ResourceType enum values
	ResourceTypeVCPU             = "VCPU"
	ResourceTypeRAM              = "RAM"
	ResourceTypeStorage          = "STORAGE"
	ResourceTypePublicIP         = "PUBLIC_IP"
	ResourceTypeNetworkBandwidth = "NETWORK_BANDWIDTH"
	ResourceTypeGPU              = "GPU"

	// StorageType enum values
	StorageTypeHDD  = "HDD"
	StorageTypeSSD  = "SSD"
	StorageTypeNVMe = "NVMe"

	// NetworkType enum values
	NetworkTypeDedicated = "dedicated"
	NetworkTypeShared    = "shared"

	// Protocol enum values
	ProtocolTCP  = "tcp"
	ProtocolUDP  = "udp"
	ProtocolSCTP = "sctp"
)

// VirtualMachineStatus represents the status of a virtual machine for query parameters
type VirtualMachineStatus string

// ExtendedDealStatus represents the extended deal status for query parameters
type ExtendedDealStatus string

// PaginationParams represents pagination parameters
type PaginationParams struct {
	Page    *int64 `json:"page,omitempty"`
	PerPage *int64 `json:"per_page,omitempty"`
}
