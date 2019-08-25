#!/bin/bash

if [ "$1" = "" ]; then
  echo "引数に新しいブランチ名を指定してください."
  exit 1
fi

git checkout master
git pull origin master
git checkout -b $1
