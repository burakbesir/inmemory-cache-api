# In Memory Cache API
The project purpose is a simple in memory cache application you can use with get and set endpoints.

## Application Startup
You can use directly `go run .` command on project root directory or if you want to use docker you can use that commands.
```sh
docker build -t cache .
docker run -p 8080:8080 cache:latest
```
After that you can connect to api on 8080 port.

## Usage
You can set a key value pair to cache with a simple post request. A curl example :
```sh
curl -X POST -F 'key=k1' -F 'value=v1' localhost:8080/cache
```
And if you want to reach value previously you saved, you can use get api with a key parameter
```sh
curl http://localhost:8080/cache?key=k1
```