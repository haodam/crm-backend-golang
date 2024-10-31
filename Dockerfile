#FROM golang:alpine AS builder

#WORKDIR /build
#
#COPY . .
#
#RUN go mod download
#
#RUN go build -o crm.shopdev.com ./cmd
#
#FROM scratch
#
#COPY ./configs /configs
#
#COPY --from=builder /build/crm.shopdev.com /
#
#ENTRYPOINT [ "/crm.shopdev.com", "configs/local.yaml" ]

FROM golang:alpine AS builder

# Cài đặt các công cụ cần thiết
RUN apk add --no-cache git

# Thiết lập thư mục làm việc
WORKDIR /app

# Copy toàn bộ mã nguồn vào container
COPY . .

# Tải xuống các dependencies của Go
RUN go mod download

# Expose port (nếu ứng dụng của bạn chạy trên một port cụ thể)
EXPOSE 8002

# Sử dụng lệnh go run để phát triển mà không cần build lại toàn bộ
CMD ["go", "run", "./cmd"]



