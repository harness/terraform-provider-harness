package project

import (
	"errors"
	"testing"
)

func TestIsGitopsProjectReferencedByApplicationsError(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
		err  error
		want bool
	}{
		{
			name: "nil",
			err:  nil,
			want: false,
		},
		{
			name: "unrelated error",
			err:  errors.New("something else"),
			want: false,
		},
		{
			name: "referenced by applications",
			err:  errors.New("rpc error: code = InvalidArgument desc = project is referenced by 1 applications"),
			want: true,
		},
		{
			name: "case insensitive",
			err:  errors.New("Project is referenced by 2 Applications"),
			want: true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			if got := isGitopsProjectReferencedByApplicationsError(tc.err); got != tc.want {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}
