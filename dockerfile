FROM node:latest

WORKDIR /app

COPY packege*.json ./

RUN npm install

COPY . .

ENV post=8080

EXPOSE 8080
