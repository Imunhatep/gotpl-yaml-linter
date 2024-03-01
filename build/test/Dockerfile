FROM hub.rms.evolutiongaming.com/hub/golang:1.22.0-alpine3.19

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

RUN go test -v -coverprofile=int-coverage.out ./internal/... \
    && go tool cover -func=int-coverage.out

#RUN go test -v -coverprofile=coverage.out ./pkg/collection/... \
#    && go tool cover -func=coverage.out
#    && go run ./cmd/policy/main.go policy -f policy.yaml