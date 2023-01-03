package rest

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type foldable interface {
	url.Values | http.Header
	Add(string, string)
}

func fold[T foldable](keyAndValue ...string) (T, error) {
	n := len(keyAndValue)
	res := T{}

	if n%2 != 0 {
		return res, fmt.Errorf("%d keyAndValues entries, not every key has a matching value entry", n)
	}

	for k, v := 0, 1; v < n; k, v = k+2, v+2 {
		res.Add(keyAndValue[k], keyAndValue[v])
	}
	return res, nil
}

func FoldParameters(keyAndValue ...string) (url.Values, error) {
	return fold[url.Values](keyAndValue...)
}

func FoldHeader(keyAndValue ...string) (http.Header, error) {
	return fold[http.Header](keyAndValue...)
}

func MustFoldHeader(keyAndValue ...string) http.Header {
	header, err := FoldHeader(keyAndValue...)
	if err != nil {
		panic(err)
	}
	return header
}

func URL(fix string, tail string, values url.Values) (*url.URL, error) {
	full, err := url.JoinPath(fix, tail)
	if err != nil {
		return nil, err
	}
	res, err := url.Parse(full)
	if err != nil {
		return nil, err
	}
	res.RawQuery = values.Encode()
	return res, nil
}

type Request struct {
	// Method specifies the HTTP method (GET, POST, PUT, etc.).
	Method string

	// URL port without parameters and base.
	Path string

	// Body is the request's body.
	Body []byte

	// The URL parameters.
	Parameters url.Values
}

type Do func(*Request) ([]byte, error)

type Client interface {
	Do(*Request) ([]byte, error)
}

type wrapClient struct {
	doFunc Do
}

func (c wrapClient) Do(req *Request) ([]byte, error) {
	return c.doFunc(req)
}

func Wrap(do Do) Client {
	return wrapClient{do}
}

func ReadBytes(reader io.Reader) ([]byte, error) {
	buf := &bytes.Buffer{}
	_, err := buf.ReadFrom(reader)
	return buf.Bytes(), err
}

func Body(data []byte) io.ReadCloser {
	return io.NopCloser(bytes.NewReader(data))
}
