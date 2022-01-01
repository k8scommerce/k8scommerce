package utils

import "encoding/json"

func TransformObj(from interface{}, to interface{}) error {
	byteFrom, err := json.Marshal(&from)
	if err != nil {
		return err
	}
	err = json.Unmarshal(byteFrom, to)
	if err != nil {
		return err
	}
	return nil
}
