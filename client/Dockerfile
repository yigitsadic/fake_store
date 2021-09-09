FROM node:16.0.0-alpine3.13 as compiler

WORKDIR /app/src

COPY ./client/package.json package.json
COPY ./client/yarn.lock yarn.lock

RUN yarn install

COPY client client

ENV NODE_ENV=production
RUN cd client && yarn run build

FROM nginx:alpine

COPY --from=compiler /app/src/client/build /usr/share/nginx/html
