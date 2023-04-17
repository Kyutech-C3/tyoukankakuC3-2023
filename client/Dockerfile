FROM node:16-alpine

RUN mkdir /client

WORKDIR /client

ENV CHOKIDAR_USEPOLLING=true

RUN apk update && \
  apk add --no-cache python3 make g++ && \
  apk add git && \
  npm update -g npm &\
  # npm install -g pnpm &\
  npm update -g pnpm

COPY package.json .
COPY pnpm-lock.yaml .

RUN pnpm install
