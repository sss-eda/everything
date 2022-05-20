package instrumentation

// Client - The interface an instrument expects to be available.
type Client struct{}

func NewClient() (*Client, error) {
	return &Client{}, nil
}
