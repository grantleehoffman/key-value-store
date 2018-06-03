#!/bin/bash
rm -rf build
mkdir -p build

# Build CLI
cd cli
make
cd ..

# Lint CFN Templates
for i in $(ls cfn); do
  cat cfn/$i | python -m json.tool > /dev/null
  if [ $? != 0 ]; then
    echo "bad json found in file cfn/$i"
    exit 1
  fi
done

# Build deployment artifact
cp -f cli/bin/linux/key-value build
cp -f scripts/key-value-test.sh build
cp -f buildspec.yml build
cp -Rf cfn build
zip -r deployment-artifact -j build/*
