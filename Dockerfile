#FROM golang:latest as build-env
FROM scratch
#COPY --from=build-env /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ARG TARGETOS
ARG TARGETARCH
ADD /${TARGETOS}/${TARGETARCH} /
