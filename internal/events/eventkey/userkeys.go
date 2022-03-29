package eventkey

import (
	"encoding/json"

	"github.com/k8scommerce/k8scommerce/services/rpc/user/pb/user"
)

// user

/////////////////
// UserCreatedKey
/////////////////

func (s *UserCreatedKey) Marshal(obj *user.User) ([]byte, error) {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (s *UserCreatedKey) Unmarshal(data []byte) (obj *user.User, err error) {
	if err = json.Unmarshal(data, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *UserCreatedKey) AsKey() EventKey {
	return EventKey(*s)
}

/////////////////
// UserUpdatedKey
/////////////////

func (s *UserUpdatedKey) Marshal(obj *user.User) ([]byte, error) {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (s *UserUpdatedKey) Unmarshal(data []byte) (obj *user.User, err error) {
	if err = json.Unmarshal(data, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *UserUpdatedKey) AsKey() EventKey {
	return EventKey(*s)
}

/////////////////
// UserForgotPasswordKey
/////////////////

func (s *UserForgotPasswordKey) Marshal(obj *user.User) ([]byte, error) {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (s *UserForgotPasswordKey) Unmarshal(data []byte) (obj *user.User, err error) {
	if err = json.Unmarshal(data, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *UserForgotPasswordKey) AsKey() EventKey {
	return EventKey(*s)
}

/////////////////
// UserChangedPasswordKey
/////////////////

func (s *UserChangedPasswordKey) Marshal(obj *user.User) ([]byte, error) {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (s *UserChangedPasswordKey) Unmarshal(data []byte) (obj *user.User, err error) {
	if err = json.Unmarshal(data, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *UserChangedPasswordKey) AsKey() EventKey {
	return EventKey(*s)
}

/////////////////
// UserLoginSuccessKey
/////////////////

func (s *UserLoginSuccessKey) Marshal(obj *user.User) ([]byte, error) {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (s *UserLoginSuccessKey) Unmarshal(data []byte) (obj *user.User, err error) {
	if err = json.Unmarshal(data, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *UserLoginSuccessKey) AsKey() EventKey {
	return EventKey(*s)
}

/////////////////
// UserLoginFailedKey
/////////////////

func (s *UserLoginFailedKey) Marshal(obj *user.LoginRequest) ([]byte, error) {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (s *UserLoginFailedKey) Unmarshal(data []byte) (obj *user.LoginRequest, err error) {
	if err = json.Unmarshal(data, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *UserLoginFailedKey) AsKey() EventKey {
	return EventKey(*s)
}
