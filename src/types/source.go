package types

type source interface {
	send(StockReport) bool
}
