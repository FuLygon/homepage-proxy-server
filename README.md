# Proxy Mapping Server for Homepage

[![Publish to GitHub Container Registry](https://github.com/FuLygon/homepage-proxy-server/actions/workflows/publish-package.yaml/badge.svg)](https://github.com/FuLygon/homepage-proxy-server/actions/workflows/publish-package.yaml)

A simple server for integrating with [Homepage](https://github.com/gethomepage/homepage) widgets.

_But homepage already have a bunch of widgets already integrated?_
Yes it does, and if you didn't exposing your Homepage instance to the public, you can ignore this and use existing integrations homepage already have.

But if you are exposing your homepage instance to the public, you need to know that some Homepage integrations are returning **way too much** data to the client, including _possibly_ sensitive data. You can verify this with browser **DevTools**, and start inspecting the response of all the `GET /api/services/proxy` API made by Homepage. Some of the example widgets that are responding too much data than needed are:

- [Adguard Home](https://gethomepage.dev/widgets/services/adguard-home): the response include information such as most queried/blocked domains, client list and their IP,...
- [Nginx Proxy Manager](https://gethomepage.dev/widgets/services/nginx-proxy-manager): the response include information such as the proxy hosts, their domain name, forwarded host/port, advanced config,...
- [Portainer](https://gethomepage.dev/widgets/services/portainer): the response include information related to docker like container command/entrypoint, which image was used, port and network settings,...
- [Gotify](https://gethomepage.dev/widgets/services/gotify): the response include a lot of sensitive data such as application/client token, message content,...

If that concern you, you can use this with Homepage [Custom API](https://gethomepage.dev/widgets/services/customapi) widget to minimize the amount the data responded that the UI needed to render.

## How it works

Fairly simple though, Homepage will make an API request of a specific integration via this service instead of directly into the target integrating service. This service will then make a request to the target to fetch the data, similar to how the existing Homepage integrations work. But this service process the response and only return the data needed by the UI. This way, no unnecessary data will be exposed to the public.

Most of the integrations and mapping will be using existing implementation from [source code](https://github.com/gethomepage/homepage/tree/dev/src/widgets) by Homepage as a reference to process the data.

## Supported Integrations

Currently only support a few integrations:

- Adguard Home
- Nginx Proxy Manager
- Portainer
- WUD (What's Up Docker)
- Gotify
- Uptime Kuma (currently not including Incident data since I can't figure out to make a field appear dynamically with Custom API)

Initially I planned to make this only for my own Homepage instance. So if there are any integrations that aren't supported, then tough luck. I don't plan to add more integrations unless I need them.

But if you implement your own integrations, feel free to open a PR if you want and I'll check it out.

I _might_ as well using this repo to implement other integrations that I need but aren't supported by Homepage.

## Usage and configuration

### Docker Installation

- Prepare the `.env` file:

```bash
wget https://raw.githubusercontent.com/fulygon/homepage-proxy-server/main/.env.example -O .env
```

- Modify the `.env` to your need. Then deploy the service with docker. Compose file example:

```yaml
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

### Source Installation

Make sure [Go](https://go.dev/doc/install) is installed.

- Clone the repo:

```bash
git clone https://github.com/FuLygon/homepage-proxy-server.git
cd homepage-proxy-server
```

- Prepare the `.env` file:

```bash
cp .env.example .env
```

- Modify the `.env` file to your need. Then build and run:

```bash
go build -o homepage-proxy-server ./cmd/main.go
./homepage-proxy-server
```

### Homepage Configuration

An example configuration for homepage Custom API widget configuration can be found [here](docs/homepage-widgets.md).

## API Documentation

I don't think this need a proper API Documentation, I did write a fairly simple API definition in [here](docs/api.md) if you want to check it out.
