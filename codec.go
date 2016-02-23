/**
 * Created by elvizlai on 2016/2/17 15:28
 * Copyright © PubCloud
 */

package AliMNS

import (
	"encoding/base64"
	"encoding/xml"
)

type base64Bytes []byte

func (b base64Bytes) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(base64.StdEncoding.EncodeToString(b), start)
}

func (b *base64Bytes) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	var content string
	if e := d.DecodeElement(&content, &start); e != nil {
		return e
	}

	if buf, e := base64.StdEncoding.DecodeString(content); e != nil {
		return e
	} else {
		*b = base64Bytes(buf)
	}

	return nil
}
