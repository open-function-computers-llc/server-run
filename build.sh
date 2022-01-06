#! /bin/bash

# kill any currently running version of the app
killall server-run

# cd frontend
# # npm install
# ng build
# cd ..

# build and run the new binary
if [[ "$OSTYPE" == "linux-gnu" ]]; then
    env GOOS=linux GOARCH=amd64 go build -o dist/server-run
    dist/server-run &
elif [[ "$OSTYPE" == "darwin"* ]]; then
    env GOOS=darwin GOARCH=amd64 go build -o dist/server-run
    dist/server-run &
elif [[ "$OSTYPE" == "cygwin" ]]; then
    echo "windows?"
elif [[ "$OSTYPE" == "msys" ]]; then
    echo "windows?"
elif [[ "$OSTYPE" == "win32" ]]; then
    echo "windows?"
elif [[ "$OSTYPE" == "freebsd"* ]]; then
    echo "free-bsd?"
else
    echo "i have no idea what's going on"
fi
