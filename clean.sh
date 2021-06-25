#!/bin/bash

for tag in $(git tag -l); do
  echo $tag
done
