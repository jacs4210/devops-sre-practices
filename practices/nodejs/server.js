// Modulos a importar para utilizar.
const { createServer } = require ('node:http');

// Puerto del servidor web
const port = 3000;

// Dominio del servidor web
const hostname = '127.0.0.1'

// Un servidor web tendrá 2 objetos importantes que son:
// 1. request: que representa la petición http que se está realizando.
// 2. response: que representa la respuesta http que se va a enviar al cliente que hizo la petición.
const server = createServer((req, res) => {
    // Creación de la cabecera de la respuesta http. Se define un codigo de la respuesta y finalmente cabeceras u opciones de la respuesta.
    res.writeHead(200, { 'Content-Type': 'text/plain' });

    // Finaliza la petición http y responde un mensaje que será convertido en el tipo de archivo mencionado en la cabecera de la respuesta.
    res.end('Hello World\n');
});

// Cuando se crea el servidor, se deberá configurar el puerto por donde va a escuchar las peticiones y un callback si se quiere para dar por enterado de que el servidor
// esta disponible.
server.listen(port, hostname, () => {
    console.log(`Server running at http://${hostname}:${port}/`);
});
