#!/bin/bash

echo "Ingrese un numero:"
read num

if [ $num -gt 0 ];
then
  echo "$num es positivo"
elif [ $num -lt 0 ];
then
  echo "$num es negativo"
else
  echo "Es cero"
fi
