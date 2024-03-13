FROM python:3.12.0-alpine3.18

# Variables
ARG US_GR_ANSIBLE=contconf
ARG HOME=/home/${US_GR_ANSIBLE}

# Se crea usuario no root.
RUN addgroup -S ${US_GR_ANSIBLE} && adduser -S ${US_GR_ANSIBLE} -G ${US_GR_ANSIBLE}

# Descarga de dependencias para instalar ansible en versión de alpine
RUN apk update && apk add --no-cache ansible openssh

# Se define el espacio de trabajo
WORKDIR ${HOME}

# Se cambia al usuario creado anteriormente
USER ${US_GR_ANSIBLE}

# Se corre proceso principal.
CMD ["tail", "-f", "/dev/null"]