cmake_minimum_required( VERSION 3.18 )

project(proto)
include(FetchContent)

message( STATUS "source dir ${PROJECT_SOURCE_DIR}" )

set( THIRDPARTY_DOWNLOAD_PATH ${PROJECT_SOURCE_DIR}/3rdparty_download/download )

add_subdirectory( protobuf )

add_custom_target( Clean-All COMMAND ${CMAKE_BUILD_TOOL} clean )