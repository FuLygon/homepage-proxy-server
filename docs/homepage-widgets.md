## Example widgets configuration for Homepage

Check the official Homepage widgets [documentation](https://gethomepage.dev/widgets) if there is any parameter you don't understand.

There's no need to include widget credentials like password or token since it already configurated in the API Gateway environment variables.

### [Adguard Home](https://gethomepage.dev/widgets/services/adguard-home)

```yaml
widget:
  type: adguard
  url: http://homepage-widgets-gateway:8080/adguard-home
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

### [WUD (What's Up Docker)](https://gethomepage.dev/widgets/services/whatsupdocker)

```yaml
widget:
  type: whatsupdocker
  url: http://homepage-widgets-gateway:8080/wud
```

### [Gotify](https://gethomepage.dev/widgets/services/gotify)

```yaml
widget:
  type: gotify
  url: http://homepage-widgets-gateway:8080/gotify
```

### [Uptime Kuma](https://gethomepage.dev/widgets/services/uptime-kuma)

```yaml
widget:
  type: uptimekuma
  url: http://homepage-widgets-gateway:8080/uptime-kuma
  slug: statuspageslug
```
