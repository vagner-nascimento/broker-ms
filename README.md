# broker-ms
A http service that publishes messages into AMQP topics

# enviroment
- `LISTEN_PORT <default 80>`: the port to starts the http server
- `AMQP_CONN`: amqp connection string

# Local: Install and run
- `dependencies`: go mod download
- `build`: go build -o main .
- `env`: export LISTEN_PORT=port_number
- `env`: export AMQP_CONN=amqp://guest:guest@localhost:5672
- `run`: /path/to/app/main

# Docker: Install and run
- `build`: dokcer build -t <your_namespace>/broker-ms:<version_tag>
- `run`: docker run -e LISTEN_PORT=<port> - e AMQP_CONN=<amqp_conn_str> -p "<local_port>:<cont_port>" <your_namespace>/broker-ms:<version_tag>

# Docker Compose: Run app and RabbitMQ
- `run`: docker-composer -f docker-compose.yml up --build

# Local access on docker or local
- `URL`: http://localhost:<port>/publish

# Payload
```json
[
	{
		"data": [{}],
	    "topic": "string"
    }
]
```
The "data" array can be any objects
