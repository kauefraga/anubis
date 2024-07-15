# Anubis

![Anubis badge](https://img.shields.io/badge/anubis-004f72)
[![GitHub's license](https://img.shields.io/github/license/kauefraga/anubis)](https://github.com/kauefraga/anubis/blob/main/LICENSE)
[![GitHub last commit (branch)](https://img.shields.io/github/last-commit/kauefraga/anubis/main)](https://github.com/kauefraga/anubis)

> A powerful and configurable load balancer, built in Go.

## Configuration

Create a file called `anubis.toml` and define the following attributes:

- `version` - Anubis version (default: `1`)
- `port` - Anubis port (default: `4000`)
- `algorithm` - algorithm you want to use, `round-robin`, `least-connection` or `weighted-response-time` (default: `round-robin`)
- `servers` - array of servers address

Example of minimal configuration:

```toml
servers = ['localhost:4001', 'localhost:4002', 'localhost:4003']
```

Resulting in `version = 1`, `port = 4000`, `algorithm = 'round-robin'` and  `servers = ['localhost:4001', 'localhost:4002', 'localhost:4003']`.

Example of full configuration:

```toml
version = 1
port = 3333
algorithm = 'round-robin'
servers = ['localhost:3334', 'localhost:3335', 'localhost:3336']
```

> [!IMPORTANT]
> More coming soon...
