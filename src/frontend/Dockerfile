FROM node:20-alpine

WORKDIR /src/frontend
COPY package.json .

RUN npm install

COPY . .
RUN npm run build

FROM nginx
EXPOSE 80
COPY --from=0 /src/frontend/dist /usr/share/nginx/html