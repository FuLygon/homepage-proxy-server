# Komodo widget configuration

## Setting up environment variables

Go to Komodo Settings > Profile > API Keys > Create New API key, after that grab the `key` and `secret` and set it to `SERVICE_KOMODO_KEY` and `SERVICE_KOMODO_SECRET` respectively.

By default, the API only fetch container data, you can set the API to fetch more data by setting `SERVICE_KOMODO_EXTRA_STATS`, separated by comma, valid values are:

- `stack`
- `build`
- `repo`
- `action`
- `builder`
- `deployment`
- `procedure`
- `resource-sync`

## Example widget config

```yaml
widget:
  type: customapi
  url: http://homepage-widgets-gateway:8080/komodo
  refreshInterval: 300000 # 5 minutes
  method: GET
  mappings:
    - field: container.running
      label: Running # customize this to your liking
      format: number
    - field: container.stopped
      label: Stopped # customize this to your liking
      format: number
    - field: container.total
      label: Total # customize this to your liking
      format: number
    - field: stack.total # if you configured `SERVICE_KOMODO_EXTRA_STATS` to include `stack`
      label: Stack # customize this to your liking
      format: number
```

## API Response

With all extra stats included:

```dotenv
SERVICE_KOMODO_EXTRA_STATS=stack,build,repo,action,builder,deployment,procedure,resource-sync
```

`GET /komodo`

```json
{
  "container": {
    "total": 10,
    "running": 6,
    "stopped": 4,
    "unhealthy": 0,
    "unknown": 0
  },
  "stack": {
    "total": 3,
    "running": 2,
    "stopped": 1,
    "down": 0,
    "unhealthy": 0,
    "unknown": 0
  },
  "build": {
    "total": 3,
    "ok": 3,
    "failed": 0,
    "building": 0,
    "unknown": 0
  },
  "repo": {
    "total": 1,
    "ok": 1,
    "cloning": 0,
    "pulling": 0,
    "building": 0,
    "failed": 0,
    "unknown": 0
  },
  "action": {
    "total": 0,
    "ok": 0,
    "running": 0,
    "failed": 0,
    "unknown": 0
  },
  "builder": {
    "total": 1
  },
  "deployment": {
    "total": 0,
    "running": 0,
    "stopped": 0,
    "not_deployed": 0,
    "unhealthy": 0,
    "unknown": 0
  },
  "procedure": {
    "total": 0,
    "ok": 0,
    "running": 0,
    "failed": 0,
    "unknown": 0
  },
  "resource-sync": {
    "total": 0,
    "ok": 0,
    "syncing": 0,
    "pending": 0,
    "failed": 0,
    "unknown": 0
  }
}
```
