FROM node:alpine as build

WORKDIR /app

COPY package*.json ./
RUN npm install

COPY . . 
RUN npm run build

FROM nginx:latest

WORKDIR /

COPY --from=build ./app/dist /usr/share/nginx/html
EXPOSE 3000