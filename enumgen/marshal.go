// Copyright (c) 2023, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Based on http://github.com/dmarkham/enumer and
// golang.org/x/tools/cmd/stringer:

// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package enumgen

// Arguments to format are:
//
//	[1]: type name
const JSONMethods = `
// MarshalJSON implements the [json.Marshaler] interface.
func (i %[1]s) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the [json.Unmarshaler] interface.
func (i *%[1]s) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return errors.New("%[1]s should be a string, but got " + string(data) + "instead")
	}
	return i.SetString(s)
}
`

func (g *Generator) BuildJSONMethods(runs [][]Value, typeName string, runsThreshold int) {
	g.Printf(JSONMethods, typeName)
}

// Arguments to format are:
//
//	[1]: type name
const TextMethods = `
// MarshalText implements the [encoding.TextMarshaler] interface.
func (i %[1]s) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *%[1]s) UnmarshalText(text []byte) error {
	return i.SetString(string(text))
}
`

func (g *Generator) BuildTextMethods(runs [][]Value, typeName string, runsThreshold int) {
	g.Printf(TextMethods, typeName)
}

// Arguments to format are:
//
//	[1]: type name
const YAMLMethods = `
// MarshalYAML implements a YAML Marshaler.
func (i %[1]s) MarshalYAML() (any, error) {
	return i.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler.
func (i *%[1]s) UnmarshalYAML(unmarshal func(any) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}
	return i.SetString(s)
}
`

func (g *Generator) BuildYAMLMethods(runs [][]Value, typeName string, runsThreshold int) {
	g.Printf(YAMLMethods, typeName)
}