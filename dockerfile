FROM golang:1.8

RUN mkdir -p /app
# RUN mkdir -p /app/data
RUN apt-get update && apt-get install -y --no-install-recommends \
		sqlite3

VOLUME ["/app/data"]

WORKDIR /app

ADD ./server /app
ADD ./cli/cli /app

EXPOSE 8080

CMD ["/app/server"]

