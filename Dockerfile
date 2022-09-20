FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod tidy
RUN go mod download

# RUN GRPC_HEALTH_PROBE_VERSION=v0.4.12 && \
#     wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-darwin-amd64 && \
#     chmod +x /bin/grpc_health_probe

COPY . .

RUN go build -o /stock-ticker


EXPOSE 9000

CMD [ "/stock-ticker" ]