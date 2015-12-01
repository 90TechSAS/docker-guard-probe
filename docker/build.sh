#/bin/bash

rm -rf dgp
mkdir dgp
cp -r ../config.yaml dgp
cp -r ../main.go dgp
cp -r ../core dgp
cp -r ../utils dgp
docker build -t dg-probe .
rm -rf dgp
