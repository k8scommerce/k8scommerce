package eventtype

type CustomerConfirmationEmail struct {
	*Required
	Code string
}
