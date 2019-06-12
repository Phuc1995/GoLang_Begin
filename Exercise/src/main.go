package enums

type ResultState uint8

const (
	Ordered ResultState = iota + 1
	Canceled
	Delivered
)

const (
	TAX1 ResultState = iota + 1
	TAX2
	TAX3
)

const (
	VAT = (iota + 1) * 0.01
	TNDN
	TTNDB
)
