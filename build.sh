#!/bin/bash

buildStacks () {
  cd stacks && make
}

buildAccount() {
    cd account && make
}

buildStats() {
    cd stats && make
}

buildFunctions() {
    buildStacks
    cd ..
    buildAccount
    cd ..
    buildStats
}

buildFunctions
