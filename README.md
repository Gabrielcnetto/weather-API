
# Weather API

Personal weather API built in pure Go using net/http with no frameworks. Fetches and serves weather data with Redis caching for fast responses and reduced upstream calls. Designed for simplicity, performance, and learning core Go networking, concurrency, and caching patterns while exposing clean, REST-style endpoints.


## Screenshots

![Logic](https://i.ibb.co/GftZcmLq/weather-api-f8i1q.png)


## How use

To deploy this project run

```bash
  curl --location 'localhost:8080/weather' \
--header 'city: porto Alegre'
```

## Authors

- [@GabrielcNetto](https://github.com/Gabrielcnetto)
- [RoadmapSh](https://roadmap.sh/projects/weather-api-wrapper-service)
