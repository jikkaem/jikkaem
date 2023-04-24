from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf import empty_pb2 as _empty_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor
FEMALE: Gender
MALE: Gender

class ArtistObject(_message.Message):
    __slots__ = ["birthplace", "country", "dob", "full_name", "gender", "group", "height", "id", "instagram", "korean_name", "korean_stage_name", "stage_name", "weight"]
    BIRTHPLACE_FIELD_NUMBER: _ClassVar[int]
    COUNTRY_FIELD_NUMBER: _ClassVar[int]
    DOB_FIELD_NUMBER: _ClassVar[int]
    FULL_NAME_FIELD_NUMBER: _ClassVar[int]
    GENDER_FIELD_NUMBER: _ClassVar[int]
    GROUP_FIELD_NUMBER: _ClassVar[int]
    HEIGHT_FIELD_NUMBER: _ClassVar[int]
    ID_FIELD_NUMBER: _ClassVar[int]
    INSTAGRAM_FIELD_NUMBER: _ClassVar[int]
    KOREAN_NAME_FIELD_NUMBER: _ClassVar[int]
    KOREAN_STAGE_NAME_FIELD_NUMBER: _ClassVar[int]
    STAGE_NAME_FIELD_NUMBER: _ClassVar[int]
    WEIGHT_FIELD_NUMBER: _ClassVar[int]
    birthplace: str
    country: str
    dob: _timestamp_pb2.Timestamp
    full_name: str
    gender: Gender
    group: str
    height: int
    id: str
    instagram: str
    korean_name: str
    korean_stage_name: str
    stage_name: str
    weight: int
    def __init__(self, id: _Optional[str] = ..., stage_name: _Optional[str] = ..., full_name: _Optional[str] = ..., korean_name: _Optional[str] = ..., korean_stage_name: _Optional[str] = ..., dob: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., group: _Optional[str] = ..., country: _Optional[str] = ..., height: _Optional[int] = ..., weight: _Optional[int] = ..., birthplace: _Optional[str] = ..., gender: _Optional[_Union[Gender, str]] = ..., instagram: _Optional[str] = ...) -> None: ...

class DeleteFancamRequest(_message.Message):
    __slots__ = ["id"]
    ID_FIELD_NUMBER: _ClassVar[int]
    id: str
    def __init__(self, id: _Optional[str] = ...) -> None: ...

class FancamList(_message.Message):
    __slots__ = ["fancams"]
    FANCAMS_FIELD_NUMBER: _ClassVar[int]
    fancams: _containers.RepeatedCompositeFieldContainer[FancamObject]
    def __init__(self, fancams: _Optional[_Iterable[_Union[FancamObject, _Mapping]]] = ...) -> None: ...

class FancamObject(_message.Message):
    __slots__ = ["channelId", "channelTitle", "description", "id", "publishedAt", "record_date", "rootThumbnail", "suggested_tags", "title"]
    CHANNELID_FIELD_NUMBER: _ClassVar[int]
    CHANNELTITLE_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    ID_FIELD_NUMBER: _ClassVar[int]
    PUBLISHEDAT_FIELD_NUMBER: _ClassVar[int]
    RECORD_DATE_FIELD_NUMBER: _ClassVar[int]
    ROOTTHUMBNAIL_FIELD_NUMBER: _ClassVar[int]
    SUGGESTED_TAGS_FIELD_NUMBER: _ClassVar[int]
    TITLE_FIELD_NUMBER: _ClassVar[int]
    channelId: str
    channelTitle: str
    description: str
    id: str
    publishedAt: _timestamp_pb2.Timestamp
    record_date: _timestamp_pb2.Timestamp
    rootThumbnail: str
    suggested_tags: SuggestedTags
    title: str
    def __init__(self, id: _Optional[str] = ..., title: _Optional[str] = ..., description: _Optional[str] = ..., publishedAt: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., channelId: _Optional[str] = ..., channelTitle: _Optional[str] = ..., rootThumbnail: _Optional[str] = ..., record_date: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., suggested_tags: _Optional[_Union[SuggestedTags, _Mapping]] = ...) -> None: ...

class GetFancamRequest(_message.Message):
    __slots__ = ["id"]
    ID_FIELD_NUMBER: _ClassVar[int]
    id: str
    def __init__(self, id: _Optional[str] = ...) -> None: ...

class GetFancamsLatestRequest(_message.Message):
    __slots__ = ["max_results"]
    MAX_RESULTS_FIELD_NUMBER: _ClassVar[int]
    max_results: int
    def __init__(self, max_results: _Optional[int] = ...) -> None: ...

class GetFancamsRequest(_message.Message):
    __slots__ = ["ids"]
    IDS_FIELD_NUMBER: _ClassVar[int]
    ids: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, ids: _Optional[_Iterable[str]] = ...) -> None: ...

class SuggestedTags(_message.Message):
    __slots__ = ["en_artist", "en_group", "en_song", "kr_artist", "kr_group", "kr_song"]
    EN_ARTIST_FIELD_NUMBER: _ClassVar[int]
    EN_GROUP_FIELD_NUMBER: _ClassVar[int]
    EN_SONG_FIELD_NUMBER: _ClassVar[int]
    KR_ARTIST_FIELD_NUMBER: _ClassVar[int]
    KR_GROUP_FIELD_NUMBER: _ClassVar[int]
    KR_SONG_FIELD_NUMBER: _ClassVar[int]
    en_artist: _containers.RepeatedScalarFieldContainer[str]
    en_group: _containers.RepeatedScalarFieldContainer[str]
    en_song: _containers.RepeatedScalarFieldContainer[str]
    kr_artist: _containers.RepeatedScalarFieldContainer[str]
    kr_group: _containers.RepeatedScalarFieldContainer[str]
    kr_song: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, en_artist: _Optional[_Iterable[str]] = ..., en_group: _Optional[_Iterable[str]] = ..., en_song: _Optional[_Iterable[str]] = ..., kr_artist: _Optional[_Iterable[str]] = ..., kr_group: _Optional[_Iterable[str]] = ..., kr_song: _Optional[_Iterable[str]] = ...) -> None: ...

class Gender(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = []
