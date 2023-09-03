// Copyright (c) 2023, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package enumgen

import (
	"os"
	"strings"
	"testing"

	"goki.dev/enums/enumgen/config"
	"goki.dev/ki/v2/kit"
)

func TestGenerate(t *testing.T) {
	c := &config.Config{}
	err := kit.SetFromDefaultTags(c)
	if err != nil {
		t.Errorf("programmer error: error setting config from default tags: %v", err)
	}
	c.Dir = "./testdata"
	c.Output = "./testdata/enumgen.go"
	err = Generate(c)
	if err != nil {
		t.Errorf("error while generating: %v", err)
	}
	have, err := os.ReadFile("testdata/enumgen.go")
	if err != nil {
		t.Errorf("error while reading generated file: %v", err)
	}
	want, err := os.ReadFile("testdata/enumgen.golden")
	if err != nil {
		t.Errorf("error while reading golden file: %v", err)
	}
	// ignore first line, which has "Code generated by" message
	// that can change based on where go test is ran.
	_, shave, got := strings.Cut(string(have), "\n")
	if !got {
		t.Errorf("expected string with newline in testdata/enumgen.go, but got %q", have)
	}
	_, swant, got := strings.Cut(string(want), "\n")
	if !got {
		t.Errorf("expected string with newline in testdata/enumgen.golden, but got %q", want)
	}
	if shave != swant {
		t.Errorf("expected generated file and expected file to be the same after the first line, but they are not (compare ./testdata/enumgen.go and ./testdata/enumgen.golden to see the difference)")
	}
}
