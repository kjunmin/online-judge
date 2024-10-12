FROM node:22-slim AS build

WORKDIR /app

COPY . .
RUN npm install
RUN npm run build

FROM node:22-slim

COPY --from=build /app/build /build

EXPOSE 3000

CMD ["node", "build"]
