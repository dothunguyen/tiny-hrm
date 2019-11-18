#!/bin/bash
echo "Build web ui..."
cd views

npm install && npm run build
echo "Done!"

cd ../

echo "Build API and create executable..."
echo "Install dependencies..."
echo "Gin framework..."

go get github.com/gin-gonic/gin
github.com/gin-gonic/contrib/static

echo "Done!"

echo "SQLite3 for go..."

go get github.com/mattn/go-sqlite3

echo "Done!"
echo "Now build the executable..."

go build

echo "Done!! now you can run the server with command 'bash run.sh'"