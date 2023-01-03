package rest

import (
	"encoding/json"
)

func DeleteJSON(c RawClient, path string, out any, keyAndValue ...string) error {
	body, err := c.Get(path, keyAndValue...)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, out)
}

func GetJSON(c RawClient, path string, out any, keyAndValue ...string) error {
	body, err := c.Get(path, keyAndValue...)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, out)
}

func PostJSON(c RawClient, path string, in any, out any, keyAndValue ...string) error {
	inBody := []byte{}
	if in != nil {
		var err error
		inBody, err = json.Marshal(in)
		if err != nil {
			return err
		}
	}
	outBody, err := c.Post(path, inBody, keyAndValue...)
	if err != nil {
		return err
	}
	if out != nil {
		return json.Unmarshal(outBody, out)
	}
	return nil
}

func PutJSON(c RawClient, path string, in any, out any, keyAndValue ...string) error {
	inBody := []byte{}
	if in != nil {
		var err error
		inBody, err = json.Marshal(in)
		if err != nil {
			return err
		}
	}
	outBody, err := c.Put(path, inBody, keyAndValue...)
	if err != nil {
		return err
	}
	if out != nil {
		return json.Unmarshal(outBody, out)
	}
	return nil
}

type jsonClientImpl struct {
	RawClient
}

func (c *jsonClientImpl) DeleteJSON(path string, out any, keyAndValue ...string) error {
	return DeleteJSON(c, path, out, keyAndValue...)
}

func (c *jsonClientImpl) GetJSON(path string, out any, keyAndValue ...string) error {
	return GetJSON(c, path, out, keyAndValue...)
}

func (c *jsonClientImpl) PostJSON(path string, in any, out any, keyAndValue ...string) error {
	return PostJSON(c, path, in, out, keyAndValue...)
}

func (c *jsonClientImpl) PutJSON(path string, in any, out any, keyAndValue ...string) error {
	return PutJSON(c, path, in, out, keyAndValue...)
}

type JSONClient interface {
	RawClient
	DeleteJSON(path string, out any, keyAndValue ...string) error
	GetJSON(path string, out any, keyAndValue ...string) error
	PostJSON(path string, in any, out any, keyAndValue ...string) error
	PutJSON(path string, in any, out any, keyAndValue ...string) error
}

func NewJSONClient(core Client) JSONClient {
	return &jsonClientImpl{NewRawClient(core)}
}
