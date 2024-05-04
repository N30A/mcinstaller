package servers

type versionManifest struct {
	Game []struct {
		Version string `json:"version"`
	} `json:"game"`
	Loader []struct {
		Version string `json:"version"`
	} `json:"loader"`
	Installer []struct {
		Version string `json:"version"`
	} `json:"installer"`
}

func (m *versionManifest) toStringSlice() []string {
	versions := make([]string, len(m.Game))

	for i := 0; i < len(m.Game); i++ {
		versions[i] = m.Game[i].Version
	}

	return versions
}
