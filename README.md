# espv2protoyaml

You may want this to build api_config.yaml [programmatically](https://cloud.google.com/endpoints/docs/grpc/get-started-cloud-run#configure_esp). Enumerating RPC calls is your own bag.

## Usage

```bash
./espv2protoyaml \
    -backend helloworld.Greeter \
    -endpoint-name hellogrpc.endpoints.YOUR_PROJECT_ID.cloud.goog \
    -endpoint-title helloworld.Greeter \
    -usage-rule hellogrpc.endpoints.YOUR_PROJECT_ID.cloud.goog.Greeter,true \ # selector,unregisteredCallFlag
    -backend-rule '*',grpcs://<CLOUD_RUN_NAME>-<HASH>-uc.a.run.app # selector,address
```

### Flags

```bash
  -backend value
    gRPC service name. Can be repeated
  -backend-rule value
    Comma separated RPC address and selector. Can be repeated.
  -config-version int
    config version (default 3)
  -endpoint-name string
    GCP endpoint name
  -endpoint-title string
    GCP endpoint title
  -h string
    Prints this message
  -o string
    API Config path (default: api_config.yaml). '-' works. (default "api_config.yaml")
  -service-type string
    service type (default "google.api.Service")
  -usage-rule value
    Comma separated RPC name and 'allow unregistered call' value. Can be repeated..
```

### API

```go
var c espv2protoyaml.Espv2Config

backends := []string{"backend1", "backend2"}
usageRules := []struct{
  Selector string
  AllowUnregistered bool
}{
  {
    Selector: "...",
    AllowUnregistered: true,
  }
}

backendRules := []struct{
  Selector string
  Address string
}{
  {
    Selector: "*",
    Address: "...",
  }
}

c.SetServiceType("...")
c.SetConfigVersion("...")
c.SetEndpointName("...")
c.SetEndpointTitle("...")
c.AppendBackend(backends...)

for _, r := range usageRules {
  c.AppendUsageRule(r.Selector, r.AllowUnregistered)
}

for _, r := range backendRules {
  c.AppendBackendRule(r.Selector, r.Address)
}

_ = c.WriteConfig(os.Stdout)
```
