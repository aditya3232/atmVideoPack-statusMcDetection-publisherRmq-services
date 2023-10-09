package publisher_status_mc_detection

type PublisherStatusMcDetectionFormatter struct {
	Tid           string `json:"tid"`
	DateTime      string `json:"date_time"`
	StatusSignal  string `json:"status_signal"`
	StatusStorage string `json:"status_storage"`
	StatusRam     string `json:"status_ram"`
	StatusCpu     string `json:"status_cpu"`
}

// data ditampilkan dari input, bukan dari entity yg sudah tercreate
func FormatPublisherStatusMcDetection(statusMcDetection StatusMcDetection) PublisherStatusMcDetectionFormatter {

	formatter := PublisherStatusMcDetectionFormatter(statusMcDetection)

	// var formatter = PublisherStatusMcDetectionFormatter{
	// 	Tid:           statusMcDetection.Tid,
	// 	DateTime:      statusMcDetection.DateTime,
	// 	StatusSignal:  statusMcDetection.StatusSignal,
	// 	StatusStorage: statusMcDetection.StatusStorage,
	// 	StatusRam:     statusMcDetection.StatusRam,
	// 	StatusCpu:     statusMcDetection.StatusCpu,
	// }

	return formatter
}
