## Practicas con Docker

1. Para crear la imagen se puede ejecutar el comando:

```bash
docker build -t <tag> -f Dockerfile .
```

Nota: Debes ubicarte en la ruta donde está el Dockerfile y el tag que quieras darle a tu imagen.

2. Para correr un contenedor con docker utilizando la imagen y que esta se quede en ejecución, se debe ejecutar el siguiente comando:

```bash
docker run -d --tty=1 --name <name_container> <tag_image>
````