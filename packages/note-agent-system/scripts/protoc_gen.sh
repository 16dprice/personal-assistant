hatch run python -m grpc_tools.protoc \
    --python_out=../src/note_agent_system/generated \
    --pyi_out=../src/note_agent_system/generated \
    --grpc_python_out=../src/note_agent_system/generated \
    --proto_path=../../../protobufs \
    --experimental_allow_proto3_optional \
    ../../../protobufs/*.proto
