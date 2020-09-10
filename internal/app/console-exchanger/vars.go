package console_exchanger

var (
	usageAnswer = "Usage:\n" +
		"\t./currency <amount:float> <src_symbol:string> <dst_symbol:string>\n" +
		"Example:\n" +
		"\t./currency 100 USD RUB"
	
	wrongArgs = "Wrong arguments! Use -h to get information."
	
	rate = []string{"CAD", "HKD", "ISK", "PHP", "DKK", "HUF", "CZK", "GBP", "RON",
		"SEK", "IDR", "INR", "BRL", "RUB", "HRK", "JPY", "THB", "CHF", "EUR", "MYR",
		"BGN", "TRY", "CNY", "NOK", "NZD", "ZAR", "USD", "MXN", "SGD", "AUD", "ILS",
		"KRW", "PLN"}
)
