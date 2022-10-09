ARG NGINX_VERSION=1.21.6
FROM nginx:${NGINX_VERSION}-alpine

RUN rm -v /etc/nginx/nginx.conf
COPY ./nginx.conf /etc/nginx/

EXPOSE 81
RUN mkdir /nginx
