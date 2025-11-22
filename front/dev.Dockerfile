FROM node:22-alpine AS build

WORKDIR /app

ENV HOST=0.0.0.0 NODE_ENV=production
ENV NODE_ENV=production

EXPOSE 3000
