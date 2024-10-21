#!/bin/bash

#set -x

echo "Ingrese el nombre a quien quieres que salude:"
read name

if [[ -n "$name" ]];
then
  echo "Hello, $name"
else
  echo "Por favor ingrese el nombre de la persona que quieres que salude."
  exit 1
fi


