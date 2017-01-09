package entrance

type message struct {
	message chan string
	outMessage chan string
	configData ConfigData
}

func NewMessage(configData ConfigData) *message {
	return &message{
		message: make(chan string, configData.ChannelMux),
		outMessage: make(chan string, configData.OutChangelMux),
		configData: configData,
	}
}

func (m *message)AddMessage(message string) {
	m.message <- message
}

func (m *message)addOutMessage(meesage string) {
	m.outMessage <- meesage
}

func (m *message)getMessageChan() <- chan string {
	return m.message
}

func (m *message)getOutMessage() <- chan string {
	return m.outMessage
}

func (m *message)PushOutMessageFile() error {

	return nil
}

func (m *message)Close() error {
	close(m.message)
	return nil
}