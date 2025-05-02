# Your Spotify widget configuration

## Setting up environment variables

You'll need to set `SERVICE_YOUR_SPOTIFY_TOKEN` to fetch data, you can grab this token by go to Your Spotify Dashboard > Settings > Account Tab > Public token.

Generate your token, then grab the token value from the URL, e.g if your public token URL is `https://yourspotify.com/?token=faab0509-de47-4282-961a-113050b3331f`, then the value for `SERVICE_YOUR_SPOTIFY_TOKEN` is `faab0509-de47-4282-961a-113050b3331f`

## Example widget config

The Custom API widget URL contain a query parameter `time_range`, all of the valid value are:

- `day`: fetch data from yesterday
- `week`: fetch data from last week
- `month`: fetch data from last month
- `year`: fetch data from last year
- `all`: fetch all data up to the current time

```yaml
widget:
  type: customapi
  url: http://homepage-widgets-gateway:8080/your-spotify/?time_range=month
  # 5 minutes - customize this to your liking
  # Keep in mind the service also cached the response data for 5 minutes since Your Spotify also doesn't fetch data regularly, so there's no need to set this to a lower value.
  refreshInterval: 300000
  method: GET
  mappings:
    - field: songs_listened
      label: Songs # customize this to your liking
      format: number
    - field: time_listened
      label: Time # customize this to your liking
      format: number
      suffix: min # customize this to your liking, this is the text that will appear after the value
    - field: artists_listened
      label: Artists # customize this to your liking
      format: number
```

## API Response

`GET /your-spotify`

```json
{
  "songs_listened": 10,
  "time_listened": 120,
  "artists_listened": 5
}
```

- `songs_listened` `int64`: The number of songs listened.
- `time_listened` `int64`: The total time (_minute_) spent listening to music.
- `artists_listened` `int`: The number of unique artists listened.
