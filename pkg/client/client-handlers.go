package client

func (c *srv) OnDisconnect(handler func(Topic, error)) {
	c.onDisconnect = handler
}

func (c *srv) OnError(handler func(Topic, error)) {
	c.onError = handler
}
