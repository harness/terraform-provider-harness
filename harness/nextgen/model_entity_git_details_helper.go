package nextgen

import (
	"github.com/jinzhu/copier"
)

func (d *EntityGitDetailsOptions) ToPipelinesApiPostPipelineOpts() *PipelinesApiPostPipelineOpts {
	opts := &PipelinesApiPostPipelineOpts{}
	err := copier.Copy(opts, d)
	if err != nil {
		return nil
	}
	return opts
}

func (d *EntityGitDetailsOptions) ToPipelinesApiUpdatePipelineV2Opts() *PipelinesApiUpdatePipelineV2Opts {
	opts := &PipelinesApiUpdatePipelineV2Opts{}
	err := copier.Copy(opts, d)
	if err != nil {
		return nil
	}
	return opts
}

func (d *EntityGitDetailsOptions) ToPipelinesApiGetPipelineOpts() *PipelinesApiGetPipelineOpts {
	opts := &PipelinesApiGetPipelineOpts{}
	err := copier.Copy(opts, d)
	if err != nil {
		return nil
	}
	return opts
}

func (d *EntityGitDetailsOptions) ToPipelinesApiDeletePipelineOpts() *PipelinesApiDeletePipelineOpts {
	opts := &PipelinesApiDeletePipelineOpts{}
	err := copier.Copy(opts, d)
	opts.LastObjectId = d.ObjectId
	if err != nil {
		return nil
	}
	return opts
}

func (d *PipelineEntityGitDetails) ToEntityGitDetails() *EntityGitDetails {
	gitEntity := &EntityGitDetails{}
	err := copier.Copy(gitEntity, d)
	if err != nil {
		return nil
	}
	return gitEntity
}

func (d *EntityGitDetails) IsEmpty() bool {
	return d.ObjectId == "" &&
		d.Branch == "" &&
		d.RepoIdentifier == "" &&
		d.RootFolder == "" &&
		d.FilePath == "" &&
		d.RepoName == ""
}
