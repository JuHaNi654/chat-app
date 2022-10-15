FROM node:16-buster

WORKDIR /app

COPY package*.json ./ 

RUN npm install

COPY . .

EXPOSE 5173

CMD ["npm", "run", "dev"]

