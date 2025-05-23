## Example widgets configuration for Homepage

There's no need to include widget credentials like password or token since it already configurated in the API Gateway environment variables.

### [Adguard Home](https://gethomepage.dev/widgets/services/adguard-home)

```yaml
widget:
  type: adguard
  url: http://homepage-widgets-gateway:8080/adguard-home
```

### ASF (ArchiSteamFarm)

```yaml
widget:
  type: customapi
  url: http://homepage-widgets-gateway:8080/asf
  refreshInterval: 300000 # 5 minutes
  method: GET
  mappings:
    - field: total
      label: Total
      format: number
    - field: online
      label: Online
      format: number
    - field: cards_remaining
      label: Cards Remaining
      format: number
```

### [Gotify](https://gethomepage.dev/widgets/services/gotify)

```yaml
widget:
  type: gotify
  url: http://homepage-widgets-gateway:8080/gotify
```

### [Komodo](komodo.md)

```yaml
widget:
  type: customapi
  url: http://homepage-widgets-gateway:8080/komodo
  refreshInterval: 300000 # 5 minutes
  method: GET
  mappings:
    - field: container.running
      label: Running
      format: number
    - field: container.stopped
      label: Stopped
      format: number
    - field: container.total
      label: Total
      format: number
```

### [Linkwarden](https://gethomepage.dev/widgets/services/linkwarden)

```yaml
widget:
  type: linkwarden
  url: http://homepage-widgets-gateway:8080/linkwarden
```

### [Nginx Proxy Manager](https://gethomepage.dev/widgets/services/nginx-proxy-manager)

```yaml
widget:
  type: npm
  url: http://homepage-widgets-gateway:8080/nginx-proxy-manager
```

### [Portainer](https://gethomepage.dev/widgets/services/portainer)

```yaml
widget:
  type: portainer
  url: http://homepage-widgets-gateway:8080/portainer
  env: 1
```

### [Uptime Kuma](https://gethomepage.dev/widgets/services/uptime-kuma)

```yaml
widget:
  type: uptimekuma
  url: http://homepage-widgets-gateway:8080/uptime-kuma
  slug: statuspageslug
```

### [WireGuard](wireguard.md)

```yaml
widget:
  type: customapi
  url: http://homepage-widgets-gateway:8080/wireguard
  refreshInterval: 300000 # 5 minutes
  method: GET
  mappings:
    - field: total
      label: Total
      format: number
    - field: connected
      label: Connected
      format: number
```

### [WUD (What's Up Docker)](https://gethomepage.dev/widgets/services/whatsupdocker)

```yaml
widget:
  type: whatsupdocker
  url: http://homepage-widgets-gateway:8080/wud
```

### [Your Spotify](your-spotify.md)

```yaml
widget:
  type: customapi
  url: http://homepage-widgets-gateway:8080/your-spotify/?time_range=month
  refreshInterval: 300000
  method: GET
  mappings:
    - field: songs_listened
      label: Songs
      format: number
    - field: time_listened
      label: Time
      format: number
      suffix: min
    - field: artists_listened
      label: Artists
      format: number
```
