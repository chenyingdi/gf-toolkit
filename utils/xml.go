package utils

import (
	"encoding/xml"
	"io"
)

type Xml map[string]interface{}

type xmlMapEntry struct {
	XMLName xml.Name
	Value   interface{} `xml:",chardata"`
}

type xmlMapStrEntry struct {
	XMLName xml.Name
	Value   string `xml:",chardata"`
}

func (x Xml) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if len(x) == 0 {
		return nil
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	for k, v := range x {
		if err := e.Encode(xmlMapEntry{XMLName: xml.Name{Local: k}, Value: v}); err != nil {
			return err
		}
	}

	return e.EncodeToken(start.End())
}

func (x *Xml) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	*x = Xml{}
	for {
		var e xmlMapStrEntry

		err := d.Decode(&e)
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		(*x)[e.XMLName.Local] = e.Value
	}

	return nil
}

