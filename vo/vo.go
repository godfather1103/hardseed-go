package vo

type DownloadInfo struct {
	Url string
	Host string
	Urlprefix string
	Proxy string
	PageStart int64
	PageEnd int64
	HateList []string
	LikeList []string
	Path string
	TimeoutDownloadPic int64
}
