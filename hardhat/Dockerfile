FROM node:16.19.0-bullseye
RUN apt update
RUN apt upgrade -y
RUN apt install -y bash curl git
WORKDIR /app
COPY ./package.json /app
COPY ./package-lock.json /app
RUN npm install
COPY . /app