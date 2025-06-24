from note_agent_system.generated.note_service_pb2 import Note, GetNoteRequest
from note_agent_system.generated.note_service_pb2_grpc import NoteServiceStub
import grpc

channel = grpc.insecure_channel("localhost:50051")
client = NoteServiceStub(channel=channel)

request = GetNoteRequest(title="1740771336 We All Write.md")
response = client.GetNote(request)

print(response)
