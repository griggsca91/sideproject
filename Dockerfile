FROM golang

ARG app_env
ENV APP_ENV $app_env


COPY ./backend /go/src/github.com/griggsca91/sideproject/app
WORKDIR /go/src/github.com/griggsca91/sideproject/app

RUN pwd
RUN ls

RUN go get ./
RUN go build

CMD go get github.com/pilu/fresh && fresh;

EXPOSE 8080