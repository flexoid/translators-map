FROM node:16 as node-builder
WORKDIR /app
COPY web/package*.json ./
RUN npm install
COPY web/ ./
RUN npm run build

FROM golang:1.18-bullseye as go-builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . ./
RUN go vet -v
RUN CGO_ENABLED=0 go build

FROM gcr.io/distroless/static-debian11:debug-nonroot
COPY --chown=nonroot:nonroot --from=node-builder /app/dist ./web
COPY --chown=nonroot:nonroot --from=go-builder /app/translators-map-go ./
EXPOSE 4000
ENTRYPOINT [ "/home/nonroot/translators-map-go" ]
