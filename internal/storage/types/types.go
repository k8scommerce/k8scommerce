package types

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type ConfigKey string
type HandlerKey string

const (
	ConfigCtx  ConfigKey  = "config"
	HandlerCtx HandlerKey = "handler"
)

type ImageSize struct {
	Tag string `json:"tag"`
	URL string `json:"url"`
}
type ImageSizes struct {
	Sizes []ImageSize `json:"sizes"`
}

func (s *ImageSizes) Marshal() ([]byte, error) {
	bytes, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (s *ImageSizes) Unmarshal(data []byte) (*ImageSizes, error) {
	if err := json.Unmarshal(data, s); err != nil {
		return nil, err
	}
	return s, nil
}

func (i ImageSizes) Value() (driver.Value, error) {
	return json.Marshal(i)
}

func (i *ImageSizes) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &i)
}
