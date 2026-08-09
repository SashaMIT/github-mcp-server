package github

import (
	"testing"

	"github.com/google/go-github/v89/github"
)

func TestSanitizeGitHubIssueTextFields_StripsInvisibleTags(t *testing.T) {
	issue := &github.Issue{
		Title: github.Ptr("Hello\U000E0001World"),
		Body:  github.Ptr("Body\U000E0020Text\U000E007FEnd"),
	}

	sanitizeGitHubIssueTextFields(issue)

	if got := issue.GetTitle(); got != "HelloWorld" {
		t.Fatalf("title = %q, want %q", got, "HelloWorld")
	}
	if got := issue.GetBody(); got != "BodyTextEnd" {
		t.Fatalf("body = %q, want %q", got, "BodyTextEnd")
	}
}

func TestSanitizeGitHubIssueTextFields_NilSafe(t *testing.T) {
	sanitizeGitHubIssueTextFields(nil)
	sanitizeGitHubIssueTextFields(&github.Issue{})
}
