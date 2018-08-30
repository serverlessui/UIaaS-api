#!/bin/bash

ENVIRONMENT=$1

deleteStacks () {
  cd stacks && make delete ENVIRONMENT=$ENVIRONMENT
}

deleteAccount() {
    cd account && make delete ENVIRONMENT=$ENVIRONMENT
}

deleteStats() {
    cd stats && make delete ENVIRONMENT=$ENVIRONMENT
}

deleteFunctions() {
    deleteStacks
    cd ..
    deleteAccount
    cd ..
    deleteStats
}