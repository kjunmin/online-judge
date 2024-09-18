FROM node:alpine AS build

WORKDIR /app

COPY package*.json yarn.lock ./
RUN yarn

COPY . . 
RUN yarn build

FROM nginx:alpine

RUN rm /etc/nginx/conf.d/default.conf
COPY ./nginx.conf /etc/nginx/conf.d/
COPY --from=build ./app/dist /usr/share/nginx/html
EXPOSE 3000
CMD [ "nginx", "-g", "daemon off;" ]
