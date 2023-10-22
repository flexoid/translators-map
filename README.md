# Sworn translator map

Polish sworn translators on map.

https://sworntranslatormap.com

![Screen recording](./docs/screen-recording.gif)

## Configuration

Run with `--help` or check out [config.go](./internal/config/config.go) for configuration options.

## Deployment

See included [docker-compose.yml](./docker-compose.yml) for a deployment example.

## Development

Run backend process:

```
go run main.go server

# or with hot reload using Air
air
```

Run scraper:

```
go run main.go scraper
```

Run frontend dev server:

```
./web
npm run dev
```

To regenerate and compile translations run:

```
npm run i18n
```
