package gollections

func (c *Collection) Map(fn func(interface{}) interface{}) *Collection {
	var mapped []interface{}
	for _, item := range c.items {
		mapped = append(mapped, fn(item))
	}
	return New(mapped...)
}
