FROM node:18-alpine AS builder
WORKDIR /app
RUN npm install -g pnpm
COPY package*.json .
RUN pnpm install
COPY . .
RUN pnpm run build
RUN pnpm prune --production

FROM node:18-alpine
WORKDIR /app

COPY --from=builder /app/build build/
COPY --from=builder /app/node_modules node_modules/
COPY --from=builder /app/src/lib/proto/ proto/
COPY package.json .

ENV NODE_ENV=production
ENV PROTO_FILES_LOCATION=proto/main.proto
CMD [ "node", "build" ]