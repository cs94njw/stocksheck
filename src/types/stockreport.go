package types

type StockReport struct {
	// Information for each stock
	stocks []StockAdvice
	// Overall advice
	generalAdvice GeneralAdvice
}

type StockAdvice struct {
	name               string
	profitEarningRatio float32
	currentPrice       float32
}

type GeneralAdvice struct {
	tenLowestPERatioStocks []StockAdvice
}
