package api

import (
	"encoding/json"
	"io"
)

type Codec interface {
	Encode(io.Writer, interface{}) error
	Decode(io.Reader, interface{}) error
}

var DefaultJSONCode = jsonCodec{}

type jsonCodec struct{}

func (_ jsonCodec) Decode(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}

func (_ jsonCodec) Encode(r io.Writer, v interface{}) error {
	return json.NewEncoder(r).Encode(v)
}
