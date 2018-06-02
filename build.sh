#!/bin/bash
rm -rf build
mkdir -p build

cd cli
make
cp bin/key-value ../build
cd ..

cp -r cfn build

zip deployment-artifact build/*
