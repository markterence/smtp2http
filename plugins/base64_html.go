package plugins

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io/ioutil"
)

func StringToBase64(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

// CompressHTMLBody compresses the input string with gzip and encodes it in base64.
// Returns the base64-encoded compressed string, or an empty string on error.
func CompressHTMLBody(s string) string {
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)
	_, err := gz.Write([]byte(s))
	if err != nil {
		return ""
	}
	gz.Close()
	return base64.StdEncoding.EncodeToString(buf.Bytes())
}

// Base64AndGunzip decodes a base64 string and decompresses it with gzip.
func Base64AndGunzip(encoded string) ([]byte, error) {
	compressed, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return nil, err
	}
	r, err := gzip.NewReader(bytes.NewReader(compressed))
	if err != nil {
		return nil, err
	}
	defer r.Close()
	return ioutil.ReadAll(r)
}
