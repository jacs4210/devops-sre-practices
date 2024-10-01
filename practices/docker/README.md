## Practicas con Docker

### Caracteristicas tecnicas utilizadas para este repositorio
- Docker version: 27.2.0, build 3ab4256
- Docker compose version: v2.29.2-desktop.2

### Utilizando solo Docker (1-only-docker)

1. Para crear la imagen se puede ejecutar el comando:

```bash
docker build -t <tag> -f Dockerfile .
```

Nota: Debes ubicarte en la ruta donde está el Dockerfile y el tag que quieras darle a tu imagen.

2. Para correr un contenedor con docker utilizando la imagen y que esta se quede en ejecución, se debe ejecutar el siguiente comando:

```bash
docker run -d --tty=1 --build-arg USER_GIT="<Name user>" --build-arg EMAIL_GIT="<email user>" --name <name_container> <tag_image>
```

Nota: Se debe ubicar el valor de los argumentos mencionados con el fin de que se realice correctamente la configuración de git.

### Utilizando docker-compose (2-docker-compose)
Otra alternativa es utilizar el archivo "docker-compose.yml" del directorio **2-docker-compose** y crear el archivo *.env* con las siguientes variables:

- USER_GIT="<name_user_git>"
- EMAIL_GIT="<email>"

Por ejemplo:

- USER_GIT="Roronoa Zoro"
- EMAIL_GIT="kitetsu_zoro@gmail.com"

Una vez definido el archivo *.env* se puede ejecutar el siguiente comando para la creación del entorno de desarrollo.

```bash
docker compose up -d
```