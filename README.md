# What is go-requests?
It's a package that makes it easy to send HTTP requests using **Go**. I had the idea to code this package because I was tired to write the same code every time I wanted to interact with an external server and I thought it can be useful for someone else with the same problem

# Installation
`go get -u github.com/jaavier/go-requests`

### How to start using it? 
Methods allowed:
- GET
- POST (JSON or Form)
- PUT
- DELETE

### Structures
- Headers
```golang
type Header struct {
	Key   string
	Value string
}
```

- Request
```golang
type Request struct {
	Url      string
	Method   string
	Headers  []Header
	BodyJSON interface{}
	Form     url.Values
}
```
- Response
```golang
type Response struct {
	Body          string
	StatusCode    int
	Headers       http.Header
	Cookies       []*http.Cookie
	ContentLength int64
}
```
# Example Sending Request

### Get your public IP
```golang
package main

import (
	"fmt"

	gorequests "github.com/jaavier/go-requests"
)

func main() {
	res, err := gorequests.SendRequest(gorequests.Request{
		Method: "GET",
		Url:    "https://ifconfig.co",
		Headers: []gorequests.Header{
			{
				Key:   "User-Agent",
				Value: "curl/7.84.0",
			},
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Your ip is: %s", res.Body)
}
```

### Sending JSON with visitor message:
```golang
package main

import (
	"fmt"

	gorequests "github.com/jaavier/go-requests"
)

type Visitor struct {
	Message string `json:"message"`
}

func main() {
	response, err := gorequests.SendRequest(gorequests.Request{
		Method: "POST",
		Url:    "https://webhook.site/a41a84e0-1659-4abb-8e06-6a3a1e23e95c",
		BodyJSON: Visitor{
			Message: "Hello from github.com/jaavier/go-requests!",
		},
		Headers: []gorequests.Header{
			{
				Key:   "Content-Type",
				Value: "application/json",
			},
		},
	})
	
	if err != nil {
		panic(err)
	}
	
	fmt.Println(response.Body)
}
```
### Sending Form with visitor message:
```golang
package main

import (
	"fmt"
	"net/url"

	gorequests "github.com/jaavier/go-requests"
)

type Visitor struct {
	Message string `json:"message"`
}

func main() {
	var form url.Values = url.Values{
		"message": {"Hello from github.com/jaavier/requests"},
	}
	
	response, err := gorequests.SendRequest(gorequests.Request{
		Method: "POST",
		Url:    "https://webhook.site/a41a84e0-1659-4abb-8e06-6a3a1e23e95c",
		Form:   form,
	})
	
	if err != nil {
		panic(err)
	}
	
	fmt.Println(response.StatusCode)
}
```
