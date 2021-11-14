// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/turnkeyca/monolith/integration/client/auth"
	"github.com/turnkeyca/monolith/integration/client/employment"
	"github.com/turnkeyca/monolith/integration/client/permission"
	"github.com/turnkeyca/monolith/integration/client/pet"
	"github.com/turnkeyca/monolith/integration/client/reference"
	"github.com/turnkeyca/monolith/integration/client/roommate"
	"github.com/turnkeyca/monolith/integration/client/shorturl"
	"github.com/turnkeyca/monolith/integration/client/user"
)

// Default of turnkey API HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http"}

// NewHTTPClient creates a new of turnkey API HTTP client.
func NewHTTPClient(formats strfmt.Registry) *OfTurnkeyAPI {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new of turnkey API HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *OfTurnkeyAPI {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new of turnkey API client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *OfTurnkeyAPI {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(OfTurnkeyAPI)
	cli.Transport = transport
	cli.Auth = auth.New(transport, formats)
	cli.Employment = employment.New(transport, formats)
	cli.Permission = permission.New(transport, formats)
	cli.Pet = pet.New(transport, formats)
	cli.Reference = reference.New(transport, formats)
	cli.Roommate = roommate.New(transport, formats)
	cli.Shorturl = shorturl.New(transport, formats)
	cli.User = user.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// OfTurnkeyAPI is a client for of turnkey API
type OfTurnkeyAPI struct {
	Auth auth.ClientService

	Employment employment.ClientService

	Permission permission.ClientService

	Pet pet.ClientService

	Reference reference.ClientService

	Roommate roommate.ClientService

	Shorturl shorturl.ClientService

	User user.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *OfTurnkeyAPI) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.Auth.SetTransport(transport)
	c.Employment.SetTransport(transport)
	c.Permission.SetTransport(transport)
	c.Pet.SetTransport(transport)
	c.Reference.SetTransport(transport)
	c.Roommate.SetTransport(transport)
	c.Shorturl.SetTransport(transport)
	c.User.SetTransport(transport)
}
