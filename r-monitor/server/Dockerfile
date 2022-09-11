FROM node:18.7.0-alpine

WORKDIR /go/src/r-monitor/server
COPY package*.json ./
RUN npm install
COPY . .
EXPOSE 5000
CMD ["npm","run","dev"]