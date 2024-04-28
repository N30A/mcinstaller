package servers

var SupportedServers = []string{
	"vanilla",
	"paper",
	"fabric",
	"forge",
}

type Server interface {
	Versions() ([]string, error)
	DownloadURL(version string) (string, error)
}
