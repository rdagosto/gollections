package gollections

func (c *Collection) Filter(fn func(interface{}) bool) *Collection {
	var filtered []interface{}
	for _, item := range c.items {
		if fn(item) {
			filtered = append(filtered, item)
		}
	}
	return New(filtered...)
}
