# from https://github.com/restic/restic/tree/master/docker
#!/bin/sh

set -e

# https://medium.com/travis-on-docker/how-to-cross-compile-go-programs-using-docker-beaa102a316d
docker run --rm -it -v "$GOPATH":/go -w /go/src/github.com/xtforgame/restfs golang:1.10-alpine3.7 sh -c '
for GOARCH in 386 amd64; do
  for GOOS in darwin linux windows freebsd; do
    echo "Building $GOOS-$GOARCH"
    export GOOS=$GOOS
    export GOARCH=$GOARCH
    go build -o build/bin/bptd-$GOOS-$GOARCH bptd.go
    go build -o build/bin/simple_client-$GOOS-$GOARCH simple_client.go
  done
done
for GOARCH in arm ; do
  for GOOS in linux freebsd; do
    echo "Building $GOOS-$GOARCH"
    export GOOS=$GOOS
    export GOARCH=$GOARCH
    go build -o build/bin/bptd-$GOOS-$GOARCH bptd.go
    go build -o build/bin/simple_client-$GOOS-$GOARCH simple_client.go
  done
done
for GOARCH in arm64 ; do
  for GOOS in linux; do
    echo "Building $GOOS-$GOARCH"
    export GOOS=$GOOS
    export GOARCH=$GOARCH
    go build -o build/bin/bptd-$GOOS-$GOARCH bptd.go
    go build -o build/bin/simple_client-$GOOS-$GOARCH simple_client.go
  done
done
'
