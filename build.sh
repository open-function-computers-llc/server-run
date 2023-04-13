#! /bin/bash

# kill any currently running version of the app
killall server-run

cd frontend
npm install
ng build
cd ..

# build and run the new binary
if [[ "$OSTYPE" == "linux-gnu" ]]; then
    env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-X main.Version=`git rev-parse HEAD`" -o dist/server-run
    dist/server-run serve &
elif [[ "$OSTYPE" == "darwin"* ]]; then
    env GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-X main.Version=`git rev-parse HEAD`" -o dist/server-run
    dist/server-run serve &
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
