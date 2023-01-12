package qiscus

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"
)

// APIResponse : is a structs that may come from Qiscus API endpoints
type APIResponse struct {
	Status     string // e.g. "200 OK"
	StatusCode int    // e.g. 200
	Proto      string // e.g. "HTTP/1.0"

	// response Header contain a map of all HTTP header keys to values.
	Header http.Header
	// response body
	RawBody []byte
	// request that was sent to obtain the response
	Request *http.Request
}

// newAPIResponse : internal function to set HTTP Raw response return to APIResponse
func newAPIResponse(res *http.Response, responseBody []byte) *APIResponse {
	return &APIResponse{
		Status:     res.Status,
		StatusCode: res.StatusCode,
		Proto:      res.Proto,
		Header:     res.Header,
		RawBody:    responseBody,
		Request:    res.Request,
	}
}

type HttpRequest interface {
	DoRequest() *Error
	AddHeader(name, value string)
	AddParameter(name, value string)
}

// HttpRequestImpl : this is for Qiscus HttpClient Implementation
type HttpRequestImpl struct {
	Method     string
	URL        string
	Body       io.Reader
	Headers    map[string]string
	Parameters map[string][]string
	Response   interface{}
	HttpClient *http.Client
}

func NewHttpRequest(method string, url string, body io.Reader, response interface{}) HttpRequest {
	return &HttpRequestImpl{
		Method:     method,
		URL:        url,
		Body:       body,
		Response:   response,
		HttpClient: DefaultHttpClient,
	}
}

func (r *HttpRequestImpl) AddHeader(name, value string) {
	if r.Headers == nil {
		r.Headers = make(map[string]string)
	}
	r.Headers[name] = value
}

func (r *HttpRequestImpl) AddParameter(name, value string) {
	if r.Parameters == nil {
		r.Parameters = make(map[string][]string)
	}
	r.Parameters[name] = append(r.Parameters[name], value)
}

func (r *HttpRequestImpl) DoRequest() *Error {
	// NewRequest is used by Call to generate an http.Request.
	req, err := http.NewRequest(r.Method, r.URL, r.Body)
	if err != nil {
		return &Error{
			Message:  fmt.Sprintf("error request creation failed: %s", err.Error()),
			RawError: err,
		}
	}

	// Get request body.
	// The request body after calling Do() function, will be closed by the underlying Transport, even on errors.
	var reqBody []byte
	if req.Body != nil {
		reqBody, _ = io.ReadAll(req.Body)
		// Restore the io.ReadCloser to its original state
		req.Body = io.NopCloser(bytes.NewBuffer(reqBody))
	}

	// Set Parameters
	if r.Parameters != nil && len(r.Parameters) > 0 {
		params := req.URL.Query()
		for name, values := range r.Parameters {
			for _, value := range values {
				params.Add(name, value)
			}
		}
		req.URL.RawQuery = params.Encode()
	}

	// Set Headers
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "Qiscus-Go/"+LibraryVersion)

	if r.Headers != nil && len(r.Headers) > 0 {
		for header, value := range r.Headers {
			req.Header.Add(header, value)
		}
	}

	start := time.Now()
	res, err := r.HttpClient.Do(req)
	if err != nil {
		return &Error{
			Message:    fmt.Sprintf("error when request via http client, cannot send request with error: %s", err.Error()),
			StatusCode: res.StatusCode,
			RawError:   err,
		}
	}

	defer res.Body.Close()
	latency := time.Since(start)

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return &Error{
			Message:    "cannot read response body: " + err.Error(),
			StatusCode: res.StatusCode,
			RawError:   err,
		}
	}

	if DefaultHttpOutboundLog {
		compactBody := func(data []byte) string {
			var js map[string]interface{}
			if json.Unmarshal(data, &js) != nil {
				return string(data)
			}

			result := new(bytes.Buffer)
			if err := json.Compact(result, data); err != nil {
				fmt.Println(err)
			}
			return result.String()
		}

		// Write http outbound log
		log.Info().
			Str("method", res.Request.Method).
			Str("url", res.Request.URL.String()).
			Str("body", compactBody(reqBody)).
			Int("status", res.StatusCode).
			Str("response", compactBody(resBody)).
			Dur("latency", latency).
			Msg("OUTBOUND LOG")
	}

	rawResponse := newAPIResponse(res, resBody)

	if r.Response != nil {
		if err = json.Unmarshal(resBody, &r.Response); err != nil {
			return &Error{
				Message:        fmt.Sprintf("invalid body response, parse error during api request to qiscus with message: %s", err.Error()),
				StatusCode:     res.StatusCode,
				RawError:       err,
				RawApiResponse: rawResponse,
			}
		}
	}

	// Check StatusCode from Qiscus HTTP response api StatusCode
	if res.StatusCode >= 400 {
		return &Error{
			Message:        fmt.Sprintf("qiscus api is returning error. http status code: %s  api response: %s", strconv.Itoa(res.StatusCode), string(resBody)),
			StatusCode:     res.StatusCode,
			RawError:       err,
			RawApiResponse: rawResponse,
		}
	}

	return nil
}
