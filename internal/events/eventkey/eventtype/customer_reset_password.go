package eventtype

type CustomerResetPassword struct {
	*Required
	Code string
}
