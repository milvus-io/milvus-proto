#!/usr/bin/env bash

SCRIPTS_DIR=$(dirname "$0")
CMAKE_DIR="${SCRIPTS_DIR}/../cmake"
BUILD_OUTPUT_DIR="${SCRIPTS_DIR}/../cmake-build"

mkdir -p ${BUILD_OUTPUT_DIR}
pushd ${BUILD_OUTPUT_DIR}
cmake ${CMAKE_DIR}
popd