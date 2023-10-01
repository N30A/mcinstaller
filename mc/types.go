package mc

func SetDefault(name string) MCServer {
	st := MCServer{}
	st.Name = name

	return st
}

type MCServer struct {
	Name     string
	Versions []MCVersion
}

type MCVersion struct {
	ID  string
	URL string
}

type VanillaManifest struct {
	Versions []struct {
		URL string `json:"url"`
	} `json:"versions"`
}

type VanillaData struct {
	ID        string `json:"id"`
	Downloads struct {
		Server struct {
			URL string `json:"url"`
		} `json:"server"`
	} `json:"downloads"`
}

type PaperManifest struct {
	Versions []string `json:"versions"`
}

type PaperData struct {
	Builds []int `json:"builds"`
}

type FabricManifest struct {
	Versions []struct {
		ID string `json:"version"`
	} `json:"game"`
	Loaders []struct {
		Stable  bool   `json:"stable"`
		Version string `json:"version"`
	} `json:"loader"`
	Installers []struct {
		Stable  bool   `json:"stable"`
		Version string `json:"version"`
	} `json:"installer"`
}
