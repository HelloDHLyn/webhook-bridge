# Webhook Bridge

> Receive webhooks on your server, and integrate them with others.

## Usage

### Supported Targets

| Target | Options |
| --- | --- |
| Slack | `url` : incoming webhook address |

### Example 1. Declarative YAML Config

First, create a configuration file like below.

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
    json: |
      {
        "text": "Docker image {{ .repository.repo_name }} has built successfully."
      }
```

Then run docker container.

```sh
docker run \
  -v '/path/to/config.yaml:/etc/config.yaml' \
  -e 'BRIDGE_CONFIG_PATH=/etc/config.yaml' \
  -p '8080:8080' \
  hellodhlyn/webhook-bridge
```

## Development

TBD
