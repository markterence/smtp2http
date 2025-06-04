SMTP2HTTP (email-to-web)
========================
smtp2http is a simple smtp server that resends the incoming email to the configured web endpoint (webhook) as a basic http post request.

Dev 
===
- `go mod vendor`
- `go build`

Dev with Docker
==============
Using pre-built docker image:

```
docker pull ghcr.io/markterence/smtp2http:latest

docker run -p 25:25 ghcr.io/markterence/smtp2http:latest \ 
    --webhook=http://some.hook/api
```

Locally :
- `go mod vendor`
- `docker build -f Dockerfile.dev -t smtp2http-dev .`
- `docker run -p 25:25 smtp2http-dev --timeout.read=50 --timeout.write=50 --webhook=http://some.hook/api`

Or build it as it comes from the repo :
- `docker build -t smtp2http .`
- `docker run -p 25:25 smtp2http --timeout.read=50 --timeout.write=50 --webhook=http://some.hook/api`

The `timeout` options are of course optional but make it easier to test in local with `telnet localhost 25`
Here is a telnet example payload : 
```
HELO zeus
# smtp answer

MAIL FROM:<email@from.com>
# smtp answer

RCPT TO:<youremail@example.com>
# smtp answer

DATA
your mail content
.

```

Docker pre-built image are available and only on ghcr
=====

**docker-compose.yml** Example:

```yml
services:
  smtp2http:
    image: ghcr.io/markterence/smtp2http:latest
    container_name: smtp2http
    restart: unless-stopped
    ports:
      - "2525:25"
    env_file:
      - .env
    command: [
      "-listen=0.0.0.0:25",
      "-webhook=${WEBHOOK_URL}",
      "-user", "${SMTP_USER}",
      "-pass", "${SMTP_PASSWORD}",
      "-name", "smtp2http"
    ]
```

Native usage
=====
`smtp2http --listen=:25 --webhook=http://localhost:8080/api/smtp-hook`
`smtp2http --help`

Contribution
============
Original repo from @alash3al
Thanks to @aranajuan and @alash3al


