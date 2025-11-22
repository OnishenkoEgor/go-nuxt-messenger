FROM node:22-alpine AS build

WORKDIR /app

COPY package.json package-lock.json nuxt.config.ts tsconfig.json /app/ /public/ ./

RUN npm install

RUN npm run build

ENV HOST=0.0.0.0 NODE_ENV=production
ENV NODE_ENV=production

EXPOSE 3000

CMD ["node", ".output/server/index.mjs"]
