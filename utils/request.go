package utils

type RequestParams struct {
	value map[string]interface{}
}

func NewRequestParams() *RequestParams {
	return &RequestParams{
		value: make(map[string]interface{}),
	}
}

func (p *RequestParams) Values() map[string]interface{} {
	return p.value
}

func (p *RequestParams) SetString(key string, val string) *RequestParams {
	p.value[key] = val
	return p
}

func (p *RequestParams) SetInt(key string, val int) *RequestParams {
	p.value[key] = val
	return p
}

func (p *RequestParams) SetInt32(key string, val int32) *RequestParams {
	p.value[key] = val
	return p
}

func (p *RequestParams) SetInt64(key string, val int64) *RequestParams {
	p.value[key] = val
	return p
}

func (p *RequestParams) GetString(key string) (val string) {
	val, _ = p.value[key].(string)
	return val
}

func (p *RequestParams) Get(key string) interface{} {
	return p.value[key]
}

func (p *RequestParams) GetInt(key string) (val int) {
	val, _ = p.value[key].(int)
	return val
}

func (p *RequestParams) GetInt32(key string) (val int32) {
	val, _ = p.value[key].(int32)
	return val
}

func (p *RequestParams) GetInt64(key string) (val int64) {
	val, _ = p.value[key].(int64)
	return val
}

