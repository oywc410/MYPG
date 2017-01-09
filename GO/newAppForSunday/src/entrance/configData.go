package entrance

type ConfigData struct {
	HttpAddr   string
	ChannelMux int
	OutChangelMux int
	NsqAddrs   []string
}
