FROM alpine:3.7

COPY build/alpine3.7/bptd /usr/bin
RUN chmod +x /usr/bin/bptd
# COPY docker/alpine3.7/restic /usr/bin
# RUN chmod +x /usr/bin/restic
# COPY docker/alpine3.7/azpbrctl /usr/bin
# RUN chmod +x /usr/bin/azpbrctl
# COPY docker/alpine3.7/rclone /usr/bin
# RUN chmod +x /usr/bin/rclone

# RUN apk add --update --no-cache ca-certificates fuse

# ENTRYPOINT ["/usr/bin/bptd"]
