FROM golang:alpine AS builder

WORKDIR /build

COPY . .

RUN go mod download

RUN go build -o crm.shopdev.com ./cmd

FROM scratch

COPY ./deploy/conf /deploy/conf

COPY --from=builder /build/crm.shopdev.com /

ENTRYPOINT ["/crm.shopdev.com", "deploy/conf/local.yaml"]



