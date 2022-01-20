package website

func (ws *Site) SetUptimeURI(uri string) error {
	ws.UptimeURI = uri
	return ws.saveStateFile()
}
