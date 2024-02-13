package client

import (
	"fmt"
	"math"
	"net/http"
	"time"
)

const totalRetryAttempts = 3

// shouldRetry checks if retry should be performed
func shouldRetry(err error, resp *http.Response) bool {
	return err != nil || resp.StatusCode == http.StatusBadGateway || resp.StatusCode == http.StatusGatewayTimeout
}

// timeout sets pause between retry request
func timeout(retryCount int) time.Duration {
	return time.Duration(math.Pow(2, float64(retryCount))) * time.Second
}

type retryableTransport struct {
	transport http.RoundTripper
}

func (t *retryableTransport) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	retryCount := 0
	resp, err = t.transport.RoundTrip(req)
	for shouldRetry(err, resp) && retryCount < totalRetryAttempts {
		pause := timeout(retryCount)
		fmt.Printf("Request returned %v --> Attepmting to retry request in %v\n", resp.StatusCode, pause)
		time.Sleep(pause)
		// retry
		resp, err = t.transport.RoundTrip(req)
		retryCount++
	}
	return
}

func NewRetryableClient() *http.Client {
	transport := &retryableTransport{transport: &http.Transport{}}
	return &http.Client{Transport: transport}
}
