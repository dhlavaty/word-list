FROM alpine:3.13.1 as buildimg

RUN wget https://github.com/dhlavaty/word-list/releases/download/v0.0.1/word-list-cli-0.0.1-linux-amd64.tar.gz -O tmp.tar.gz
RUN tar x -f tmp.tar.gz
RUN chmod +x /word-list-cli


FROM scratch
WORKDIR /usr/local/bin
COPY --from=buildimg /word-list-cli .
CMD ["/usr/local/bin/word-list-cli"]
