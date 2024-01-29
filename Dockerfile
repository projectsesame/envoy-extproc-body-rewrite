FROM docker.m.daocloud.io/busybox

COPY ./grpc_health_probe /bin/grpc_health_probe
RUN chmod +x /bin/grpc_health_probe

COPY ./envoy-extproc-demo-go /bin/extproc
RUN chmod +x /bin/extproc

ARG EXAMPLE=body-check

EXPOSE 50051

CMD [ "extproc", "body-check", "--log-stream", "--log-phases", "body-size", "32"  ]
