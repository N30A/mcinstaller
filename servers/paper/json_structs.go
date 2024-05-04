package servers

type versionManifest struct {
	Versions []string `json:"versions"`
}

type versionBuilds struct {
	Builds []int `json:"builds"`
}
