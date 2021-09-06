package client

func (c *srv) OnDisconnect(handler func(error)) {
	c.onDisconnect = handler
}

func (c *srv) OnError(handler func(error)) {
	c.onError = handler
}
