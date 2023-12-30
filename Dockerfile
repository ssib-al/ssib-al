FROM golang:1.21-alpine AS backend-builder

WORKDIR /build/back
COPY ./back/go.mod ./back/go.sum ./
RUN go mod download
COPY ./back/ .
RUN go build -o main .

FROM node:20-alpine AS frontend-builder

WORKDIR /build/front
COPY ./front/ ./
RUN npm install
RUN npm run build

FROM scratch

COPY --from=backend-builder /build/back/main .
COPY --from=frontend-builder /build/back/public ./public
ENTRYPOINT ["/main"]