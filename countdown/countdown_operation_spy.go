package countdown

type CountdownOperationSpy struct {
	Calls []string
}

func (c *CountdownOperationSpy) Sleep() {
	c.Calls = append(c.Calls, "sleep")
}

func (c *CountdownOperationSpy) Write(p []byte) (n int, err error) {
	c.Calls = append(c.Calls, "write")
	return
}
