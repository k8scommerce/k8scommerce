package eventkey

import (
	"encoding/json"
)

// debug/testing

/////////////////
// DebugTestingKey
/////////////////

func (s *DebugTestingKey) Marshal(obj interface{}) ([]byte, error) {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (s *DebugTestingKey) Unmarshal(data []byte) (obj interface{}, err error) {
	if err = json.Unmarshal(data, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *DebugTestingKey) AsKey() EventKey {
	return EventKey(*s)
}
