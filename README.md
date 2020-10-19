# stupid-http-mock

![Docker Cloud Automated build](https://img.shields.io/docker/cloud/automated/kuskoman/stupid-http-mock)
![Docker Pulls](https://img.shields.io/docker/pulls/kuskoman/stupid-http-mock)

This is a really small application used to mock
server in very basic cases.

## Installation

You can build this app from source or use prebuilt image from [DockerHub](https://hub.docker.com/).

## Building from source

1. Clone git repository

   ```shell
   git clone git@github.com:kuskoman/stupid-http-mock.git
   ```

2. Build repository

   ```shell
   go build main.go
   ```

This should result with ready-to-use executable.

## Using image from DockerHub

Simply use ([Docker](https://docs.docker.com/get-docker/) required)

```shell
docker run kuskoman/stupid-http-mock:latest -p <host port>:<container port>
```

for example:

```shell
docker run kuskoman/stupid-http-mock:latest -p 1234:8080
```

will run app on host port 1234.

## Usage

The application takes requests, responds with [204 no content](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/204)
and simply logs it's properties.

Example of logs (after entering app url in browser using firefox
with random user-agent plugin):

```log
2020/10/19 15:30:29 Request received
URL: /
Method: GET
Host: localhost:8080
Body:
Headers:
  Cookie: csrftoken=zukg4pEjcCjeQfi3QpvmFaNH288TPxgeau5h6HpQ0GMcHBboQg8VEr61oJk2quFH
  Upgrade-Insecure-Requests: 1
  User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36
  Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8
  Accept-Language: en-US,en;q=0.5
  Accept-Encoding: gzip, deflate
  Connection: keep-alive
```

## Why

Because I needed very simple mock server many times
and I decided to create 'full' app, instead of short
scripts every time it was needed.
