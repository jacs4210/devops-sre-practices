FROM python:3.10.4-alpine

# Variables
ENV US_GR_ANSIBLE contconf
ENV HOME /home/${US_GR_ANSIBLE}
ENV ROOT root

# Se crea usuario no root.
RUN addgroup -S ${US_GR_ANSIBLE} && adduser -S ${US_GR_ANSIBLE} -G ${US_GR_ANSIBLE}

# Descarga de dependencias para instalar ansible en versión de alpine
RUN apk update && apk add --no-cache musl-dev libffi-dev make gcc openssh

# Se define el espacio de trabajo
WORKDIR ${HOME}

# Se copia el archivo requerido para instalar pip
COPY get-pip.py .

# Se dan permisos sobre el archivo para instalar pip
RUN chown ${US_GR_ANSIBLE}:${US_GR_ANSIBLE} get-pip.py

# Se cambia al usuario creado anteriormente
USER ${US_GR_ANSIBLE}

# Instalación de pip y de Ansible con pip.
RUN python3 get-pip.py --user && pip install --user cffi && pip install --user ansible

USER ${ROOT}

# Se crea enlace simbolico para la ejecución de ansible.
RUN cd /usr/bin && ln -s ~/.local/bin/ansible ansible

# Se cambia al usuario no-root del sistema.
USER ${US_GR_ANSIBLE}

# Se corre proceso principal.
CMD ["tail", "-f", "/dev/null"]