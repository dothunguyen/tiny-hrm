FROM golang:1.13.4
# install glide
RUN go get github.com/Masterminds/glide

RUN curl -sL https://deb.nodesource.com/setup_8.x | bash \
    && apt-get install -y nodejs \
    && apt-get install -y npm
# create a working directory
WORKDIR /go/src/tiny-hrm
# add glide.yaml and glide.lock
ADD glide.yaml glide.yaml
ADD glide.lock glide.lock
# install packages
RUN glide install
# add source code
ADD src src
# build web ui
WORKDIR /go/src/tiny-hrm/src/views

RUN npm install \
    && npm run build
# run main.go
WORKDIR /go/src/tiny-hrm
RUN mkdir .run
CMD ["go", "run", "src/main.go"]