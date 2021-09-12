FROM node:16.0.0-alpine3.13 as compiler

ARG TARGET_API_URL

WORKDIR /app/src

COPY client .

RUN yarn install --production=false

ENV NODE_ENV=production

ENV API_URL=$TARGET_API_URL
RUN yarn run build

FROM nginx:alpine

COPY --from=compiler /app/src/build /usr/share/nginx/html
