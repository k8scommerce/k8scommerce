package eventkey

import (
	"encoding/json"

	"github.com/k8scommerce/k8scommerce/services/rpc/cart/pb/cart"
)

// cart

/////////////////
// CartProductAddedKey
/////////////////

func (s *CartProductAddedKey) Marshal(obj *cart.Cart) ([]byte, error) {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (s *CartProductAddedKey) Unmarshal(data []byte) (obj *cart.Cart, err error) {
	obj = &cart.Cart{}
	if err = json.Unmarshal(data, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *CartProductAddedKey) AsKey() EventKey {
	return EventKey(*s)
}

/////////////////
// CartProductUpdatedKey
/////////////////

func (s *CartProductUpdatedKey) Marshal(obj *cart.Cart) ([]byte, error) {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (s *CartProductUpdatedKey) Unmarshal(data []byte) (obj *cart.Cart, err error) {
	obj = &cart.Cart{}
	if err = json.Unmarshal(data, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *CartProductUpdatedKey) AsKey() EventKey {
	return EventKey(*s)
}

/////////////////
// CartProductDeletedKey
/////////////////

func (s *CartProductDeletedKey) Marshal(obj *cart.Cart) ([]byte, error) {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (s *CartProductDeletedKey) Unmarshal(data []byte) (obj *cart.Cart, err error) {
	obj = &cart.Cart{}
	if err = json.Unmarshal(data, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *CartProductDeletedKey) AsKey() EventKey {
	return EventKey(*s)
}
