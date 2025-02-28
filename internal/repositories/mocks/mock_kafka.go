package mocks

// MockKafkaProducer is a mock implementation of the KafkaProducer interface
type MockKafkaProducer struct {
	PublishFunc func(event interface{})
	CloseFunc   func()
}

func (m *MockKafkaProducer) Publish(event interface{}) {
	m.PublishFunc(event)
}

func (m *MockKafkaProducer) Close() {
	m.CloseFunc()
}
