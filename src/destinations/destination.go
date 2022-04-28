package destinations

type destination interface {
	send(StockReport) bool
}
