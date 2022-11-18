# What is go-requests?
It's a package that makes it easy to send HTTP requests using **Go**. I had the idea to code this package because I was tired to write the same code every time I wanted to interact with an external server and I thought it can be useful for someone else with the same problem

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
### Sending JSON with visitor message:
```golang
type Visitor struct {
    Message string `json:"message"`
}
```
```golang
response, err := SendRequest(Request{
	Method: "POST",
	Url:    "https://webhook.site/a41a84e0-1659-4abb-8e06-6a3a1e23e95c",
	BodyJSON: Visitor{
		Message: "Hello from github.com/jaavier/go-requests!",
	},
	Headers: []Header{
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
```
### Sending Form with visitor message:
```golang
var form url.Values = url.Values{
	"message": {"Hello from github.com/jaavier/requests"},
}

response, err := SendRequest(Request{
	Method: "POST",
	Url:    "https://faas-fra1-afec6ce7.doserverless.co/api/v1/web/fn-95a06ffa-77ad-4531-8214-8b96897b95d7/default/delete-soon",
	Form:   form,
})

if err != nil {
	panic(err)
}

fmt.Println(response.Body)
```
