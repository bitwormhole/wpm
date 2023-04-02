package v3filters

const orderBase = 0

const (
	order0 = orderBase + iota
	orderPrepareX
	orderPrepareAction
	orderFindTemplate
	orderCheckTemplate
	orderPrepareProperties
	orderCLIMaker
	orderCLIRunner
)
