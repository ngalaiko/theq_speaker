FROM node:7.9.0

RUN mkdir web
WORKDIR ./web
ADD . ./

RUN npm install || true
RUN npm install -g @angular/cli || true

ENTRYPOINT ["ng", "serve", "-prod", "--host", "galayko.rocks"]
EXPOSE 4200
