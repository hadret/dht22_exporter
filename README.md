# DHT22 Exporter for Prometheus

[DHT22](https://www.sparkfun.com/datasheets/Sensors/Temperature/DHT22.pdf) Prometheus exporter for temperature and humidity.

## Getting up to speed

Grab one of the pre-built binaries available in [tags](https://github.com/hadret/dht22_exporter/tags) or a Docker container from GitHub' registry.

Binary run:

```help
./dht22_exporter [<flags>]
```

In the [contrib](contrib) directory there's also sample Systemd service file.

Docker run:

```shell
docker run --rm -it ghcr.io/hadret/dht22_exporter:latest -h
```

## Usage

```help
./dht22_exporter -h
Usage of ./dht22_exporter:
  -gpio-port int
        The GPIO port where DHT22 is connected. (default 2)
  -interval duration
        Temperature and humidity evaluation interval. (default 60ns)
  -listen-address string
        The address to listen on for HTTP requests. (default ":10005")
```

## License

Apache License 2.0, see [LICENSE](LICENSE).
