// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package chart

import (
	"encoding/xml"

	"github.com/unidoc/unioffice"
)

type CT_View3D struct {
	RotX         *CT_RotX
	HPercent     *CT_HPercent
	RotY         *CT_RotY
	DepthPercent *CT_DepthPercent
	RAngAx       *CT_Boolean
	Perspective  *CT_Perspective
	ExtLst       *CT_ExtensionList
}

func NewCT_View3D() *CT_View3D {
	ret := &CT_View3D{}
	return ret
}

func (m *CT_View3D) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.RotX != nil {
		serotX := xml.StartElement{Name: xml.Name{Local: "c:rotX"}}
		e.EncodeElement(m.RotX, serotX)
	}
	if m.HPercent != nil {
		sehPercent := xml.StartElement{Name: xml.Name{Local: "c:hPercent"}}
		e.EncodeElement(m.HPercent, sehPercent)
	}
	if m.RotY != nil {
		serotY := xml.StartElement{Name: xml.Name{Local: "c:rotY"}}
		e.EncodeElement(m.RotY, serotY)
	}
	if m.DepthPercent != nil {
		sedepthPercent := xml.StartElement{Name: xml.Name{Local: "c:depthPercent"}}
		e.EncodeElement(m.DepthPercent, sedepthPercent)
	}
	if m.RAngAx != nil {
		serAngAx := xml.StartElement{Name: xml.Name{Local: "c:rAngAx"}}
		e.EncodeElement(m.RAngAx, serAngAx)
	}
	if m.Perspective != nil {
		seperspective := xml.StartElement{Name: xml.Name{Local: "c:perspective"}}
		e.EncodeElement(m.Perspective, seperspective)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "c:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_View3D) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_View3D:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "rotX"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "rotX"}:
				m.RotX = NewCT_RotX()
				if err := d.DecodeElement(m.RotX, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "hPercent"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "hPercent"}:
				m.HPercent = NewCT_HPercent()
				if err := d.DecodeElement(m.HPercent, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "rotY"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "rotY"}:
				m.RotY = NewCT_RotY()
				if err := d.DecodeElement(m.RotY, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "depthPercent"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "depthPercent"}:
				m.DepthPercent = NewCT_DepthPercent()
				if err := d.DecodeElement(m.DepthPercent, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "rAngAx"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "rAngAx"}:
				m.RAngAx = NewCT_Boolean()
				if err := d.DecodeElement(m.RAngAx, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "perspective"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "perspective"}:
				m.Perspective = NewCT_Perspective()
				if err := d.DecodeElement(m.Perspective, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_View3D %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_View3D
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_View3D and its children
func (m *CT_View3D) Validate() error {
	return m.ValidateWithPath("CT_View3D")
}

// ValidateWithPath validates the CT_View3D and its children, prefixing error messages with path
func (m *CT_View3D) ValidateWithPath(path string) error {
	if m.RotX != nil {
		if err := m.RotX.ValidateWithPath(path + "/RotX"); err != nil {
			return err
		}
	}
	if m.HPercent != nil {
		if err := m.HPercent.ValidateWithPath(path + "/HPercent"); err != nil {
			return err
		}
	}
	if m.RotY != nil {
		if err := m.RotY.ValidateWithPath(path + "/RotY"); err != nil {
			return err
		}
	}
	if m.DepthPercent != nil {
		if err := m.DepthPercent.ValidateWithPath(path + "/DepthPercent"); err != nil {
			return err
		}
	}
	if m.RAngAx != nil {
		if err := m.RAngAx.ValidateWithPath(path + "/RAngAx"); err != nil {
			return err
		}
	}
	if m.Perspective != nil {
		if err := m.Perspective.ValidateWithPath(path + "/Perspective"); err != nil {
			return err
		}
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
