# pull official base image
FROM node:14.12.0-alpine
RUN mkdir /app
RUN mkdir /app/log
WORKDIR /app

# install app dependencies
ADD package.json /app/package.json
ADD package-lock.json /app/package-lock.json
ADD src /app/src
ADD public /app/public

RUN npm install
RUN npm install react-script

CMD ["npm", "start"]