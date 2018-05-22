package webhook

import (
	"reflect"
	"testing"
)

const JSON_STR = `
{
  "push_data": {
    "pushed_at": 1526689682,
    "images": [],
    "tag": "alpine",
    "pusher": "joaofnfernandes"
  },
  "callback_url": "https://registry.hub.docker.com/u/joaofnfernandes/test/hook/2c2e331dd21a14e5fd0ddgagce2eb5d0h/",
  "repository": {
    "status": "Active",
    "description": "",
    "is_trusted": false,
    "full_description": "",
    "repo_url": "https://hub.docker.com/r/joaofnfernandes/test",
    "owner": "joaofnfernandes",
    "is_official": false,
    "is_private": false,
    "name": "test",
    "namespace": "joaofnfernandes",
    "star_count": 0,
    "comment_count": 0,
    "date_created": 1492469385,
    "repo_name": "joaofnfernandes/test"
  }
}
`

var (
	expected = Webhook{
		CallbackUrl: "https://registry.hub.docker.com/u/joaofnfernandes/test/hook/2c2e331dd21a14e5fd0ddgagce2eb5d0h/",
		PushData: PushData{
			PushedAt: 1526689682,
			Images:   []string{},
			Tag:      "alpine",
			Pusher:   "joaofnfernandes",
		},
		Repository: Repository{
			Status:          "Active",
			Description:     "",
			IsTrusted:       false,
			FullDescription: "",
			RepoUrl:         "https://hub.docker.com/r/joaofnfernandes/test",
			Owner:           "joaofnfernandes",
			IsOfficial:      false,
			IsPrivate:       false,
			Name:            "test",
			Namespace:       "joaofnfernandes",
			StarCount:       0,
			CommentCount:    0,
			DateCreated:     1492469385,
			RepoName:        "joaofnfernandes/test",
		},
	}
)

func TestUnmarshal(t *testing.T) {
	got := Unmarshal(JSON_STR)
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("Failed to unmarshal webhook from json. Got:\n%v\nexpected:\n%v", got, expected)
	}
}

func TestIsDefault(t *testing.T) {
	defaultWebhook := Webhook{}
	nonDefaultWebhook := Webhook{CallbackUrl: "http://example.org"}

	if !defaultWebhook.isDefault() {
		t.Fatalf("IsDefault says webhook is not default, when it is")
	}
	if nonDefaultWebhook.isDefault() {
		t.Fatalf("IsDefault says webhook is default, when it isn't")
	}

}

func TestIsValid(t *testing.T) {
	validWebhook := Webhook{
		Repository: Repository{
			Namespace: "joaofnfernandes",
			Name:      "test",
		},
		PushData: PushData{
			Tag: "latest",
		},
	}
	if !validWebhook.IsValid() {
		t.Fatalf("IsValid says webhook is not valid, when it is")
	}
}
