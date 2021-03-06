# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: topos/scores/v1/scores.proto

from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='topos/scores/v1/scores.proto',
  package='topos.scores.v1',
  syntax='proto3',
  serialized_options=b'ZAgithub.com/topos-ai/topos-apis/genproto/go/topos/scores/v1;scores',
  serialized_pb=b'\n\x1ctopos/scores/v1/scores.proto\x12\x0ftopos.scores.v1\x1a\x1cgoogle/api/annotations.proto\":\n\x05Score\x12\x10\n\x08vertex_a\x18\x01 \x01(\t\x12\x10\n\x08vertex_b\x18\x02 \x01(\t\x12\r\n\x05score\x18\x03 \x01(\x01\"H\n\x14GetGraphScoreRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x10\n\x08vertex_a\x18\x02 \x01(\t\x12\x10\n\x08vertex_b\x18\x03 \x01(\t\"\x8d\x01\n\x15TopGraphScoresRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x11\n\tpage_size\x18\x02 \x01(\x05\x12\x10\n\x08vertex_a\x18\x03 \x01(\t\x12\x10\n\x08vertex_b\x18\x04 \x01(\t\x12\x17\n\x0f\x61scending_order\x18\x05 \x01(\x08\x12\x16\n\x0evertex_filters\x18\x06 \x03(\t\"@\n\x16TopGraphScoresResponse\x12&\n\x06scores\x18\x01 \x03(\x0b\x32\x16.topos.scores.v1.Score2\x88\x02\n\x06Scores\x12r\n\rGetGraphScore\x12%.topos.scores.v1.GetGraphScoreRequest\x1a\x16.topos.scores.v1.Score\"\"\x82\xd3\xe4\x93\x02\x1c\x12\x1a/v1/{name=graphs/*}/scores\x12\x89\x01\n\x0eTopGraphScores\x12&.topos.scores.v1.TopGraphScoresRequest\x1a\'.topos.scores.v1.TopGraphScoresResponse\"&\x82\xd3\xe4\x93\x02 \x12\x1e/v1/{name=graphs/*}/scores:topBCZAgithub.com/topos-ai/topos-apis/genproto/go/topos/scores/v1;scoresb\x06proto3'
  ,
  dependencies=[google_dot_api_dot_annotations__pb2.DESCRIPTOR,])




_SCORE = _descriptor.Descriptor(
  name='Score',
  full_name='topos.scores.v1.Score',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='vertex_a', full_name='topos.scores.v1.Score.vertex_a', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='vertex_b', full_name='topos.scores.v1.Score.vertex_b', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='score', full_name='topos.scores.v1.Score.score', index=2,
      number=3, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=79,
  serialized_end=137,
)


_GETGRAPHSCOREREQUEST = _descriptor.Descriptor(
  name='GetGraphScoreRequest',
  full_name='topos.scores.v1.GetGraphScoreRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='topos.scores.v1.GetGraphScoreRequest.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='vertex_a', full_name='topos.scores.v1.GetGraphScoreRequest.vertex_a', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='vertex_b', full_name='topos.scores.v1.GetGraphScoreRequest.vertex_b', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=139,
  serialized_end=211,
)


_TOPGRAPHSCORESREQUEST = _descriptor.Descriptor(
  name='TopGraphScoresRequest',
  full_name='topos.scores.v1.TopGraphScoresRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='topos.scores.v1.TopGraphScoresRequest.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='page_size', full_name='topos.scores.v1.TopGraphScoresRequest.page_size', index=1,
      number=2, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='vertex_a', full_name='topos.scores.v1.TopGraphScoresRequest.vertex_a', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='vertex_b', full_name='topos.scores.v1.TopGraphScoresRequest.vertex_b', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='ascending_order', full_name='topos.scores.v1.TopGraphScoresRequest.ascending_order', index=4,
      number=5, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='vertex_filters', full_name='topos.scores.v1.TopGraphScoresRequest.vertex_filters', index=5,
      number=6, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=214,
  serialized_end=355,
)


_TOPGRAPHSCORESRESPONSE = _descriptor.Descriptor(
  name='TopGraphScoresResponse',
  full_name='topos.scores.v1.TopGraphScoresResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='scores', full_name='topos.scores.v1.TopGraphScoresResponse.scores', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=357,
  serialized_end=421,
)

_TOPGRAPHSCORESRESPONSE.fields_by_name['scores'].message_type = _SCORE
DESCRIPTOR.message_types_by_name['Score'] = _SCORE
DESCRIPTOR.message_types_by_name['GetGraphScoreRequest'] = _GETGRAPHSCOREREQUEST
DESCRIPTOR.message_types_by_name['TopGraphScoresRequest'] = _TOPGRAPHSCORESREQUEST
DESCRIPTOR.message_types_by_name['TopGraphScoresResponse'] = _TOPGRAPHSCORESRESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Score = _reflection.GeneratedProtocolMessageType('Score', (_message.Message,), {
  'DESCRIPTOR' : _SCORE,
  '__module__' : 'topos.scores.v1.scores_pb2'
  # @@protoc_insertion_point(class_scope:topos.scores.v1.Score)
  })
_sym_db.RegisterMessage(Score)

GetGraphScoreRequest = _reflection.GeneratedProtocolMessageType('GetGraphScoreRequest', (_message.Message,), {
  'DESCRIPTOR' : _GETGRAPHSCOREREQUEST,
  '__module__' : 'topos.scores.v1.scores_pb2'
  # @@protoc_insertion_point(class_scope:topos.scores.v1.GetGraphScoreRequest)
  })
_sym_db.RegisterMessage(GetGraphScoreRequest)

TopGraphScoresRequest = _reflection.GeneratedProtocolMessageType('TopGraphScoresRequest', (_message.Message,), {
  'DESCRIPTOR' : _TOPGRAPHSCORESREQUEST,
  '__module__' : 'topos.scores.v1.scores_pb2'
  # @@protoc_insertion_point(class_scope:topos.scores.v1.TopGraphScoresRequest)
  })
_sym_db.RegisterMessage(TopGraphScoresRequest)

TopGraphScoresResponse = _reflection.GeneratedProtocolMessageType('TopGraphScoresResponse', (_message.Message,), {
  'DESCRIPTOR' : _TOPGRAPHSCORESRESPONSE,
  '__module__' : 'topos.scores.v1.scores_pb2'
  # @@protoc_insertion_point(class_scope:topos.scores.v1.TopGraphScoresResponse)
  })
_sym_db.RegisterMessage(TopGraphScoresResponse)


DESCRIPTOR._options = None

_SCORES = _descriptor.ServiceDescriptor(
  name='Scores',
  full_name='topos.scores.v1.Scores',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=424,
  serialized_end=688,
  methods=[
  _descriptor.MethodDescriptor(
    name='GetGraphScore',
    full_name='topos.scores.v1.Scores.GetGraphScore',
    index=0,
    containing_service=None,
    input_type=_GETGRAPHSCOREREQUEST,
    output_type=_SCORE,
    serialized_options=b'\202\323\344\223\002\034\022\032/v1/{name=graphs/*}/scores',
  ),
  _descriptor.MethodDescriptor(
    name='TopGraphScores',
    full_name='topos.scores.v1.Scores.TopGraphScores',
    index=1,
    containing_service=None,
    input_type=_TOPGRAPHSCORESREQUEST,
    output_type=_TOPGRAPHSCORESRESPONSE,
    serialized_options=b'\202\323\344\223\002 \022\036/v1/{name=graphs/*}/scores:top',
  ),
])
_sym_db.RegisterServiceDescriptor(_SCORES)

DESCRIPTOR.services_by_name['Scores'] = _SCORES

# @@protoc_insertion_point(module_scope)
