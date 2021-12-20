// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package dml_test

import (
	"encoding/xml"
	"testing"

	"github.com/unidoc/unioffice/schema/soo/dml"
)

func TestCT_PositiveSize2DConstructor(t *testing.T) {
	v := dml.NewCT_PositiveSize2D()
	if v == nil {
		t.Errorf("dml.NewCT_PositiveSize2D must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_PositiveSize2D should validate: %s", err)
	}
}

func TestCT_PositiveSize2DMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_PositiveSize2D()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_PositiveSize2D()
	xml.Unmarshal(buf, v2)
}
