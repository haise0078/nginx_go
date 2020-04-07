FROM nginx:latest

RUN apt-get update && apt-get install -y \
    procps \
    strace
RUN mkdir \
    /var/www \
    /var/www/public 

COPY nginx.conf /etc/nginx/nginx.conf
COPY conf.d /etc/nginx/conf.d
COPY public /var/www/public