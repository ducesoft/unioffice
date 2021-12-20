// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

package chart

import "github.com/unidoc/unioffice/color"

var defaultColors = []color.Color{color.RGB(0x33, 0x66, 0xcc), color.RGB(0xDC, 0x39, 0x12), color.RGB(0xFF, 0x99, 0x00),
	color.RGB(0x10, 0x96, 0x18), color.RGB(0x99, 0x00, 0x99), color.RGB(0x3B, 0x3E, 0xAC), color.RGB(0x00, 0x99, 0xC6),
	color.RGB(0xDD, 0x44, 0x77), color.RGB(0x66, 0xAA, 0x00), color.RGB(0xB8, 0x2E, 0x2E), color.RGB(0x31, 0x63, 0x95),
	color.RGB(0x99, 0x44, 0x99), color.RGB(0x22, 0xAA, 0x99), color.RGB(0xAA, 0xAA, 0x11), color.RGB(0x66, 0x33, 0xCC),
	color.RGB(0xE6, 0x73, 0x00), color.RGB(0x8B, 0x07, 0x07), color.RGB(0x32, 0x92, 0x62), color.RGB(0x55, 0x74, 0xA6),
	color.RGB(0x3B, 0x3E, 0xAC)}

type chartBase struct {
}

func (c chartBase) nextColor(idx int) color.Color {
	return defaultColors[idx%len(defaultColors)]
}
