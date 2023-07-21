package nextgen

type FileUsage string

var FileUsages = struct {
	ManifestFile FileUsage
	Config       FileUsage
	Script       FileUsage
}{
	ManifestFile: "ManifestFile",
	Config:       "Config",
	Script:       "Script",
}

var FileUsageValues = []string{
	FileUsages.ManifestFile.String(),
	FileUsages.Config.String(),
	FileUsages.Script.String(),
}

func (e FileUsage) String() string {
	return string(e)
}
