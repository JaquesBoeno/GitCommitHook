package commitMessage

import (
	"testing"
)

func TestCommitMessageBuilder(t *testing.T) {
	got := CommitMessageBuilder("<type>(<scope>): <subject>\n\n<desc>\n", []Value{
		{Id: "type", Value: "feat"},
		{Id: "scope", Value: "front-end"},
		{Id: "subject", Value: "user card added"},
		{Id: "desc", Value: "some description"},
	})

	want := "feat(front-end): user card added\n\nsome description\n"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
