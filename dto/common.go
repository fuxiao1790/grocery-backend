package dto

import "encoding/json"

type Err struct {
	Error error
}

func (e Err) MarshalJSON() ([]byte, error) {
	if e.Error == nil {
		return json.Marshal(nil)
	}
	return json.Marshal(e.Error.Error())
}
