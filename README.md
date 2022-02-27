# OAuth1

Modified to export some usefull methods in order to not re-create an http.Client instance
for each authorized user. Forked from [dghubble/oauth1)](https://github.com/dghubble/oauth1).

```go
import "github.com/kataras/iris/v12/x/client"
```

```go
var myClient = client.New(client.BaseURL("https://apis.garmin.com"))
```

```go
import "github.com/iris-contrib/oauth1"

var config = &oauth1.Config{
	ConsumerKey:    "xxx",
	ConsumerSecret: "xxx",
	CallbackURL:    "http://localhost:8080/callback",
	Endpoint: oauth1.Endpoint{
		RequestTokenURL: "https://connectapi.garmin.com/oauth-service/oauth/request_token",
		AuthorizeURL:    "https://connect.garmin.com/oauthConfirm",
		AccessTokenURL:  "https://connectapi.garmin.com/oauth-service/oauth/access_token",
	},
}
```

```go
func testPreFilledAccessToken(ctx iris.Context) {
	var (
		accessToken  = "xxx"
		accessSecret = "xxx"
	)

	endpoint := "xxx"
	opt := oauth1.RequestOption(config, accessToken, accessSecret)

	var resp interface{}
	err := garminClient.ReadJSON(ctx, &resp, iris.MethodGet, endpoint, nil, opt)
	if err != nil {
        ctx.StopWithError(iris.StatusBadGateway, err)
		return
	}

	ctx.JSON(resp)
}
```


## Callback

```go
func requestToken(ctx iris.Context) {
	requestToken, requestSecret, err := config.RequestToken()
	if err != nil {
		ctx.Application().Logger().Errorf("request token: %s", err.Error())
		return
	}

	authorizationURL, err := config.AuthorizationURL(requestToken)
	if err != nil {
		ctx.Application().Logger().Errorf("authorize: %s", err.Error())
		return
	}

	// You have to keep "requestSecret" for the next request, it's up to you.
	ctx.Redirect(authorizationURL.String())
}
```

```go
func oauth1Callback(ctx iris.Context) {
	requestToken, verifier, err := oauth1.ParseAuthorizationCallback(ctx.Request())
	if err != nil {
		ctx.Application().Logger().Errorf("callback: parse auth callback: %s", err.Error())
		return
	}

	// Pass it through url parameters or anything, 
    // just fill it with the previous handler's result.
	var requestSecret string

	accessToken, accessSecret, err := config.AccessToken(requestToken, requestSecret, verifier)
	if err != nil {
		ctx.Application().Logger().Errorf("callback: access token: %s", err.Error())
		return
	}
}
```
