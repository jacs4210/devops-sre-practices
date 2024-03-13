FROM python:3.10.4-alpine3.15

# Variables
ARG US_GR_ANSIBLE=contconf
ARG HOME=/home/${US_GR_ANSIBLE}

# Se crea usuario no root.
RUN addgroup -S ${US_GR_ANSIBLE} && adduser -S ${US_GR_ANSIBLE} -G ${US_GR_ANSIBLE}

# Se instala el cliente openssh.
RUN apk update && apk upgrade && apk add --no-cache openssh

# Se define el espacio de trabajo
WORKDIR ${HOME}

# Se expone puerto ssh
EXPOSE 22

# Se cambia al usuario creado anteriormente
USER ${US_GR_ANSIBLE}

# Se define el proceso principal.
CMD ["tail", "-f", "/dev/null"]