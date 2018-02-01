#!/bin/bash

FOLDER=${PWD##*/}

x=https://github.com/StellarKit/$FOLDER
echo $x


git remote add upstream https://github.com/StellarKit/$FOLDER
git fetch upstream
git checkout master
git merge upstream/master

echo ""
echo "If there are conflicts keep HEAD and delete upstream/master"
echo "run git push origin master"
