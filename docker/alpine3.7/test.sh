# from https://github.com/restic/restic/tree/master/docker
#!/bin/sh

set -e

docker run --rm -ti \
  -h btpd \
  -p 8080:8080 \
  -v $(pwd)/assets:/usr/bptd/assets \
	-v $(pwd)/tmp/data:/usr/bptd/data \
  -w /usr/bptd \
  xtforgame/bptd:latest bptd ./data
