package main

import (
	"errors"
	"flag"
	"io"
	"os"
	"strings"

	"github.com/sendcoffee/espv2protoyaml"
	"gopkg.in/yaml.v2"
)

var (
	espv2pyFlags  = flag.NewFlagSet("espv2protoyaml", flag.ExitOnError)
	usageMsg      = espv2pyFlags.String("h", "", "Prints this message")
	svcType       = espv2pyFlags.String("service-type", espv2protoyaml.DefaultServiceType, "service type")
	configVersion = espv2pyFlags.Int("config-version", espv2protoyaml.DefaultConfigVersion, "config version")
	endpointName  = espv2pyFlags.String("endpoint-name", "", "GCP endpoint name")
	endpointTile  = espv2pyFlags.String("endpoint-title", "", "GCP endpoint title")
	backend       backends
	ur            usageRules
	br            backendRules
	outFile       = espv2pyFlags.String("o", "api_config.yaml", "API Config path (default: api_config.yaml). '-' works.")
)

func usage() {
	espv2pyFlags.PrintDefaults()
}

func run(args []string) error {
	espv2pyFlags.Var(&backend, "backend", "gRPC service name. Can be repeated")
	espv2pyFlags.Var(&br, "backend-rule", "Comma separated RPC address and selector. Can be repeated.")
	espv2pyFlags.Var(&ur, "usage-rule", "Comma separated RPC name and 'allow unregistered call' value. Can be repeated.")
	espv2pyFlags.Usage = usage
	var c espv2protoyaml.Espv2Config
	parseErr := espv2pyFlags.Parse(args)

	if parseErr != nil {
		espv2pyFlags.Usage()
		return parseErr
	}

	if *usageMsg != "" {
		espv2pyFlags.Usage()
	}

	if *outFile == "" {
		espv2pyFlags.Usage()
		return errors.New("Must provide a path to write")
	}

	dd := strings.NewReader(espv2protoyaml.DefaultData)
	dec := yaml.NewDecoder(dd)

	var of io.Writer
	// out file when not writing to stdout
	if *outFile != "-" {
		f, err := os.Create(*outFile)
		of = f
		if err != nil {
			return err
		}
	}

	dec.Decode(&c)

	c.SetServiceType(*svcType)
	c.SetConfigVersion(*configVersion)
	c.SetEndpointName(*endpointName)
	c.SetEndpointTitle(*endpointTile)
	c.AppendBackend(backend...)

	for _, r := range ur {
		c.AppendUsageRule(r.Selector, r.AllowUnregistered)
	}

	for _, r := range br {
		c.AppendBackendRule(r.Selector, r.Address)
	}

	enc := yaml.NewEncoder(of)
	if *outFile == "-" {
		stdOutEnc := yaml.NewEncoder(os.Stdout)
		return stdOutEnc.Encode(c)
	}

	return enc.Encode(c)
}
