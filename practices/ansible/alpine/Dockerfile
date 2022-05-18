FROM alpine:3.15.4

# Variables
ENV US_GR_ANSIBLE contconf
ENV HOME /home/${US_GR_ANSIBLE}

# Se crea usuario no root.
RUN addgroup -S ${US_GR_ANSIBLE} && adduser -S ${US_GR_ANSIBLE} -G ${US_GR_ANSIBLE}

# Se instala ansible y OpenSSH.
RUN apk update && apk upgrade && apk add --no-cache ansible openssh

# Se define el espacio de trabajo
WORKDIR ${HOME}

# Se cambia al usuario creado anteriormente
USER ${US_GR_ANSIBLE}

# Se define el proceso principal.
CMD ["tail", "-f", "/dev/null"]