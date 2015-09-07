# Docker Guard Probe

## What?

Docker Guard is a powerful monitoring tool to watch your containers (running or not, memory/disk/netio usage, ...)

This tool is the probe of the [Docker Guard Monitoring system](https://github.com/90TechSAS/docker-guard-monitoring)

## How to install?

You must have Docker 1.8.2 or newer installed (to display the version, type in the console: ```docker version```).

Now, let's run the probe:

```bash
git clone https://github.com/90TechSAS/docker-guard-probe.git
cd docker-guard-probe/docker
```

Edit the file ```config.yaml``` at your own sweet will.
Type these commands to build a container with the probe inside and run it!

```bash
./build.sh
./run.sh
```

An ID will be displayed (example: b0ae690e631af71cd768d08b33205adc9474e57302b943805168becf82d75045), copy and paste it like this:

```bash
docker inspect -f {{.State.Running}} b0ae690e631af71cd768d08b33205adc9474e57302b943805168becf82d75045
```

If everything is ok, this command displays "true".

## How to configure?

The file config.yaml is the config file of the probe. When you clone the project, this file is absent and ignored by git to prevent commiting your configuration (and you password btw). So first, you need to copy the file config.yaml.example like this:

```bash
cp config.yaml.example config.yaml
```

And then you can edit the configuration by editing config.yaml.
There are comments in this file who explains how it works, don't worry.

## How to contribute?

Feel free to fork the project and make a pull request!

## Thanks to

* [LoGo](https://github.com/Nurza/LoGo)
* [Gorilla Mux](https://github.com/gorilla/mux)
* [Go-yaml](https://github.com/go-yaml/yaml)

## License

MIT
