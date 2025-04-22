# Proxy Mapping Server for Homepage

## Usage
Using docker compose:
```yaml
version: '3.8'

services:
  homepage:
    image: ghcr.io/gethomepage/homepage:latest
    container_name: homepage
    # other configurations for homepage service

  homepage-proxy-server:
    image: ghcr.io/fulygon/homepage-proxy-server:latest
    container_name: homepage-proxy-server
    env_file: .env
```