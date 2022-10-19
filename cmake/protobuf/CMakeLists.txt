set( PROTOBUF_SOURCE_URL
     "https://github.com/protocolbuffers/protobuf/releases/download/v3.9.0/protobuf-cpp-3.9.0.zip" )

message( STATUS "Building protobuf-cpp-3.9.0 from source" )

FetchContent_Declare( protobuf
    URL             ${PROTOBUF_SOURCE_URL}
    URL_MD5         "9562b27cc6ac5ebd087f201f1310c885"
    DOWNLOAD_DIR    ${CMAKE_CURRENT_BINARY_DIR}/../3rdparty_download
    SOURCE_DIR      ${CMAKE_CURRENT_BINARY_DIR}/protobuf-src
    BINARY_DIR      ${CMAKE_CURRENT_BINARY_DIR}/protobuf-build
)

set( protobuf_BUILD_TESTS CACHE BOOL OFF FORCE )
set( protobuf_WITH_ZLIB CACHE BOOL OFF FORCE )
if ( NOT protobuf_POPULATED )
    FetchContent_Populate( protobuf )

    # Adding the following targets:
    # protobuf::libprotobuf - static target
    # protobuf::protoc - executable target
    add_subdirectory( ${protobuf_SOURCE_DIR}/cmake
                      ${protobuf_BINARY_DIR}
                      EXCLUDE_FROM_ALL )
endif()