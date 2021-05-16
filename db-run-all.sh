#!/bin/bash

case $1 in
    down)
        for filename in db/migrations/*; do
            echo "$filename"
            if [[ "$filename" == *.down.sql ]]
            then
                FILE="$filename" make db-run-file-local
            fi
        done;
    ;;
    up)
        for filename in db/migrations/*; do
            echo "$filename"
            if [[ "$filename" == *.up.sql ]]
            then
                FILE="$filename" make db-run-file-local
            fi
        done;
    ;;
    *)
        echo "choose up or down"
    ;;
esac