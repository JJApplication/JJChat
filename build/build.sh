#!/usr/bin/env bash

dir=$(pwd);
package=$dir/dist

echo "dir is: $dir"
echo "dist is: $package"

echo "start to build"
export GOPROXY=https://goproxy.io
export GOOS=linux
export GOARCH=amd64

cd $dir/server && go build -o $package/jjchatgpt
cp $dir/server/config.yaml $package
cp -R $dir/ui $package

echo "package done"