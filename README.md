# Webhook Bridge

> Receive webhooks on your server, and integrate them with others.

## Usage

### Prerequisites

  - Go 1.X

### Supported Hooks

  - Input
    - Docker Hub
  - Output
    - Discord
    - Slack

### Example 1. Simple way (using docker)

TBD

### Example 2. Normal way

```yaml
### Version of the configration file template.
version: '1'


### Configure server for incoming webhooks.
server:
  # URL path. In the example, path for the bridge named `example-bridge` becomes `/webhook/example-bridge`.
  path_prefix: '/webhook'

  # HTTP port to receive hooks.
  port: 8080


### Configs for integrations.
bridges:
  - name: 'example-bridge'
    input:
      source: 'docker-hub'
    output:
      target: 'slack'
      options:
        url: 'https://hooks.slack.com/services/...'
    converter:
      name: 'DockerHubToSlack'
```


## Development

TBD
