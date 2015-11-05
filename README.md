
# go-digits [![Build Status](https://travis-ci.org/dghubble/go-digits.png)](https://travis-ci.org/dghubble/go-digits) [![Coverage](http://gocover.io/_badge/github.com/dghubble/go-digits/digits)](http://gocover.io/github.com/dghubble/go-digits/digits) [![GoDoc](http://godoc.org/github.com/dghubble/go-digits?status.png)](http://godoc.org/github.com/dghubble/go-digits)
<img align="right" src="http://storage.googleapis.com/dghubble/digits-gopher.png">

go-digits is a Go client library for the [Digits](https://get.digits.com/) API.

To implement *Login with Digits* for web or mobile, see the [gologin](https://github.com/dghubble/gologin) package and [examples](https://github.com/dghubble/gologin/tree/master/examples/digits).

### Features

* Digits API Client
* AccountService for getting Digits accounts
* Accepts any OAuth1 `http.Client` for authorization

## Install

    go get github.com/dghubble/go-digits/digits

## Docs

Read [GoDoc](https://godoc.org/github.com/dghubble/go-digits)

## Usage

The `digits` package provides a Client for accessing Digits API services. Here is an example request for a Digit user's Account.

```go
import (
    "github.com/dghubble/go-digits/digits"
    "github.com/dghubble/oauth1"
)

config := oauth1.NewConfig("consumerKey", "consumerSecret")
token := oauth1.NewToken("accessToken", "accessSecret")
// OAuth1 http.Client will automatically authorize Requests
httpClient := config.Client(oauth1.NoContext, token)

// Digits client
client := digits.NewClient(httpClient)
// get current user's Digits Account
account, resp, err := client.Accounts.Account()
```

The API client accepts any `http.Client` capable of signing OAuth1 requests to handle authorization.

See the OAuth1 package [dghubble/oauth1](https://github.com/dghubble/oauth1) for details and examples.

## Contributing

See the [Contributing Guide](https://gist.github.com/dghubble/be682c123727f70bcfe7).

## License

[MIT License](LICENSE)


