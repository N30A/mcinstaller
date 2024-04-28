package servers

const (
	versionManifestURL = "https://launchermeta.mojang.com/mc/game/version_manifest.json"
)

type Vanilla struct{}

func (v *Vanilla) Versions() ([]string, error) {
	return nil, nil
}

func (v *Vanilla) DownloadURL(version string) (string, error) {
	return "", nil
}
