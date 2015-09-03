#/bin/bash

docker run -it \
	-v /var/run/:/var/run \
	-v /var/lib/docker:/var/lib/docker \
	dg-probe
