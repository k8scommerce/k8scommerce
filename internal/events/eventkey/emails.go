package eventkey

import (
	"encoding/json"
	"k8scommerce/internal/events/eventkey/eventtype"
)

// customer

/////////////////
// AdminCancelledOrderKey
/////////////////

func (s *AdminCancelledOrderKey) Marshal(obj *eventtype.AdminCancelledOrder) ([]byte, error) {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (s *AdminCancelledOrderKey) Unmarshal(data []byte) (obj *eventtype.AdminCancelledOrder, err error) {
	obj = &eventtype.AdminCancelledOrder{}
	if err = json.Unmarshal(data, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *AdminCancelledOrderKey) AsKey() EventKey {
	return EventKey(*s)
}

/////////////////
// AdminFailedOrderKey
/////////////////

func (s *AdminFailedOrderKey) Marshal(obj *eventtype.AdminFailedOrder) ([]byte, error) {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (s *AdminFailedOrderKey) Unmarshal(data []byte) (obj *eventtype.AdminFailedOrder, err error) {
	obj = &eventtype.AdminFailedOrder{}
	if err = json.Unmarshal(data, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *AdminFailedOrderKey) AsKey() EventKey {
	return EventKey(*s)
}

/////////////////
// AdminNewOrderKey
/////////////////

func (s *AdminNewOrderKey) Marshal(obj *eventtype.AdminNewOrder) ([]byte, error) {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (s *AdminNewOrderKey) Unmarshal(data []byte) (obj *eventtype.AdminNewOrder, err error) {
	obj = &eventtype.AdminNewOrder{}
	if err = json.Unmarshal(data, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *AdminNewOrderKey) AsKey() EventKey {
	return EventKey(*s)
}

/////////////////
// CustomerAccountConfirmedEmailKey
/////////////////

func (s *CustomerAccountConfirmedEmailKey) Marshal(obj *eventtype.CustomerAccountConfirmedEmail) ([]byte, error) {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (s *CustomerAccountConfirmedEmailKey) Unmarshal(data []byte) (obj *eventtype.CustomerAccountConfirmedEmail, err error) {
	obj = &eventtype.CustomerAccountConfirmedEmail{}
	if err = json.Unmarshal(data, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *CustomerAccountConfirmedEmailKey) AsKey() EventKey {
	return EventKey(*s)
}

/////////////////
// CustomerCompletedOrderKey
/////////////////

func (s *CustomerCompletedOrderKey) Marshal(obj *eventtype.CustomerCompletedOrder) ([]byte, error) {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (s *CustomerCompletedOrderKey) Unmarshal(data []byte) (obj *eventtype.CustomerCompletedOrder, err error) {
	obj = &eventtype.CustomerCompletedOrder{}
	if err = json.Unmarshal(data, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *CustomerCompletedOrderKey) AsKey() EventKey {
	return EventKey(*s)
}

/////////////////
// CustomerConfirmationEmailKey
/////////////////

func (s *CustomerConfirmationEmailKey) Marshal(obj *eventtype.CustomerConfirmationEmail) ([]byte, error) {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (s *CustomerConfirmationEmailKey) Unmarshal(data []byte) (obj *eventtype.CustomerConfirmationEmail, err error) {
	obj = &eventtype.CustomerConfirmationEmail{}
	if err = json.Unmarshal(data, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *CustomerConfirmationEmailKey) AsKey() EventKey {
	return EventKey(*s)
}

/////////////////
// CustomerNewAccountKey
/////////////////

func (s *CustomerNewAccountKey) Marshal(obj *eventtype.CustomerNewAccount) ([]byte, error) {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (s *CustomerNewAccountKey) Unmarshal(data []byte) (obj *eventtype.CustomerNewAccount, err error) {
	obj = &eventtype.CustomerNewAccount{}
	if err = json.Unmarshal(data, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *CustomerNewAccountKey) AsKey() EventKey {
	return EventKey(*s)
}

/////////////////
// CustomerNoteKey
/////////////////

func (s *CustomerNoteKey) Marshal(obj *eventtype.CustomerNote) ([]byte, error) {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (s *CustomerNoteKey) Unmarshal(data []byte) (obj *eventtype.CustomerNote, err error) {
	obj = &eventtype.CustomerNote{}
	if err = json.Unmarshal(data, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *CustomerNoteKey) AsKey() EventKey {
	return EventKey(*s)
}

/////////////////
// CustomerOnHoldOrderKey
/////////////////

func (s *CustomerOnHoldOrderKey) Marshal(obj *eventtype.CustomerOnHoldOrder) ([]byte, error) {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (s *CustomerOnHoldOrderKey) Unmarshal(data []byte) (obj *eventtype.CustomerOnHoldOrder, err error) {
	obj = &eventtype.CustomerOnHoldOrder{}
	if err = json.Unmarshal(data, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *CustomerOnHoldOrderKey) AsKey() EventKey {
	return EventKey(*s)
}

/////////////////
// CustomerPasswordChangedKey
/////////////////

func (s *CustomerPasswordChangedKey) Marshal(obj *eventtype.CustomerPasswordChanged) ([]byte, error) {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (s *CustomerPasswordChangedKey) Unmarshal(data []byte) (obj *eventtype.CustomerPasswordChanged, err error) {
	obj = &eventtype.CustomerPasswordChanged{}
	if err = json.Unmarshal(data, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *CustomerPasswordChangedKey) AsKey() EventKey {
	return EventKey(*s)
}

/////////////////
// CustomerProcessingOrderKey
/////////////////

func (s *CustomerProcessingOrderKey) Marshal(obj *eventtype.CustomerProcessingOrder) ([]byte, error) {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (s *CustomerProcessingOrderKey) Unmarshal(data []byte) (obj *eventtype.CustomerProcessingOrder, err error) {
	obj = &eventtype.CustomerProcessingOrder{}
	if err = json.Unmarshal(data, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *CustomerProcessingOrderKey) AsKey() EventKey {
	return EventKey(*s)
}

/////////////////
// CustomerRefundedOrderKey
/////////////////

func (s *CustomerRefundedOrderKey) Marshal(obj *eventtype.CustomerRefundedOrder) ([]byte, error) {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (s *CustomerRefundedOrderKey) Unmarshal(data []byte) (obj *eventtype.CustomerRefundedOrder, err error) {
	obj = &eventtype.CustomerRefundedOrder{}
	if err = json.Unmarshal(data, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *CustomerRefundedOrderKey) AsKey() EventKey {
	return EventKey(*s)
}

/////////////////
// CustomerResetPasswordKey
/////////////////

func (s *CustomerResetPasswordKey) Marshal(obj *eventtype.CustomerResetPassword) ([]byte, error) {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (s *CustomerResetPasswordKey) Unmarshal(data []byte) (obj *eventtype.CustomerResetPassword, err error) {
	obj = &eventtype.CustomerResetPassword{}
	if err = json.Unmarshal(data, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *CustomerResetPasswordKey) AsKey() EventKey {
	return EventKey(*s)
}

/////////////////
// CustomerSaleKey
/////////////////

func (s *CustomerSaleKey) Marshal(obj *eventtype.CustomerSale) ([]byte, error) {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (s *CustomerSaleKey) Unmarshal(data []byte) (obj *eventtype.CustomerSale, err error) {
	obj = &eventtype.CustomerSale{}
	if err = json.Unmarshal(data, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *CustomerSaleKey) AsKey() EventKey {
	return EventKey(*s)
}
