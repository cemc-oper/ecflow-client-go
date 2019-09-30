package ecflow_client

import "time"

type EcflowClient struct {
	ServerHost    string
	ServerPort    string
	CollectedTime time.Time
	wrapper       EcflowClientWrapper
}

type StatusRecord struct {
	Path   string `json:"path"`
	Status string `json:"status"`
}

func CreateEcflowClient(host string, port string) *EcflowClient {
	client := &EcflowClient{
		ServerHost: host,
		ServerPort: port,
	}
	client.wrapper = NewEcflowClientWrapper(client.ServerHost, client.ServerPort)
	return client
}

func (c *EcflowClient) SetConnectTimeout(timeout int) {
	c.wrapper.SetConnectTimeout(timeout)
}

func (c *EcflowClient) Sync() int {
	if c.wrapper == nil {
		c.wrapper = NewEcflowClientWrapper(c.ServerHost, c.ServerPort)
	}
	errorCode := c.wrapper.Sync()
	c.CollectedTime = time.Now()
	return errorCode
}

func (c *EcflowClient) StatusRecords() []StatusRecord {
	var returnRecords []StatusRecord
	records := c.wrapper.StatusRecords()
	count := int(records.Size())
	for i := 0; i < count; i++ {
		record := records.Get(i)
		returnRecords = append(returnRecords, StatusRecord{
			Path:   record.GetPath_(),
			Status: record.GetStatus_(),
		})
	}
	return returnRecords
}

func (c *EcflowClient) Close() {
	DeleteEcflowClientWrapper(c.wrapper)
	c.wrapper = nil
}
