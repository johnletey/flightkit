# FlightKit

Go clients for flight tracking APIs.

| Package                      | Network                                        |
| ---------------------------- | ---------------------------------------------- |
| [`adsblol`](adsblol)         | [adsb.lol](https://adsb.lol)                   |
| [`airnavradar`](airnavradar) | [AirNav Radar](https://www.airnavradar.com)    |
| [`flightaware`](flightaware) | [FlightAware](https://www.flightaware.com)     |
| [`fr24`](fr24)               | [Flightradar24](https://www.flightradar24.com) |
| [`opensky`](opensky)         | [OpenSky Network](https://opensky-network.org) |
| [`planefinder`](planefinder) | [Plane Finder](https://planefinder.net)        |

Every package exposes a `ServerURL` constant for the provider's API base URL, and packages whose API needs credentials have a `WithAuthentication` option.

```go
client, err := fr24.NewClientWithResponses(fr24.ServerURL, fr24.WithAuthentication(apiKey))
```

If you're missing a provider, open an issue or a pull request.

## Regenerating clients

A [weekly GitHub Action](.github/workflows/sync-openapi.yml) pulls each provider's latest spec, regenerates the client with [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen), and opens a pull request if anything changed.

## License

[Apache 2.0](LICENSE)
