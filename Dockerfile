FROM node:22.15.0-alpine3.21
WORKDIR /opt/app

COPY . .
RUN npm install

EXPOSE 3000

CMD ["node", "index.js"]