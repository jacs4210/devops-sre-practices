FROM node:11.1.0-alpine

WORKDIR /app

ADD package.json package-lock.json /app/
RUN npm install

EXPOSE 3000

ADD . /app

# Al utilizar node como comando, permite que el proceso principal se le asigne al comando node, mientras que si utilizo npm start, el PID 1 estará es directamente a npm y no a node.
CMD ["node", "index"]
