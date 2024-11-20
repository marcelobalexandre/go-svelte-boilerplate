FROM node:23.2.0-alpine

ARG NODE_ENV
ENV NODE_ENV=$NODE_ENV

WORKDIR /app

COPY package.json package-lock.json ./
RUN npm install

COPY . .

EXPOSE 3000

CMD ["npm", "run", "dev"]
