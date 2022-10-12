set( PROTOBUF_SOURCE_URL
     "https://github.com/protocolbuffers/protobuf/releases/download/v3.9.0/protobuf-cpp-3.9.0.zip" )

message( STATUS "Building protobuf-cpp-3.9.0 from source" )

FetchContent_Declare( protobuf
    URL             ${PROTOBUF_SOURCE_URL}
    URL_MD5         "9562b27cc6ac5ebd087f201f1310c885"
    DOWNLOAD_DIR    ${THIRDPARTY_DOWNLOAD_PATH}
    SOURCE_DIR      ${CMAKE_CURRENT_BINARY_DIR}/protobuf-src
    BINARY_DIR      ${CMAKE_CURRENT_BINARY_DIR}/protobuf-build
)

FetchContent_MakeAvailable(protobuf)