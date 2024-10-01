FROM ubuntu:18.04

# Argumentos utilizados para configurar git.
ARG USER_GIT
ARG EMAIL_GIT

# Descarga de herramientas de desarrollo
RUN apt-get update && apt-get install git -y

# Configuración global de git
RUN git config --global user.name ${USER_GIT} && \
    git config --global user.email ${EMAIL_GIT}

# Creación de directorio para almacenar llaves SSH
RUN mkdir -p ${HOME}/.ssh

CMD [ "/bin/bash" ]