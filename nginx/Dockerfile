FROM nginx:latest

RUN apt-get update -y
RUN apt-get upgrade -y
RUN apt-get install libssl-dev -y

RUN mkdir -p /etc/nginx/ssl/
RUN openssl genrsa -des3 -passout pass:NotSecure -out /etc/nginx/ssl/cert.pass.key 2048 \
    && openssl rsa -passin pass:NotSecure -in /etc/nginx/ssl/cert.pass.key -out /etc/nginx/ssl/cert.key \
    && rm /etc/nginx/ssl/cert.pass.key \
    && openssl req -new -passout pass:NotSecure -key /etc/nginx/ssl/cert.key -out /etc/nginx/ssl/cert.csr \
    -subj '/C=SS/ST=SS/L=Gotham City/O=API Platform Dev/CN=localhost' \
    && openssl x509 -req -sha256 -days 365 -in /etc/nginx/ssl/cert.csr -signkey /etc/nginx/ssl/cert.key -out /etc/nginx/ssl/cert.crt
