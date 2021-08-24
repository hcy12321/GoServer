FROM golang:1.15.10 

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
ENV PROJECT_NAME=GoServer
COPY . /${PROJECT_NAME}
RUN cd /${PROJECT_NAME} && go build main.go
EXPOSE 8101
CMD cd /${PROJECT_NAME} && ./main