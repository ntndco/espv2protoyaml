// Package espv2protoyaml describes ...
package espv2protoyaml

import (
	"io"

	"gopkg.in/yaml.v2"
)

const (
	// DefaultServiceType ...
	DefaultServiceType string = "google.api.Service"
	// DefaultConfigVersion ...
	DefaultConfigVersion int = 3
)

// Espv2Config a struct for creating the API config manifest
type Espv2Config struct {
	ServiceType   string        `yaml:"type"`
	ConfigVersion int           `yaml:"config_version"`
	Name          string        `yaml:"name"`
	Title         string        `yaml:"title"`
	APIs          []BackendName `yaml:"apis"`
	Usage         struct {
		Rules []UsageRule `yaml:"rules"`
	} `yaml:"usage"`
	Backend struct {
		Rules []BackendRule `yaml:"rules"`
	} `yaml:"backend"`
}

// BackendName provides the name for the gRPC service
type BackendName struct {
	Name string
}

// UsageRule required for creating API config manifest
type UsageRule struct {
	Selector               string `yaml:"selector"`
	AllowUnregisteredCalls bool   `yaml:"allow_unregistered_calls"`
}

// BackendRule required for creating API config manifest
type BackendRule struct {
	Selector string `yaml:"selector"`
	Address  string `yaml:"address"`
}

// SetServiceType ...
func (c *Espv2Config) SetServiceType(t string) {
	st := t
	if t == "" {
		st = DefaultServiceType
	}
	c.ServiceType = st
}

// SetConfigVersion ...
func (c *Espv2Config) SetConfigVersion(v int) {
	cv := v
	if v == 0 {
		cv = DefaultConfigVersion
	}
	c.ConfigVersion = cv
}

// SetEndpointName ...
func (c *Espv2Config) SetEndpointName(v string) {
	c.Name = v
}

// SetEndpointTitle ...
func (c *Espv2Config) SetEndpointTitle(v string) {
	c.Title = v
}

// AppendBackend ...
func (c *Espv2Config) AppendBackend(v ...string) {
	var bns []BackendName
	for _, bn := range v {
		b := BackendName{Name: bn}
		bns = append(bns, b)
	}
	c.APIs = append(c.APIs, bns...)
}

// AppendBackendRule ...
func (c *Espv2Config) AppendBackendRule(selector, address string) {
	r := BackendRule{
		Selector: selector,
		Address:  address,
	}
	c.Backend.Rules = append(c.Backend.Rules, r)
}

// AppendUsageRule ...
func (c *Espv2Config) AppendUsageRule(selector string, auc bool) {
	r := UsageRule{
		Selector:               selector,
		AllowUnregisteredCalls: auc,
	}
	c.Usage.Rules = append(c.Usage.Rules, r)
}

// WriteConfig writes config to sink
func (c *Espv2Config) WriteConfig(w io.Writer) error {
	enc := yaml.NewEncoder(w)
	err := enc.Encode(c)
	if err != nil {
		return err
	}

	err = enc.Close()

	return err
}

// DefaultData for building the API config
var DefaultData = `
type:
config_version:
name:
title:
apis:
usage:
  rules:
backend:
  rules:
`
