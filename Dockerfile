FROM docker.m.daocloud.io/golang:1.21.6-bullseye

SHELL ["/bin/bash", "-c"]

RUN apt-get update && apt-get -y upgrade \
    && apt-get autoremove -y \
    && rm -rf /var/lib/apt/lists/* \
    && apt-get -y clean

WORKDIR /build

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct
COPY . .
RUN go mod tidy \
    && go mod download \
    && go build -o /extproc


FROM docker.m.daocloud.io/busybox

COPY --from=0 /extproc /bin/extproc
RUN chmod +x /bin/extproc

ARG EXAMPLE=body-check

EXPOSE 50051

ENTRYPOINT [ "/bin/extproc" ]
CMD [ "body-check", "--log-stream", "--log-phases", "body-size", "32"  ]
