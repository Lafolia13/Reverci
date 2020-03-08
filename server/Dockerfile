FROM gcc:latest AS gcc
COPY solvers/ /solvers/
WORKDIR /solvers/
RUN mkdir /solvers/cpu/ && \
  make OUTDIR="/solvers/cpu/" && \
  ls /solvers/cpu

FROM golang:latest
COPY --from=gcc /solvers/cpu/ /solvers/
COPY golang/ /server/
WORKDIR /server/
ENV GOPATH="/server/"
RUN go get -u github.com/gorilla/mux && \
  go install $GOPATH/src/cmd/app/

CMD ["/server/bin/app"]
