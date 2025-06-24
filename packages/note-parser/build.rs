fn main() -> Result<(), Box<dyn std::error::Error>> {
    tonic_build::compile_protos("../../protobufs/note_service.proto")?;
    return Ok(());
}
