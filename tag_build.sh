#!/usr/bin/env bash

# 生成tag
git tag $1

#推送tag
git push origin $1