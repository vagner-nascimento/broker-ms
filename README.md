# broker-ms
A http service that publishes messages into AMQP topics

# enviroment
- `LISTEN_PORT <default 80>`: the port to starts the http server

# Local: Install and run
- `dependencies`: go mod download
- `build`: go build -o main .
- `env`: export LISTEN_PORT=<port>
- `run`: /path/to/app/main

# Docker: Install and run
- `build`: dokcer build -t <your_namespace>/broker-ms:<version_tag>
- `run`: docker run -e LISTEN_PORT=<port> -p "<local_port>:<cont_port>" <your_namespace>/broker-ms:<version_tag>

# Local access on docker or local
- `URL`
