#!/bin/bash

case $1 in
    down)
        for filename in db/migrations/*; do
            if [[ "$filename" == *.down.sql ]]
            then
                FILE="$filename" make db-local-file
            fi
        done;
    ;;
    up)
        for filename in db/migrations/*; do
            if [[ "$filename" == *.up.sql ]]
            then
                FILE="$filename" make db-local-file
            fi
        done;
    ;;
    *)
        echo "choose up or down"
    ;;
esac