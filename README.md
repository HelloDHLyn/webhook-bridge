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

### Supported Converters

  - `func`

### Configurations

```yaml
### Version of the configration file template.
version: '1'


### Configs for incoming webhooks.
# URL path. Path for the bridge named `example-bridge` becomes `/webhook/example-bridge`.
input_path_prefix: '/webhook'

# HTTP port to receive hooks.
input_port: 8080


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
