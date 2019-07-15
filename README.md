Simple golang http client app which tries to call $URL $SESSION times.

```
docker build -t http .
docker run --rm -e URL=https://google.com -e SESSIONS=30 http # this will call https://google.com 30 times
```
