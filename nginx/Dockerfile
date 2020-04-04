FROM nginx:latest

RUN apt-get update && apt-get install -y \
    procps \
    strace

COPY nginx.conf /etc/nginx/nginx.conf
COPY conf.d /etc/nginx/conf.d