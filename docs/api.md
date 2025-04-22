## API Definition

- `GET /health`

  - Response:
    ```json
    {
      "status": "healthy"
    }
    ```

- `GET /adguard-home`

  - Response:
    ```json
    {
      "queries": 50,
      "blocked": 10,
      "filtered": 2,
      "latency": 5.15
    }
    ```
- `GET /nginx-proxy-manager`

  - Response:
    ```json
    {
      "total": 15,
      "enabled": 10,
      "disabled": 5
    }
    ```
- `GET /portainer`

  - Response:
    ```json
    {
      "total": 35,
      "running": 25,
      "stopped": 10
    }
    ```

- `GET /wud`

  - Response:
    ```json
    {
      "monitoring": 10,
      "updates": 2
    }
    ```

- `GET /gotify`

  - Response:
    ```json
    {
      "applications": 5,
      "clients": 3,
      "messages": 150
    }
    ```

- `GET /uptime-kuma`

  - Response:
    ```json
    {
      "sites-up": 12,
      "sites-down": 3,
      "uptime": 99.5
    }
    ```
