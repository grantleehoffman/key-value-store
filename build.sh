#!/bin/bash
rm -rf build
mkdir -p build

cd cli
make
cp bin/linux/key-value ../build
cd ..

cp -f scripts/key-value-test.sh build
cp -f buildspec.yml build
cp -Rf cfn build

zip -r deployment-artifact -j build/*
