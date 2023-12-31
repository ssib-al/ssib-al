FROM golang:1.21-alpine AS backend-builder

WORKDIR /build/back
COPY ./back/go.mod ./back/go.sum ./
RUN go mod download
COPY ./back/ .
RUN go build -o main .

FROM node:20-alpine AS frontend-builder

ARG MODE=production
ARG DOMAIN_PREFIX

WORKDIR /build/front

RUN echo "VITE_DOMAIN_PREFIX=${DOMAIN_PREFIX}" > .env
RUN echo "VITE_MODE=${MODE}" >> .env

COPY ./front/package.json ./front/package-lock.json ./
RUN npm install

COPY ./front/ ./

RUN npm run build

FROM alpine:3.18

# Install Doppler CLI
RUN wget -q -t3 'https://packages.doppler.com/public/cli/rsa.8004D9FF50437357.key' -O /etc/apk/keys/cli@doppler-8004D9FF50437357.rsa.pub && \
    echo 'https://packages.doppler.com/public/cli/alpine/any-version/main' | tee -a /etc/apk/repositories && \
    apk add doppler

COPY --from=backend-builder /build/back/main .
COPY --from=frontend-builder /build/back/public ./public

CMD ["doppler", "run", "--", "/main"]