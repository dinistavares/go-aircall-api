# go-aircall-api

A Aircall API Golang Wrapper for the [Aircall API](https://developer.aircall.io/api-references)

# Install

```console
$ go get github.com/dinistavares/go-aircall-api
```

# Usage

Import the Aircall package.

```go
import "github.com/dinistavares/go-aircall-api"
```

Create a new Aircall Client and use the provided services.

```go
  client := aircall.New()

  // List Aircall calls
  calls, _, err := client.Call.List(nil)
```

## Authenticate

Authenticate with access token for OAuth app

```go
import (
  "github.com/dinistavares/go-aircall-api"
)

func main(){
  accessToken := "xxxxxxx"

  client := aircall.New()
  client.Authenticate(accessToken)

  // List Aircall calls
  calls, _, err := client.Call.List(nil)
}
```

Authenticate with api_key and api_token for Aircall customers

```go
import (
  "github.com/dinistavares/go-aircall-api"
)

func main(){
  apiKey := "xxxxxxx"
  apiToken := "xxxxxxx"

  client := aircall.New()
  client.AuthenticateBasic(apiKey, apiToken)

  // List Aircall calls
  calls, _, err := client.Call.List(nil)
}
```

# Examples

### Calls

**List calls**
```go
  // Create query filter and set 'from' and 'to'
  opts := client.Call.Query().NewListCalls()
  opts.From("1728722536")
  opts.To("1728922536")

  response, _, err := client.Call.List(opts)
```

**Get call by ID**
```go
  id := 23456

  response, _, err := client.Call.GetByID(id)
```

**Search calls**
```go
  opts := client.Call.Query().NewSearchCalls()
  opts.From("1728722536")
  opts.To("1728922536")
  opts.Direction("inbound")

  response, _, err := client.Call.Search(opts)
```
