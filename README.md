# Docker Guard Probe

## What?

Docker Guard is a powerful monitoring tool to watch your containers (running or not, memory/disk/netio usage, ...)

This tool is the probe of the ![Docker Guard Monitoring system](https://github.com/90TechSAS/docker-guard-monitoring)

## How to install?

You must have Docker 1.8.2 or newer (to display the version, type in the console: ```docker version```)

Now, let's run the probe:

```bash
git clone https://github.com/90TechSAS/docker-guard-monitoring.git
cd docker-guard-monitoring/docker
```

Edit the file ```config.yaml``` at your own sweet will.
Type these commands to build a container with the probe inside and run it!

```bash
./build.sh
./run.sh
```

## How to configure?

TODO

## How to contribute?

Feel free to fork the project a make a pull request!

## Thanks to

* [LoGo](https://github.com/Nurza/LoGo)
* [Gorilla Mux](github.com/gorilla/mux)
* [Go-yaml](https://github.com/go-yaml/yaml)

## License

MIT
