package nextgen

type FileUsage string

var FileUsages = struct {
	ManifestFile FileUsage
	Config       FileUsage
	Script       FileUsage
}{
	ManifestFile: "MANIFEST_FILE",
	Config:       "CONFIG",
	Script:       "SCRIPT",
}

var FileUsageValues = []string{
	FileUsages.ManifestFile.String(), ",",
	FileUsages.Config.String(), ",",
	FileUsages.Script.String(),
}

func (e FileUsage) String() string {
	return string(e)
}
