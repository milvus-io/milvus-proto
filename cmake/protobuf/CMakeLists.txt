set( PROTOBUF_SOURCE_URL
     "https://github.com/protocolbuffers/protobuf/archive/refs/tags/v3.21.4.zip" )

message( STATUS "Building protobuf-cpp from source" )

FetchContent_Declare( protobuf
    URL             ${PROTOBUF_SOURCE_URL}
    URL_MD5         "4f524caf15beb4354d64b63901587932"
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
