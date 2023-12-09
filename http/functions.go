package http

import (
	"encoding/json"
	"io"
	"net/http"
	"reflect"
	"strings"

	"cyberpull.com/gotk/v2/errors"
)

func Get[T any](url string, opts ...*RequestOptions) (data T, err error) {
	return Request[T](http.MethodGet, url, opts...)
}

func Post[T any](url string, opts ...*RequestOptions) (data T, err error) {
	return Request[T](http.MethodPost, url, opts...)
}

func Patch[T any](url string, opts ...*RequestOptions) (data T, err error) {
	return Request[T](http.MethodPatch, url, opts...)
}

func Put[T any](url string, opts ...*RequestOptions) (data T, err error) {
	return Request[T](http.MethodPut, url, opts...)
}

func Delete[T any](url string, opts ...*RequestOptions) (data T, err error) {
	return Request[T](http.MethodDelete, url, opts...)
}

func Request[T any](method, url string, opts ...*RequestOptions) (data T, err error) {
	var req *http.Request
	var resp *http.Response

	method = strings.ToUpper(method)

	opt := defaultRequestOptions(opts...)

	if req, err = http.NewRequest(method, url, opt.Body); err != nil {
		return
	}

	opt.mergeTo(req)

	if resp, err = http.DefaultClient.Do(req); err != nil {
		return
	}

	defer resp.Body.Close()

	var b []byte

	if b, err = responseData(resp); err != nil {
		return
	}

	vType := reflect.TypeOf(data)

	contentType := resp.Header.Get("Content-Type")

	if opt.ExpectsJSON || contentType == "application/json" {
		// Parse JSON Content
		switch vType.Kind() {
		case reflect.Pointer:
			data = reflect.New(vType).Interface().(T)
			err = json.Unmarshal(b, data)
		default:
			err = json.Unmarshal(b, &data)
		}

		return
	}

	// Get Content
	if vType.Kind() == reflect.String {
		data = reflect.ValueOf(string(b)).Interface().(T)
		return
	}

	err = errors.New("Invalid return type")

	return
}

// Private Functions ===========================

func trim(uri string) string {
	uri = strings.TrimSpace(uri)
	uri = strings.Trim(uri, "/")
	return uri
}

func join(paths ...string) string {
	entries := make([]string, 0)

	for _, path := range paths {
		path = trim(path)

		if path != "" {
			entries = append(entries, path)
		}
	}

	return strings.Join(entries, "/")
}

func isUrl(uri string) bool {
	return (strings.HasPrefix(uri, "https://") ||
		strings.HasPrefix(uri, "http://"))
}

func isOk(statusCode int) bool {
	return statusCode >= http.StatusOK && statusCode < 300
}

func responseData(resp *http.Response) (data []byte, err error) {
	data, err = io.ReadAll(resp.Body)

	if err != nil || isOk(resp.StatusCode) {
		return
	}

	if len(data) == 0 {
		data = []byte(resp.Status)
	}

	return
}
