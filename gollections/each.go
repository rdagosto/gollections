package gollections

func (c *Collection) Each(fn func(interface{})) {
	for _, item := range c.items {
		fn(item)
	}
}
