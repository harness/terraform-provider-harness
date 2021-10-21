package nextgen

type ClientKeyAlgorithm string

var ClientKeyAlgorithms = struct {
	RSA   ClientKeyAlgorithm
	ECDSA ClientKeyAlgorithm
}{
	RSA:   "RSA",
	ECDSA: "EC",
}

var ClientKeyAlgorithmsSlice = []string{
	ClientKeyAlgorithms.RSA.String(),
	ClientKeyAlgorithms.ECDSA.String(),
}

func (enum ClientKeyAlgorithm) String() string {
	return string(enum)
}
