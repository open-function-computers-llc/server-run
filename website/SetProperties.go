package website

func (ws *Site) SetUptimeURI(uri string) error {
	ws.UptimeURI = uri
	return ws.saveStateFile()
}

func (ws *Site) SetLocked(locked bool) error {
	ws.IsLocked = locked
	return ws.saveStateFile()
}
