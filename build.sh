#!/bin/bash
rm -rf build
mkdir -p build

cd cli
make
cp bin/key-value ../build
cd ..

cp -f buildspec.yml build
cp -Rf cfn build

zip -r deployment-artifact -j build/*
