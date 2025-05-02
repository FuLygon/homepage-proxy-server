# WireGuard Widgets Configuration

Note that if you're using [wg-easy](https://github.com/wg-easy/wg-easy), Homepage already has it integrated [here](https://gethomepage.dev/widgets/services/wgeasy)

Currently the WireGuard integration support both locally installed WireGuard and WireGuard deployed via Docker.

## Setting up the method for fetching data

The env `SERVICE_WIREGUARD_METHOD` supported values are `local`, `docker` and `external`:

- `local`: this is mean to use only for local installation of WireGuard and local installation of the API Gateway, the service will execute `wg` command to fetch the data.
- `docker`: can be used with both local and docker installation of WireGuard.
- `external`: this is mean to use with the external script that notifies client connections, the script can be found [here](https://github.com/FuLygon/wireguard-client-connection-notification).

### Local Method

Required environment variable for this method is `SERVICE_WIREGUARD_INTERFACE`. Example `.env`:

```dotenv
SERVICE_WIREGUARD_ENABLED=true
SERVICE_WIREGUARD_METHOD=local
SERVICE_WIREGUARD_INTERFACE=wg0
SERVICE_WIREGUARD_TIMEOUT=5
```

You'll need to build the API Gateway from source, then run it locally, if the user current doesn't have required permissions to run `wg` command, you'll need to run the API Gateway as `root` or with `sudo`.

### Docker Method

Required environment variables for this method are `SERVICE_WIREGUARD_INTERFACE` and `SERVICE_WIREGUARD_DOCKER_CONTAINER`. Example `.env`:

```dotenv
SERVICE_WIREGUARD_ENABLED=true
SERVICE_WIREGUARD_METHOD=docker
SERVICE_WIREGUARD_INTERFACE=wg0
SERVICE_WIREGUARD_TIMEOUT=5
SERVICE_WIREGUARD_DOCKER_CONTAINER=<your_wireguard_container_name_or_id>
```

You also need to mount Docker socket to the API Gateway container:

```yaml
services:
  homepage-widgets-gateway:
    # other configurations for homepage-widgets-gateway service
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
```

#### Local installation of WireGuard

You'll need a container that has access to the host network and running in privileged mode, compose file example:

```yaml
services:
  alpine:
    image: alpine:latest
    container_name: wireguard-bridge
    entrypoint: ["/bin/sh", "-c"]
    network_mode: host
    privileged: true
    command:
      - |
        apk update
        apk add --no-cache wireguard-tools
        tail -f /dev/null
```

Then grab the container name or id and set it to `SERVICE_WIREGUARD_DOCKER_CONTAINER` in the `.env` file.

#### Docker installation of WireGuard

If you're using docker installation of WireGuard, you can set the `SERVICE_WIREGUARD_DOCKER_CONTAINER` to the container name or id of the WireGuard container.

Make sure the `wg` command is available in the container, you can check this by `exec` into the container.

### External Method

Example `.env`:

```dotenv
SERVICE_WIREGUARD_ENABLED=true
SERVICE_WIREGUARD_METHOD=external
```

You'll need to set up [the script](https://github.com/FuLygon/wireguard-client-connection-notification) and a cronjob for it.

The script will generate a list of clients in the `clients` folder, you'll need to mount this folder to the API Gateway container:

```yaml
services:
  homepage-widgets-gateway:
    # other configurations for homepage-widgets-gateway service
    volumes:
      - /path_to_the_clients_folder:/app/wireguard-clients
```

The API Gateway will read the content of the folder and generate the API response.

## Example widget config

```yaml
widget:
  type: customapi
  url: http://homepage-widgets-gateway:8080/wireguard
  refreshInterval: 300000 # 5 minutes
  method: GET
  mappings:
    - field: total
      label: Total # customize this to your liking
      format: number
    - field: connected
      label: Connected # customize this to your liking
      format: number
```

## API Response

```json
{
  "total": 5,
  "connected": 2
}
```

- `total` `int`: The total number of clients.
- `connected` `int`: The number of connected clients.
