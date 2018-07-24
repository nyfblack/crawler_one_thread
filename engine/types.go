package engine

type Request struct {
	Url    string
	Parser func([]byte) ParserResult
}

type ParserResult struct {
	Requests []Request
	Items    []interface{}
}

func NilParser(context []byte) ParserResult {
	return ParserResult{}
}
