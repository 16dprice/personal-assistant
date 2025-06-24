from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from collections.abc import Iterable as _Iterable, Mapping as _Mapping
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Note(_message.Message):
    __slots__ = ("title", "content", "tags", "created_at", "updated_at", "linked_notes")
    TITLE_FIELD_NUMBER: _ClassVar[int]
    CONTENT_FIELD_NUMBER: _ClassVar[int]
    TAGS_FIELD_NUMBER: _ClassVar[int]
    CREATED_AT_FIELD_NUMBER: _ClassVar[int]
    UPDATED_AT_FIELD_NUMBER: _ClassVar[int]
    LINKED_NOTES_FIELD_NUMBER: _ClassVar[int]
    title: str
    content: str
    tags: _containers.RepeatedScalarFieldContainer[str]
    created_at: int
    updated_at: int
    linked_notes: _containers.RepeatedCompositeFieldContainer[Note]
    def __init__(
        self,
        title: _Optional[str] = ...,
        content: _Optional[str] = ...,
        tags: _Optional[_Iterable[str]] = ...,
        created_at: _Optional[int] = ...,
        updated_at: _Optional[int] = ...,
        linked_notes: _Optional[_Iterable[_Union[Note, _Mapping]]] = ...,
    ) -> None: ...

class CreateNoteRequest(_message.Message):
    __slots__ = ("title", "content", "tags", "linked_note_titles")
    TITLE_FIELD_NUMBER: _ClassVar[int]
    CONTENT_FIELD_NUMBER: _ClassVar[int]
    TAGS_FIELD_NUMBER: _ClassVar[int]
    LINKED_NOTE_TITLES_FIELD_NUMBER: _ClassVar[int]
    title: str
    content: str
    tags: _containers.RepeatedScalarFieldContainer[str]
    linked_note_titles: _containers.RepeatedScalarFieldContainer[str]
    def __init__(
        self,
        title: _Optional[str] = ...,
        content: _Optional[str] = ...,
        tags: _Optional[_Iterable[str]] = ...,
        linked_note_titles: _Optional[_Iterable[str]] = ...,
    ) -> None: ...

class CreateNoteResponse(_message.Message):
    __slots__ = ("note",)
    NOTE_FIELD_NUMBER: _ClassVar[int]
    note: Note
    def __init__(self, note: _Optional[_Union[Note, _Mapping]] = ...) -> None: ...

class GetNoteRequest(_message.Message):
    __slots__ = ("title",)
    TITLE_FIELD_NUMBER: _ClassVar[int]
    title: str
    def __init__(self, title: _Optional[str] = ...) -> None: ...

class GetNoteResponse(_message.Message):
    __slots__ = ("note",)
    NOTE_FIELD_NUMBER: _ClassVar[int]
    note: Note
    def __init__(self, note: _Optional[_Union[Note, _Mapping]] = ...) -> None: ...

class ListNotesRequest(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class ListNotesResponse(_message.Message):
    __slots__ = ("notes",)
    NOTES_FIELD_NUMBER: _ClassVar[int]
    notes: _containers.RepeatedCompositeFieldContainer[Note]
    def __init__(
        self, notes: _Optional[_Iterable[_Union[Note, _Mapping]]] = ...
    ) -> None: ...

class UpdateNoteRequest(_message.Message):
    __slots__ = ("title", "content", "tags", "linked_note_titles")
    TITLE_FIELD_NUMBER: _ClassVar[int]
    CONTENT_FIELD_NUMBER: _ClassVar[int]
    TAGS_FIELD_NUMBER: _ClassVar[int]
    LINKED_NOTE_TITLES_FIELD_NUMBER: _ClassVar[int]
    title: str
    content: str
    tags: _containers.RepeatedScalarFieldContainer[str]
    linked_note_titles: _containers.RepeatedScalarFieldContainer[str]
    def __init__(
        self,
        title: _Optional[str] = ...,
        content: _Optional[str] = ...,
        tags: _Optional[_Iterable[str]] = ...,
        linked_note_titles: _Optional[_Iterable[str]] = ...,
    ) -> None: ...

class UpdateNoteResponse(_message.Message):
    __slots__ = ("note",)
    NOTE_FIELD_NUMBER: _ClassVar[int]
    note: Note
    def __init__(self, note: _Optional[_Union[Note, _Mapping]] = ...) -> None: ...

class DeleteNoteRequest(_message.Message):
    __slots__ = ("title",)
    TITLE_FIELD_NUMBER: _ClassVar[int]
    title: str
    def __init__(self, title: _Optional[str] = ...) -> None: ...

class DeleteNoteResponse(_message.Message):
    __slots__ = ("success",)
    SUCCESS_FIELD_NUMBER: _ClassVar[int]
    success: bool
    def __init__(self, success: bool = ...) -> None: ...
