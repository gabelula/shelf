
# request
    import "github.com/coralproject/shelf/internal/platform/request"







## type Client
``` go
type Client struct {
    BaseURL string
    Signer  auth.Signer
}
```
Client contains the necessary pieces to perform requests down to a service
layer that has platform authentication enabled.











### func (\*Client) Do
``` go
func (c *Client) Do(req *http.Request) ([]byte, error)
```
Do executes the http request on the default http client and returns the bytes
in the event that the response code was < 400.



### func (\*Client) New
``` go
func (c *Client) New(context interface{}, verb, path string, body io.Reader) (*http.Request, error)
```
New creates a new request sourced from the client. If a signer is present on
the client, requests will automatically be signed.









- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)