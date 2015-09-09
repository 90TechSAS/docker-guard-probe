#/bin/bash

rm -rf dgp
cp -r .. dgp
docker build -t dg-probe .
rm -rf dgp
