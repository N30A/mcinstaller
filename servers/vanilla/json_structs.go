package servers

type versionManifest struct {
	Versions []struct {
		ID  string `json:"id"`
		URL string `json:"url"`
	} `json:"versions"`
}

type versionPackage struct {
	ID        string `json:"id"`
	Downloads struct {
		Server struct {
			URL string `json:"url"`
		} `json:"server"`
	} `json:"downloads"`
}

func (m *versionManifest) indexOf(version string) int {
	for i, v := range m.Versions {
		if v.ID == version {
			return i
		}
	}
	return -1
}

func (m *versionManifest) toStringSlice() []string {
	versions := make([]string, len(m.Versions))

	for i := 0; i < len(m.Versions); i++ {
		versions[i] = m.Versions[i].ID
	}

	return versions
}
