package bridge

import (
	"net/http"
	"strings"
	"testing"
)

const dockerHubInputString = `
{
  "callback_url": "https://registry.hub.docker.com/u/svendowideit/testhook/hook/2141b5bi5i5b02bec211i4eeih0242eg11000a/",
  "push_data": {
    "images": [
      "27d47432a69bca5f2700e4dff7de0388ed65f9d3fb1ec645e2bc24c223dc1cc3",
      "51a9c7c1f8bb2fa19bcd09789a34e63f35abb80044bc10196e304f6634cc582c",
      "..."
    ],
    "pushed_at": 1.417566161e+09,
    "pusher": "trustedbuilder",
    "tag": "latest"
  },
  "repository": {
    "comment_count": 0,
    "date_created": 1.417494799e+09,
    "description": "",
    "dockerfile": "FROM scratch\nRUN sleep 10",
    "full_description": "Docker Hub based automated build from a GitHub repo",
    "is_official": false,
    "is_private": true,
    "is_trusted": true,
    "name": "testhook",
    "namespace": "svendowideit",
    "owner": "svendowideit",
    "repo_name": "svendowideit/testhook",
    "repo_url": "https://registry.hub.docker.com/u/svendowideit/testhook/",
    "star_count": 0,
    "status": "Active"
  }
}
`

func TestDockerHubParseHTTPRequest(t *testing.T) {
	req, err := http.NewRequest("POST", "/webhook/example-bridge", strings.NewReader(dockerHubInputString))
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	if err != nil {
		t.Error(err)
	}

	var payload DockerHubInput
	err = payload.ParseHTTPRequest(req)
	if err != nil {
		t.Error(err)
		return
	}

	if payload.CallbackURL != "https://registry.hub.docker.com/u/svendowideit/testhook/hook/2141b5bi5i5b02bec211i4eeih0242eg11000a/" ||
		payload.PushData.Images[0] != "27d47432a69bca5f2700e4dff7de0388ed65f9d3fb1ec645e2bc24c223dc1cc3" {
		t.Fail()
	}
}