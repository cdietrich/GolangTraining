#!/bin/bash
docker run -it -u `id -u $USER` -p 3000:3000 -v "$(pwd):/home/project:cached" theiaide/theia-go
