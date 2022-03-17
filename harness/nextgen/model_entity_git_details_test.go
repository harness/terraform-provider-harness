package nextgen_test

import (
	"testing"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/jinzhu/copier"
	"github.com/stretchr/testify/require"
)

func TestSerialization(t *testing.T) {

	gitEntity := &nextgen.EntityGitDetails{
		ObjectId:       "object_id",
		Branch:         "branch",
		RepoIdentifier: "repo_id",
		RootFolder:     "root_folder",
		FilePath:       "file_path",
		RepoName:       "repo_name",
	}

	opts := &nextgen.PipelinesApiUpdatePipelineOpts{}

	err := copier.Copy(opts, gitEntity)
	require.NoError(t, err)

}
