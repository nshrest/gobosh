# gobosh - BOSH client API for golang applications

This project is a golang library for applications wanting to talk to a BOSH/Micro-BOSH or bosh-lite. This is forked from https://github.com/cloudfoundry-community/gogobosh

## API

The following client functions are available, as a subset of the full BOSH Director API.

- client.GetInfo()
- client.GetStemcells()
- client.GetReleases()
- client.GetDeployments()
- client.GetDeployment("cf-warden")
- client.GetDeploymentVMs("cf-warden")
- client.GetTasks()
- client.GetTask(123)
- client.GetTaskResult(123)
- client.GetCertificates()

## Install

```
go get github.com/nshrest/gobosh
```

## Documentation

```
godoc -goroot=$GOPATH github.com/nshrest/gobosh
```

### Use

As a short getting started guide:

```
package main

import (
  "github.com/nshrest/gogobosh"
  "fmt"
)

func main() {
  c, _ := gogobosh.NewClient(gogobosh.DefaultConfig())
  info, _ := c.GetInfo()

  fmt.Println("Director")
  fmt.Printf("  Name       %s\n", info.Name)
  fmt.Printf("  Version    %s\n", info.Version)
  fmt.Printf("  User       %s\n", info.User)
  fmt.Printf("  UUID       %s\n", info.UUID)
  fmt.Printf("  CPI        %s\n", info.CPI)
}
```

## Tests

Tests  are all local currently; and do not test against a running bosh or bosh-lite.