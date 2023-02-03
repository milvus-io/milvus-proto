# Brief
Share Milvus proto files across SDK repositories.
SDK list for Milvus:
- [pymilvus](https://github.com/milvus-io/pymilvus)
- [Java SDK](https://github.com/milvus-io/milvus-sdk-java)
- [C++ SDK](https://github.com/milvus-io/milvus-sdk-cpp)
- [Go SDK](https://github.com/milvus-io/milvus-sdk-go)
- [Node SDK](https://github.com/milvus-io/milvus-sdk-node)
- [C# SDK](https://github.com/milvus-io/milvus-sdk-csharp)


# Usage
Each SDK project use milvus-proto as a git submodule. The milvus-proto only shares the proto files, each SDK project compiles the proto files by itself.

## First time to add milvus-proto as a submodule

```git submodule add https://github.com/milvus-io/milvus-proto.git [target path]```

"target path" is the folder to add the submodule 

## Update exist submodule

```git submodule update --init```

## Develop SDK on a particular version
The milvus-proto has only one branch "master", we have created tags for official version on the branch: v2.0.0, v2.0.1, etc.
To develop SDK on a particular version, checkout the tag/commit and submit a commit to the SDK project:

1. go into milvus-proto submodule folder
   
    ```cd milvus-proto```

2. checkout a particular tag/commit
   
    ```git checkout v2.0.1```

3. go to SDK project folder

    ```cd ..```

4. submit a commit to the SDK project
5. develop SDK on current proto version

## Sync proto files from Milvus main project
To avoid complex workflow on Milvus main project, currently we don't introduce milvus-proto as submodule to the main project.
When we implement a new feature with new API, the following steps are recommended:

1. add new API to the proto files of the main project, submit a commit
2. implement the new API make sure the API can be invoked
3. sync the modification to the milvus-proto, submit a commit
4. in SDK project, sync the milvus-proto submodule to the commit and submit a commit
5. develop SDK and main project parallelly

## License
Milvus proto repo is part of LF AI & Data Milvus project under Apache License 2.0.


