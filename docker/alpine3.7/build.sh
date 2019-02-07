# from https://github.com/restic/restic/tree/master/docker
#!/bin/sh

set -e

echo "Build binary using golang docker image"
docker run --rm -ti \
  -v $(pwd):/go/src/github.com/xtforgame/restfs \
  -w /go/src/github.com/xtforgame/restfs \
  -e CGO_ENABLED=1 -e GOOS=linux golang:1.10-alpine3.7 go build -o ./build/alpine3.7/bptd bptd.go

docker run --rm -ti \
  -v $(pwd):/go/src/github.com/xtforgame/restfs \
  -w /go/src/github.com/xtforgame/restfs \
  -e CGO_ENABLED=1 -e GOOS=linux golang:1.10-alpine3.7 go build -o ./build/alpine3.7/simple_client simple_client.go

echo "Build docker image xtforgame/bptd:latest"
docker build --rm -t xtforgame/bptd:latest -f docker/alpine3.7/Dockerfile .

# docker run --rm -ti \
#   -p 8080:8080 \
#   -v $(pwd)/tmp:/usr/restfs \
#   -w /usr/restfs \
#   xtforgame/bptd:latest bptd ./forweb ./pgbackrest-backup ./output
