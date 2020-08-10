# FROM ubuntu 
FROM nginx
# RUN mkdir /app
# ADD . /app
# RUN apt-get update
# RUN apt-get install -y nginx
ADD index.html /usr/share/nginx/html/

RUN chmod +r /usr/share/nginx/html/index.html
CMD ["nginx", "-g", "daemon off"]
# EXPOSE 80
# ENTRYPOINT [ "/usr/share/nginx", "-g", "daemon off;" ]


# FROM ubuntu 

# RUN apt-get update

# RUN apt-get -y install httpd

# COPY index.html /usr/www/html

# EXPOSE 80

# CMD apachectl -D FORGROUND
