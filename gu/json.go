package gu

import (
	"bytes"
	"encoding/json"
)

// ToJSONString convert data to json string
func ToJSONString(data interface{}) string {
	if res, err := json.Marshal(data); err != nil {
		return "{}" // output empty json object string if got an error
	} else {
		return string(res)
	}
}

// ToJSONStringEscapeHTML  convert data to json string escape html
func ToJSONStringEscapeHTML(data interface{}) string {
	bf := bytes.NewBufferString("")
	jsonEncoder := json.NewEncoder(bf)
	jsonEncoder.SetEscapeHTML(false)
	err := jsonEncoder.Encode(data)
	if err != nil {
		return ""
	}
	return bf.String()
}

// IsJSON check if the data is json
func IsJSON(str string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}
