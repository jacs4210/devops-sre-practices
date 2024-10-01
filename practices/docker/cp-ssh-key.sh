#!/bin/bash

# Parámetros:
# $1: Ruta de llave privada incluyendo nombre
# $2: Ruta de llave publica incluyendo nombre
# $3: Nombre del contenedor

PT_KEY_PRIV=$1
PT_KEY_PUB=$2
CONTAINER_NAME=$3

STATE_CONTAINER=$(docker ps --format '{{json .}}' | jq '. | select(.Names=="envdev") | .State' | sed 's/\"//g')

# Se verifica que las llaves proporcionadas existan en el sistema y el contenedor está en ejecución.
if [[ -f "$PT_KEY_PRIV" ]] && [[ -f "$PT_KEY_PUB" ]] && [[ "$STATE_CONTAINER" = "running" ]]; then
    echo "[INFO] Copiado de llaves SSH"
    # Copia de llave privada
    docker cp $PT_KEY_PRIV $CONTAINER_NAME:/root/.ssh/

    # Copia de llave publica
    docker cp $PT_KEY_PUB $CONTAINER_NAME:/root/.ssh/

    # Actualización de permisos de las llaves SSH copiadas
    docker exec $CONTAINER_NAME chown root:root -R /root/.ssh

    echo "[INFO] Finalización de copiado de llaves SSH"
else
    echo "[ERROR] Verifique que los parámetros sean correctos"
    exit 1
fi

exit 0