FROM node:7.9.0

RUN mkdir web
WORKDIR ./web
ADD . ./

RUN npm install

ENTRYPOINT ["ng", "build", "-prod]
VOLUME /web/dist
EXPOSE 4200
