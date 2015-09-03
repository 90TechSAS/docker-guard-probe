#/bin/bash

docker run -itd \
	-v /var/run/:/var/run \
	-v /var/lib/docker:/var/lib/docker \
	-p 8123:8123 \
	dg-probe
