## Example Custom API widgets configuration for homepage

For more information on custom API widgets, please see https://gethomepage.dev/widgets/services/customapi

### Adguard Home

```yaml
widget:
  type: customapi
  url: http://homepage-proxy-server:8080/adguard-home
  refreshInterval: 10000
  method: GET
  mappings:
    - field: queries
      label: queries
      format: number
    - field: blocked
      label: blocked
      format: number
    - field: filtered
      label: filtered
      format: number
    - field: latency
      label: latency
      format: float
      suffix: ms
```

### Nginx Proxy Manager

```yaml
widget:
  type: customapi
  url: http://homepage-proxy-server:8080/nginx-proxy-manager
  refreshInterval: 10000
  method: GET
  mappings:
    - field: enabled
      label: enabled
      format: number
    - field: disabled
      label: disabled
      format: number
    - field: total
      label: total
      format: number
```

### Portainer

```yaml
widget:
  type: customapi
  url: http://homepage-proxy-server:8080/portainer
  refreshInterval: 10000
  method: GET
  mappings:
    - field: running
      label: running
      format: number
    - field: stopped
      label: stopped
      format: number
    - field: total
      label: total
      format: number
```

### WUD (What's Up Docker)

```yaml
widget:
  type: customapi
  url: http://homepage-proxy-server:8080/wud
  refreshInterval: 10000
  method: GET
  mappings:
    - field: monitoring
      label: monitoring
      format: number
    - field: updates
      label: updates
      format: number
```

### Gotify

```yaml
widget:
  type: customapi
  url: http://homepage-proxy-server:8080/gotify
  refreshInterval: 10000
  method: GET
  mappings:
    - field: applications
      label: applications
      format: number
    - field: clients
      label: clients
      format: number
    - field: messages
      label: messages
      format: number
```

### Uptime Kuma

```yaml
widget:
  type: customapi
  url: http://homepage-proxy-server:8080/uptime-kuma
  refreshInterval: 10000
  method: GET
  mappings:
    - field: sites-up
      label: sites up
      format: number
    - field: sites-down
      label: sites down
      format: number
    - field: uptime
      label: uptime
      format: percent
```
