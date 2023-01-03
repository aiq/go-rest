package rest

func Delete(c Client, path string, keyAndValue ...string) ([]byte, error) {
	params, err := FoldParameters(keyAndValue...)
	if err != nil {
		return []byte{}, err
	}
	req := &Request{
		Method:     "GET",
		Path:       path,
		Parameters: params,
	}
	return c.Do(req)
}

func Get(c Client, path string, keyAndValue ...string) ([]byte, error) {
	params, err := FoldParameters(keyAndValue...)
	if err != nil {
		return []byte{}, err
	}
	req := &Request{
		Method:     "GET",
		Path:       path,
		Parameters: params,
	}
	return c.Do(req)
}

func Post(c Client, path string, body []byte, keyAndValue ...string) ([]byte, error) {
	params, err := FoldParameters(keyAndValue...)
	if err != nil {
		return []byte{}, err
	}
	req := &Request{
		Method:     "POST",
		Path:       path,
		Body:       body,
		Parameters: params,
	}
	return c.Do(req)
}

func Put(c Client, path string, body []byte, keyAndValue ...string) ([]byte, error) {
	params, err := FoldParameters(keyAndValue...)
	if err != nil {
		return []byte{}, err
	}
	req := &Request{
		Method:     "PUT",
		Path:       path,
		Body:       body,
		Parameters: params,
	}
	return c.Do(req)
}

type rawClientImpl struct {
	Client
}

func (c *rawClientImpl) Delete(path string, keyAndValue ...string) ([]byte, error) {
	return Delete(c, path, keyAndValue...)
}

func (c *rawClientImpl) Get(path string, keyAndValue ...string) ([]byte, error) {
	return Get(c, path, keyAndValue...)
}

func (c *rawClientImpl) Post(path string, body []byte, keyAndValue ...string) ([]byte, error) {
	return Post(c, path, body, keyAndValue...)
}

func (c *rawClientImpl) Put(path string, body []byte, keyAndValue ...string) ([]byte, error) {
	return Put(c, path, body, keyAndValue...)
}

type RawClient interface {
	Client
	Delete(path string, keyAndValue ...string) ([]byte, error)
	Get(path string, keyAndValue ...string) ([]byte, error)
	Post(path string, body []byte, keyAndValue ...string) ([]byte, error)
	Put(path string, body []byte, keyAndValue ...string) ([]byte, error)
}

func NewRawClient(core Client) RawClient {
	return &rawClientImpl{core}
}
