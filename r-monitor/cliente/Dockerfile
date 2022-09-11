FROM node:18.7.0-alpine

WORKDIR /go/src/frontend
COPY package*.json ./
RUN npm install
COPY . .
EXPOSE 3000
CMD ["npm","start"]