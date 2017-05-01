FROM node:7.9.0

RUN mkdir web
WORKDIR ./web
ADD . ./

RUN npm install || true
RUN npm install -g @angular/cli || true

VOLUME ./web/dst
EXPOSE 4200
ENTRYPOINT ["ng", "build", "-prod"]
