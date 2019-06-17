[![Codefresh build status](https://g.codefresh.io/api/badges/pipeline/ameier38/ameier38%2Fonepassword%2Fonepassword?branch=master&key=eyJhbGciOiJIUzI1NiJ9.NWMzMjE0ODA3YTJkOGI3ZjkxMzVhZjlm.WFn4I6XuUDBfWsKEp6LIuG-IlDsT4JCDTjMzeH7kGu8&type=cf-1)]( https://g.codefresh.io/pipelines/onepassword/builds?repoOwner=ameier38&repoName=onepassword&serviceName=ameier38%2Fonepassword&filter=trigger:build~Build;branch:master;pipeline:5d079684c8d990545f03f911~onepassword)
[![Go Report Card](https://goreportcard.com/badge/github.com/ameier38/onepassword-client)](https://goreportcard.com/report/github.com/ameier38/onepassword-client)

# 1Password Client
Thin wrapper around the 1Password CLI for use in Golang.

__This is a stub and should not be used in production.__
> Hopefully 1Password creates their own client library. See 
[discussion](https://discussions.agilebits.com/discussion/80802/why-a-cli-and-not-a-library).

## Usage
First [install the 1Password CLI](https://support.1password.com/command-line/).

Import the package, create a client, and retrieve an item.
```go
import (
    "os"

    op "github.com/ameier38/onepassword"
)

func main() {
    password := os.GetEnv("OP_PASSWORD")
    secretKey := os.GetEnv("OP_SECRET_KEY")
    client := op.NewClient("op", "subdomain", "test@subdomain.com", password, secretKey)
    item := client.GetItem(VaultName("test-vault"), ItemName("test-item"))
}
```

## Resources
- [terraform-provider-onepassword](https://github.com/anasinnyk/terraform-provider-1password)
- [1Password CLI](https://support.1password.com/command-line/)
