FROM python:3.10.4-alpine

# Variables
ENV USER_GROUP_SYS contconf
ENV HOME /home/$USER_GROUP_SYS

# Se crea usuario no root.
RUN addgroup -S $USER_GROUP_SYS && adduser -S $USER_GROUP_SYS -G $USER_GROUP_SYS

# Descarga de dependencias para instalar ansible en versión de alpine
RUN apk update && apk add --no-cache musl-dev libffi-dev make gcc

# Se cambia al usuario creado anteriormente
USER $USER_GROUP_SYS

# Se copia el archivo requerido para instalar pip
COPY get-pip.py $HOME

WORKDIR $HOME

# Instalación de pip y de Ansible con pip.
RUN python3 get-pip.py --user && pip install --user cffi && pip install --user ansible

# Se corre proceso principal.
CMD ["tail", "-f", "/dev/null"]