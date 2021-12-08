FROM golang:1.16.0 as builder

WORKDIR /github.com/GoogleCloudPlatform/terraformer

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go run build/main.go sumologic

FROM alpine:3.12.0 

ENV TERRAFORM_VERSION=1.0.11

RUN apk update && \
    apk add curl jq python3 bash ca-certificates git openssl unzip wget && \
    cd /tmp && \
    wget https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_linux_amd64.zip && \
    unzip terraform_${TERRAFORM_VERSION}_linux_amd64.zip -d /usr/local/bin && \
    rm -rf /tmp/*

COPY --from=builder /github.com/GoogleCloudPlatform/terraformer/terraformer-sumologic /usr/local/bin/terraformer-sumologic
COPY OfficialProviders.tf /root
RUN cd /root && terraform init
WORKDIR /root/