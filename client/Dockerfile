FROM node:16.0.0-alpine3.13 as compiler

WORKDIR /app/src

COPY client client

RUN cd client

RUN yarn install --production=false

ENV NODE_ENV=production
RUN yarn run build

FROM nginx:alpine

COPY --from=compiler /app/src/client/build /usr/share/nginx/html
