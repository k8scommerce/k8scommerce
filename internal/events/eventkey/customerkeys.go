package eventkey

import (
	"encoding/json"

	"github.com/k8scommerce/k8scommerce/services/rpc/customer/pb/customer"
)

// customer

/////////////////
// CustomerCreatedKey
/////////////////

func (s *CustomerCreatedKey) Marshal(obj *customer.Customer) ([]byte, error) {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (s *CustomerCreatedKey) Unmarshal(data []byte) (obj *customer.Customer, err error) {
	obj = &customer.Customer{}
	if err = json.Unmarshal(data, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *CustomerCreatedKey) AsKey() EventKey {
	return EventKey(*s)
}

/////////////////
// CustomerUpdatedKey
/////////////////

func (s *CustomerUpdatedKey) Marshal(obj *customer.Customer) ([]byte, error) {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (s *CustomerUpdatedKey) Unmarshal(data []byte) (obj *customer.Customer, err error) {
	obj = &customer.Customer{}
	if err = json.Unmarshal(data, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *CustomerUpdatedKey) AsKey() EventKey {
	return EventKey(*s)
}

/////////////////
// CustomerForgotPasswordKey
/////////////////

func (s *CustomerForgotPasswordKey) Marshal(obj *customer.Customer) ([]byte, error) {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (s *CustomerForgotPasswordKey) Unmarshal(data []byte) (obj *customer.Customer, err error) {
	obj = &customer.Customer{}
	if err = json.Unmarshal(data, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *CustomerForgotPasswordKey) AsKey() EventKey {
	return EventKey(*s)
}

/////////////////
// CustomerChangedPasswordKey
/////////////////

func (s *CustomerChangedPasswordKey) Marshal(obj *customer.Customer) ([]byte, error) {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (s *CustomerChangedPasswordKey) Unmarshal(data []byte) (obj *customer.Customer, err error) {
	obj = &customer.Customer{}
	if err = json.Unmarshal(data, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *CustomerChangedPasswordKey) AsKey() EventKey {
	return EventKey(*s)
}

/////////////////
// CustomerLoginSuccessKey
/////////////////

func (s *CustomerLoginSuccessKey) Marshal(obj *customer.Customer) ([]byte, error) {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (s *CustomerLoginSuccessKey) Unmarshal(data []byte) (obj *customer.Customer, err error) {
	obj = &customer.Customer{}
	if err = json.Unmarshal(data, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *CustomerLoginSuccessKey) AsKey() EventKey {
	return EventKey(*s)
}

/////////////////
// CustomerLoginFailedKey
/////////////////

func (s *CustomerLoginFailedKey) Marshal(obj *customer.LoginRequest) ([]byte, error) {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (s *CustomerLoginFailedKey) Unmarshal(data []byte) (obj *customer.LoginRequest, err error) {
	obj = &customer.LoginRequest{}
	if err = json.Unmarshal(data, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *CustomerLoginFailedKey) AsKey() EventKey {
	return EventKey(*s)
}

/////////////////
// CustomerAddedBillingAddressKey
/////////////////

func (s *CustomerAddedBillingAddressKey) Marshal(obj *customer.Address) ([]byte, error) {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (s *CustomerAddedBillingAddressKey) Unmarshal(data []byte) (obj *customer.Address, err error) {
	obj = &customer.Address{}
	if err = json.Unmarshal(data, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *CustomerAddedBillingAddressKey) AsKey() EventKey {
	return EventKey(*s)
}

/////////////////
// CustomerUpdatedBillingAddressKey
/////////////////

func (s *CustomerUpdatedBillingAddressKey) Marshal(obj *customer.Address) ([]byte, error) {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (s *CustomerUpdatedBillingAddressKey) Unmarshal(data []byte) (obj *customer.Address, err error) {
	obj = &customer.Address{}
	if err = json.Unmarshal(data, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *CustomerUpdatedBillingAddressKey) AsKey() EventKey {
	return EventKey(*s)
}

/////////////////
// CustomerDeletedBillingAddressKey
/////////////////

func (s *CustomerDeletedBillingAddressKey) Marshal(obj *customer.Address) ([]byte, error) {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (s *CustomerDeletedBillingAddressKey) Unmarshal(data []byte) (obj *customer.Address, err error) {
	obj = &customer.Address{}
	if err = json.Unmarshal(data, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *CustomerDeletedBillingAddressKey) AsKey() EventKey {
	return EventKey(*s)
}

/////////////////
// CustomerAddedShippingAddressKey
/////////////////

func (s *CustomerAddedShippingAddressKey) Marshal(obj *customer.Address) ([]byte, error) {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (s *CustomerAddedShippingAddressKey) Unmarshal(data []byte) (obj *customer.Address, err error) {
	obj = &customer.Address{}
	if err = json.Unmarshal(data, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *CustomerAddedShippingAddressKey) AsKey() EventKey {
	return EventKey(*s)
}

/////////////////
// CustomerUpdatedShippingAddressKey
/////////////////

func (s *CustomerUpdatedShippingAddressKey) Marshal(obj *customer.Address) ([]byte, error) {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (s *CustomerUpdatedShippingAddressKey) Unmarshal(data []byte) (obj *customer.Address, err error) {
	obj = &customer.Address{}
	if err = json.Unmarshal(data, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *CustomerUpdatedShippingAddressKey) AsKey() EventKey {
	return EventKey(*s)
}

/////////////////
// CustomerDeletedShippingAddressKey
/////////////////

func (s *CustomerDeletedShippingAddressKey) Marshal(obj *customer.Address) ([]byte, error) {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (s *CustomerDeletedShippingAddressKey) Unmarshal(data []byte) (obj *customer.Address, err error) {
	obj = &customer.Address{}
	if err = json.Unmarshal(data, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *CustomerDeletedShippingAddressKey) AsKey() EventKey {
	return EventKey(*s)
}
