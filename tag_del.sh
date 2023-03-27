#!/usr/bin/env bash

# 删除本地tag
git tag -d $1

# 推送远程
git push origin :refs/tags/$1

