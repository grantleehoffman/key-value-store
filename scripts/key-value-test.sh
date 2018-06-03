#!/bin/bash

script_exit=0
./key-value put --key test --value testvalue -s ${server_url}
if [ $? != 0 ];then
  echo '=== ./key-value put --key test --value testvalue exit code is not 0 x ==='
  script_exit=1
else
  echo '=== ./key-value put --key test --value testvalue exit code 0 ✓ ==='
fi
echo ""

output=$(./key-value get --key test -s ${server_url})
if [ $? != 0 ] || [ ${output} != "testvalue" ]; then
  echo '=== ./key-value get --key test output is not "testvalue" x ==='
  script_exit=1
else
  echo '=== ./key-value get --key test output is "testvalue" ✓ ==='
fi
echo ""

./key-value delete --key test -s ${server_url} &> /dev/null
if [ $? != 0 ];then
  echo '=== ./key-value delete --key test exit code is not 0 x ==='
  script_exit=1
else
  echo '=== ./key-value delete --key test exit code 0 ✓ ==='
fi
echo ""

./key-value get --key test -s ${server_url} &> /dev/null
if [ $? != 1 ];then
  echo '=== ./key-value get --key test exit code is not 1 x ==='
  script_exit=1
else
  echo '=== ./key-value get --key test exit code 1 ✓ ==='
fi
echo ""

if [[ ${script_exit} != 0 ]];then
  echo "Test failure"
  exit 1
fi

echo "All Tests Passed"
