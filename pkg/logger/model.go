package logger

type LoggerRequestData struct {
	Method  string            `json:"method"`
	Path    string            `json:"path"`
	Header  map[string]string `json:"header"`
	Body    []byte            `json:"body"`
	Network NetworkInfo       `json:"network"`
}

type NetworkInfo struct {
	ServerIP string `json:"server_ip"`
	ClientIP string `json:"client_ip"`
}
