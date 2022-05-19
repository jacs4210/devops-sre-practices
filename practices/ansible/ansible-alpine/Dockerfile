FROM alpine:3.15.4

# Variables
ARG US_GR_ANSIBLE=contconf
ARG HOME=/home/${US_GR_ANSIBLE}
ARG SSH_KEY=${HOME}/.ssh/ansible_ssh
ARG PASSPHRASE=dummy

# Se crea usuario no root.
RUN addgroup -S ${US_GR_ANSIBLE} && adduser -S ${US_GR_ANSIBLE} -G ${US_GR_ANSIBLE}

# Se instala ansible y OpenSSH.
RUN apk update && apk upgrade && apk add --no-cache ansible openssh

# Se define el espacio de trabajo
WORKDIR ${HOME}

# Se cambia al usuario creado anteriormente
USER ${US_GR_ANSIBLE}

# Se crea llave SSH
RUN ssh-keygen -q -f ${SSH_KEY} -P ${PASSPHRASE}

# Se define el proceso principal.
CMD ["tail", "-f", "/dev/null"]