// Copyright 2025 Gin Core Team. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

//go:build jsoniter

package json

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// Package indicates what library is being used for JSON encoding.
const Package = "encoding/json"

func init() {
	fmt.Fprintln(os.Stderr, `[GIN-WARNING] build tag "jsoniter" is obsolete: `+
		`github.com/json-iterator/go is archived and its support was removed in gin v1.13.0, `+
		`encoding/json is used instead. Drop the tag, or plug your own codec into `+
		`gin/codec/json.API — see docs/doc.md#custom-json-codec-at-runtime`)

	API = jsoniterApi{}
}

type jsoniterApi struct{}

func (j jsoniterApi) Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

func (j jsoniterApi) Unmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}

func (j jsoniterApi) MarshalIndent(v any, prefix, indent string) ([]byte, error) {
	return json.MarshalIndent(v, prefix, indent)
}

func (j jsoniterApi) NewEncoder(writer io.Writer) Encoder {
	return json.NewEncoder(writer)
}

func (j jsoniterApi) NewDecoder(reader io.Reader) Decoder {
	return json.NewDecoder(reader)
}
