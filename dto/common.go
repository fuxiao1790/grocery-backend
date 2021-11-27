package dto

import "encoding/json"

type Err struct {
	Error error
}

func (e Err) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.Error.Error())
}
