package types

type destination interface {
	send(StockReport) bool
}
