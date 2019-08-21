# urlshort

> [exercise 2 from Gophercises](https://gophercises.com/exercises/urlshort)

```bash
# Build
make compile
# Run db + webapp
docker-compose up --build -d
# Test mapping from maps
curl -L http://localhost:8080/urlshort-godoc
curl -L http://localhost:8080/yaml-godoc
# Test mapping from YAML file
curl -L http://localhost:8080/urlshort
curl -L http://localhost:8080/urlshort-final
# Test mapping from JSON file
curl -L http://localhost:8080/so
curl -L http://localhost:8080/tw
# Test mapping from DB
curl -L http://localhost:8080/gh
curl -L http://localhost:8080/gophercises
```
