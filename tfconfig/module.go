package tfconfig

import (
	"github.com/hashicorp/hcl/v2"
)

// Module is the top-level type representing a parsed and processed Terraform
// module.
type Module struct {
	// Path is the local filesystem directory where the module was loaded from.
	Path string `json:"path"`

	Variables map[string]*Variable `json:"variables"`
	Outputs   map[string]*Output   `json:"outputs"`

	RequiredCore      []string                        `json:"required_core,omitempty"`
	RequiredProviders map[string]*ProviderRequirement `json:"required_providers"`

	// A configuration can only provide one backend block.
	Backend *Backend

	ProviderConfigs  map[string]*ProviderConfig `json:"provider_configs,omitempty"`
	ManagedResources map[string]*Resource       `json:"managed_resources"`
	DataResources    map[string]*Resource       `json:"data_resources"`
	ModuleCalls      map[string]*ModuleCall     `json:"module_calls"`

	// Diagnostics records any errors and warnings that were detected during
	// loading, primarily for inclusion in serialized forms of the module
	// since this slice is also returned as a second argument from LoadModule.
	Diagnostics Diagnostics `json:"diagnostics,omitempty"`
}

// Backend represents a backend block in the configuration.
// https://www.terraform.io/docs/language/settings/backends/configuration.html
type Backend struct {
	Name string
	Body hcl.Body
}

// ProviderConfig represents a provider block in the configuration
type ProviderConfig struct {
	Name  string `json:"name"`
	Alias string `json:"alias,omitempty"`
}

// NewModule creates new Module representing Terraform module at the given path
func NewModule(path string) *Module {
	return &Module{
		Path:              path,
		Variables:         make(map[string]*Variable),
		Outputs:           make(map[string]*Output),
		RequiredProviders: make(map[string]*ProviderRequirement),
		ProviderConfigs:   make(map[string]*ProviderConfig),
		ManagedResources:  make(map[string]*Resource),
		DataResources:     make(map[string]*Resource),
		ModuleCalls:       make(map[string]*ModuleCall),
	}
}
