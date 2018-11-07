package main

import (
	"log"
	"webhook-bridge"
)

func main() {
	bridge.LoadConfigurationFromFile("./examples/config.yaml")
	bridge.RegisterConverter("DockerHubToSlack", func(input interface{}) (interface{}, error) {
		hubInput := input.(*bridge.DockerHubInput)
		output := &bridge.SlackOutput{
			Text: hubInput.Repository.RepoName + " build successfully.",
			Attachments: []bridge.SlackOutputAttachment{
				bridge.SlackOutputAttachment{
					Fields: []bridge.SlackOutputAttachmentField{
						bridge.SlackOutputAttachmentField{Title: "Tag", Value: hubInput.PushData.Tag, Short: true},
						bridge.SlackOutputAttachmentField{Title: "Status", Value: hubInput.Repository.Status, Short: true},
					},
				},
			},
		}
		return output, nil
	})

	err := bridge.Start()
	if err != nil {
		log.Fatal(err)
		return
	}
}
