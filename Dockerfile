FROM golang

ARG app_env
ENV APP_ENV $app_env

COPY ./service /go/src/github.com/efrainmunoz/go-microservice-template/service
WORKDIR /go/src/github.com/efrainmunoz/go-microservice-template/service

RUN go get ./
RUN go build

CMD if [ ${APP_ENV} = production ]; \
	then \
	service; \
	else \
	go get github.com/pilu/fresh && \
	fresh; \
	fi
	
EXPOSE 8080