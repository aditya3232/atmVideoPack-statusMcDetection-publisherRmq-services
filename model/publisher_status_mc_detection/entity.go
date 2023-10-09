package publisher_status_mc_detection

// disini entity buat return
type StatusMcDetection struct {
	Tid           string `json:"tid"`
	DateTime      string `json:"date_time"`
	StatusSignal  string `json:"status_signal"`
	StatusStorage string `json:"status_storage"`
	StatusRam     string `json:"status_ram"`
	StatusCpu     string `json:"status_cpu"`
}
