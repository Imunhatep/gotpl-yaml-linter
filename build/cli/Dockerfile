FROM hub.rms.evolutiongaming.com/hub/golang:1.22.0-alpine3.19 AS builder

ARG TARGETPLATFORM
ARG TARGETOS
ARG TARGETARCH
ARG DBUILD_DATE

ENV CGO_ENABLED=0
ENV GOOS=$TARGETOS
ENV GOARCH=$TARGETARCH

RUN apk add --update bash git make

ADD . /project/
WORKDIR /project/

RUN mkdir -p ./bin \
    && env GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o ./bin/yamltpl_linter_${TARGETOS}-${TARGETARCH} -ldflags "-s -w" ./cmd/cli/main.go

FROM hub.rms.evolutiongaming.com/hub/alpine:3.19

RUN apk add --update bash ca-certificates curl

ENV GOMEMLIMIT=256MiB

COPY --from=builder /project/bin/yamltpl_linter_${TARGETOS}-${TARGETARCH}  /usr/bin/yamltpl_linter

RUN chmod +x /usr/bin/yamltpl_lint \
    && adduser -D service

USER service
WORKDIR /home/service

CMD ["yamltpl_linter", "lint"]
ENTRYPOINT []