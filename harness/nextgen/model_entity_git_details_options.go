package nextgen

import "github.com/antihax/optional"

// EntityGitDetailsOptions used for translating between the different API input options
type EntityGitDetailsOptions struct {
	Branch         optional.String
	RepoIdentifier optional.String
	RootFolder     optional.String
	FilePath       optional.String
	CommitMsg      optional.String
	IsNewBranch    optional.Bool
	BaseBranch     optional.String
	ObjectId       optional.String
}
