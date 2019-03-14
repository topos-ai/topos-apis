# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: topos/points/v1/points.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='topos/points/v1/points.proto',
  package='topos.points.v1',
  syntax='proto3',
  serialized_options=_b('Z\026topos/points/v1;points'),
  serialized_pb=_b('\n\x1ctopos/points/v1/points.proto\x12\x0ftopos.points.v1\x1a\x1cgoogle/api/annotations.proto\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\x8d\x01\n\x05\x42rand\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x14\n\x0c\x64isplay_name\x18\x02 \x01(\t\x12/\n\x0b\x63reate_time\x18\x03 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12/\n\x0bupdate_time\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\"-\n\x06LatLng\x12\x10\n\x08latitude\x18\x01 \x01(\x01\x12\x11\n\tlongitude\x18\x02 \x01(\x01\"\xe2\x01\n\x05Point\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\r\n\x05\x62rand\x18\x02 \x01(\t\x12\x14\n\x0c\x64isplay_name\x18\x03 \x01(\t\x12\x19\n\x11\x66ormatted_address\x18\x04 \x01(\t\x12)\n\x08location\x18\x05 \x01(\x0b\x32\x17.topos.points.v1.LatLng\x12/\n\x0b\x63reate_time\x18\x06 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12/\n\x0bupdate_time\x18\x07 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\"\xf6\x01\n\x0bPointSource\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\r\n\x05point\x18\x02 \x01(\t\x12\x14\n\x0c\x64isplay_name\x18\x03 \x01(\t\x12\x19\n\x11\x66ormatted_address\x18\x04 \x01(\t\x12)\n\x08location\x18\x05 \x01(\x0b\x32\x17.topos.points.v1.LatLng\x12\x0c\n\x04tags\x18\x06 \x03(\t\x12/\n\x0b\x63reate_time\x18\x07 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12/\n\x0bupdate_time\x18\x08 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\"\x1f\n\x0fGetBrandRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\";\n\x12\x43reateBrandRequest\x12%\n\x05\x62rand\x18\x01 \x01(\x0b\x32\x16.topos.points.v1.Brand\";\n\x12UpdateBrandRequest\x12%\n\x05\x62rand\x18\x01 \x01(\x0b\x32\x16.topos.points.v1.Brand\"\"\n\x12\x44\x65leteBrandRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\"\x1f\n\x0fGetPointRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\";\n\x12\x43reatePointRequest\x12%\n\x05point\x18\x01 \x01(\x0b\x32\x16.topos.points.v1.Point\";\n\x12UpdatePointRequest\x12%\n\x05point\x18\x01 \x01(\x0b\x32\x16.topos.points.v1.Point\"\"\n\x12\x44\x65letePointRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\"%\n\x15GetPointSourceRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\"^\n\x18\x43reatePointSourceRequest\x12\x0e\n\x06source\x18\x01 \x01(\t\x12\x32\n\x0cpoint_source\x18\x02 \x01(\x0b\x32\x1c.topos.points.v1.PointSource\"N\n\x18UpdatePointSourceRequest\x12\x32\n\x0cpoint_source\x18\x01 \x01(\x0b\x32\x1c.topos.points.v1.PointSource\"(\n\x18\x44\x65letePointSourceRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\"<\n\x16LinkPointSourceRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x14\n\x0cpoint_source\x18\x02 \x01(\t2\xcf\x0b\n\x06Points\x12\x61\n\x08GetBrand\x12 .topos.points.v1.GetBrandRequest\x1a\x16.topos.points.v1.Brand\"\x1b\x82\xd3\xe4\x93\x02\x15\x12\x13/v1/{name=brands/*}\x12\x65\n\x0b\x43reateBrand\x12#.topos.points.v1.CreateBrandRequest\x1a\x16.topos.points.v1.Brand\"\x19\x82\xd3\xe4\x93\x02\x13\"\n/v1/brands:\x05\x62rand\x12t\n\x0bUpdateBrand\x12#.topos.points.v1.UpdateBrandRequest\x1a\x16.topos.points.v1.Brand\"(\x82\xd3\xe4\x93\x02\"\x1a\x19/v1/{brand.name=brands/*}:\x05\x62rand\x12g\n\x0b\x44\x65leteBrand\x12#.topos.points.v1.DeleteBrandRequest\x1a\x16.google.protobuf.Empty\"\x1b\x82\xd3\xe4\x93\x02\x15*\x13/v1/{name=brands/*}\x12\x61\n\x08GetPoint\x12 .topos.points.v1.GetPointRequest\x1a\x16.topos.points.v1.Point\"\x1b\x82\xd3\xe4\x93\x02\x15\x12\x13/v1/{name=points/*}\x12t\n\x0bUpdatePoint\x12#.topos.points.v1.UpdatePointRequest\x1a\x16.topos.points.v1.Point\"(\x82\xd3\xe4\x93\x02\"\x1a\x19/v1/{point.name=points/*}:\x05point\x12g\n\x0b\x44\x65letePoint\x12#.topos.points.v1.DeletePointRequest\x1a\x16.google.protobuf.Empty\"\x1b\x82\xd3\xe4\x93\x02\x15*\x13/v1/{name=points/*}\x12\x84\x01\n\x0eGetPointSource\x12&.topos.points.v1.GetPointSourceRequest\x1a\x1c.topos.points.v1.PointSource\",\x82\xd3\xe4\x93\x02&\x12$/v1/{name=sources/*/point_sources/*}\x12\x98\x01\n\x11\x43reatePointSource\x12).topos.points.v1.CreatePointSourceRequest\x1a\x1c.topos.points.v1.PointSource\":\x82\xd3\xe4\x93\x02\x34\"$/v1/{source=sources/*}/point_sources:\x0cpoint_source\x12\xa5\x01\n\x11UpdatePointSource\x12).topos.points.v1.UpdatePointSourceRequest\x1a\x1c.topos.points.v1.PointSource\"G\x82\xd3\xe4\x93\x02\x41\x1a\x31/v1/{point_source.name=sources/*/point_sources/*}:\x0cpoint_source\x12\x84\x01\n\x11\x44\x65letePointSource\x12).topos.points.v1.DeletePointSourceRequest\x1a\x16.google.protobuf.Empty\",\x82\xd3\xe4\x93\x02&*$/v1/{name=sources/*/point_sources/*}\x12\x88\x01\n\x0fLinkPointSource\x12\'.topos.points.v1.LinkPointSourceRequest\x1a\x16.topos.points.v1.Point\"4\x82\xd3\xe4\x93\x02.\")/v1/{name=sources/*/point_sources/*}:link:\x01*B\x18Z\x16topos/points/v1;pointsb\x06proto3')
  ,
  dependencies=[google_dot_api_dot_annotations__pb2.DESCRIPTOR,google_dot_protobuf_dot_empty__pb2.DESCRIPTOR,google_dot_protobuf_dot_timestamp__pb2.DESCRIPTOR,])




_BRAND = _descriptor.Descriptor(
  name='Brand',
  full_name='topos.points.v1.Brand',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='topos.points.v1.Brand.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='display_name', full_name='topos.points.v1.Brand.display_name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='create_time', full_name='topos.points.v1.Brand.create_time', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='update_time', full_name='topos.points.v1.Brand.update_time', index=3,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
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
  serialized_start=142,
  serialized_end=283,
)


_LATLNG = _descriptor.Descriptor(
  name='LatLng',
  full_name='topos.points.v1.LatLng',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='latitude', full_name='topos.points.v1.LatLng.latitude', index=0,
      number=1, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='longitude', full_name='topos.points.v1.LatLng.longitude', index=1,
      number=2, type=1, cpp_type=5, label=1,
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
  serialized_start=285,
  serialized_end=330,
)


_POINT = _descriptor.Descriptor(
  name='Point',
  full_name='topos.points.v1.Point',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='topos.points.v1.Point.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='brand', full_name='topos.points.v1.Point.brand', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='display_name', full_name='topos.points.v1.Point.display_name', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='formatted_address', full_name='topos.points.v1.Point.formatted_address', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='location', full_name='topos.points.v1.Point.location', index=4,
      number=5, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='create_time', full_name='topos.points.v1.Point.create_time', index=5,
      number=6, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='update_time', full_name='topos.points.v1.Point.update_time', index=6,
      number=7, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
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
  serialized_start=333,
  serialized_end=559,
)


_POINTSOURCE = _descriptor.Descriptor(
  name='PointSource',
  full_name='topos.points.v1.PointSource',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='topos.points.v1.PointSource.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='point', full_name='topos.points.v1.PointSource.point', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='display_name', full_name='topos.points.v1.PointSource.display_name', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='formatted_address', full_name='topos.points.v1.PointSource.formatted_address', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='location', full_name='topos.points.v1.PointSource.location', index=4,
      number=5, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='tags', full_name='topos.points.v1.PointSource.tags', index=5,
      number=6, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='create_time', full_name='topos.points.v1.PointSource.create_time', index=6,
      number=7, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='update_time', full_name='topos.points.v1.PointSource.update_time', index=7,
      number=8, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
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
  serialized_start=562,
  serialized_end=808,
)


_GETBRANDREQUEST = _descriptor.Descriptor(
  name='GetBrandRequest',
  full_name='topos.points.v1.GetBrandRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='topos.points.v1.GetBrandRequest.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
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
  serialized_start=810,
  serialized_end=841,
)


_CREATEBRANDREQUEST = _descriptor.Descriptor(
  name='CreateBrandRequest',
  full_name='topos.points.v1.CreateBrandRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='brand', full_name='topos.points.v1.CreateBrandRequest.brand', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
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
  serialized_start=843,
  serialized_end=902,
)


_UPDATEBRANDREQUEST = _descriptor.Descriptor(
  name='UpdateBrandRequest',
  full_name='topos.points.v1.UpdateBrandRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='brand', full_name='topos.points.v1.UpdateBrandRequest.brand', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
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
  serialized_start=904,
  serialized_end=963,
)


_DELETEBRANDREQUEST = _descriptor.Descriptor(
  name='DeleteBrandRequest',
  full_name='topos.points.v1.DeleteBrandRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='topos.points.v1.DeleteBrandRequest.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
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
  serialized_start=965,
  serialized_end=999,
)


_GETPOINTREQUEST = _descriptor.Descriptor(
  name='GetPointRequest',
  full_name='topos.points.v1.GetPointRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='topos.points.v1.GetPointRequest.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
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
  serialized_start=1001,
  serialized_end=1032,
)


_CREATEPOINTREQUEST = _descriptor.Descriptor(
  name='CreatePointRequest',
  full_name='topos.points.v1.CreatePointRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='point', full_name='topos.points.v1.CreatePointRequest.point', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
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
  serialized_start=1034,
  serialized_end=1093,
)


_UPDATEPOINTREQUEST = _descriptor.Descriptor(
  name='UpdatePointRequest',
  full_name='topos.points.v1.UpdatePointRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='point', full_name='topos.points.v1.UpdatePointRequest.point', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
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
  serialized_start=1095,
  serialized_end=1154,
)


_DELETEPOINTREQUEST = _descriptor.Descriptor(
  name='DeletePointRequest',
  full_name='topos.points.v1.DeletePointRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='topos.points.v1.DeletePointRequest.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
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
  serialized_start=1156,
  serialized_end=1190,
)


_GETPOINTSOURCEREQUEST = _descriptor.Descriptor(
  name='GetPointSourceRequest',
  full_name='topos.points.v1.GetPointSourceRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='topos.points.v1.GetPointSourceRequest.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
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
  serialized_start=1192,
  serialized_end=1229,
)


_CREATEPOINTSOURCEREQUEST = _descriptor.Descriptor(
  name='CreatePointSourceRequest',
  full_name='topos.points.v1.CreatePointSourceRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='source', full_name='topos.points.v1.CreatePointSourceRequest.source', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='point_source', full_name='topos.points.v1.CreatePointSourceRequest.point_source', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
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
  serialized_start=1231,
  serialized_end=1325,
)


_UPDATEPOINTSOURCEREQUEST = _descriptor.Descriptor(
  name='UpdatePointSourceRequest',
  full_name='topos.points.v1.UpdatePointSourceRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='point_source', full_name='topos.points.v1.UpdatePointSourceRequest.point_source', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
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
  serialized_start=1327,
  serialized_end=1405,
)


_DELETEPOINTSOURCEREQUEST = _descriptor.Descriptor(
  name='DeletePointSourceRequest',
  full_name='topos.points.v1.DeletePointSourceRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='topos.points.v1.DeletePointSourceRequest.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
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
  serialized_start=1407,
  serialized_end=1447,
)


_LINKPOINTSOURCEREQUEST = _descriptor.Descriptor(
  name='LinkPointSourceRequest',
  full_name='topos.points.v1.LinkPointSourceRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='topos.points.v1.LinkPointSourceRequest.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='point_source', full_name='topos.points.v1.LinkPointSourceRequest.point_source', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
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
  serialized_start=1449,
  serialized_end=1509,
)

_BRAND.fields_by_name['create_time'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_BRAND.fields_by_name['update_time'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_POINT.fields_by_name['location'].message_type = _LATLNG
_POINT.fields_by_name['create_time'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_POINT.fields_by_name['update_time'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_POINTSOURCE.fields_by_name['location'].message_type = _LATLNG
_POINTSOURCE.fields_by_name['create_time'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_POINTSOURCE.fields_by_name['update_time'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_CREATEBRANDREQUEST.fields_by_name['brand'].message_type = _BRAND
_UPDATEBRANDREQUEST.fields_by_name['brand'].message_type = _BRAND
_CREATEPOINTREQUEST.fields_by_name['point'].message_type = _POINT
_UPDATEPOINTREQUEST.fields_by_name['point'].message_type = _POINT
_CREATEPOINTSOURCEREQUEST.fields_by_name['point_source'].message_type = _POINTSOURCE
_UPDATEPOINTSOURCEREQUEST.fields_by_name['point_source'].message_type = _POINTSOURCE
DESCRIPTOR.message_types_by_name['Brand'] = _BRAND
DESCRIPTOR.message_types_by_name['LatLng'] = _LATLNG
DESCRIPTOR.message_types_by_name['Point'] = _POINT
DESCRIPTOR.message_types_by_name['PointSource'] = _POINTSOURCE
DESCRIPTOR.message_types_by_name['GetBrandRequest'] = _GETBRANDREQUEST
DESCRIPTOR.message_types_by_name['CreateBrandRequest'] = _CREATEBRANDREQUEST
DESCRIPTOR.message_types_by_name['UpdateBrandRequest'] = _UPDATEBRANDREQUEST
DESCRIPTOR.message_types_by_name['DeleteBrandRequest'] = _DELETEBRANDREQUEST
DESCRIPTOR.message_types_by_name['GetPointRequest'] = _GETPOINTREQUEST
DESCRIPTOR.message_types_by_name['CreatePointRequest'] = _CREATEPOINTREQUEST
DESCRIPTOR.message_types_by_name['UpdatePointRequest'] = _UPDATEPOINTREQUEST
DESCRIPTOR.message_types_by_name['DeletePointRequest'] = _DELETEPOINTREQUEST
DESCRIPTOR.message_types_by_name['GetPointSourceRequest'] = _GETPOINTSOURCEREQUEST
DESCRIPTOR.message_types_by_name['CreatePointSourceRequest'] = _CREATEPOINTSOURCEREQUEST
DESCRIPTOR.message_types_by_name['UpdatePointSourceRequest'] = _UPDATEPOINTSOURCEREQUEST
DESCRIPTOR.message_types_by_name['DeletePointSourceRequest'] = _DELETEPOINTSOURCEREQUEST
DESCRIPTOR.message_types_by_name['LinkPointSourceRequest'] = _LINKPOINTSOURCEREQUEST
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Brand = _reflection.GeneratedProtocolMessageType('Brand', (_message.Message,), dict(
  DESCRIPTOR = _BRAND,
  __module__ = 'topos.points.v1.points_pb2'
  # @@protoc_insertion_point(class_scope:topos.points.v1.Brand)
  ))
_sym_db.RegisterMessage(Brand)

LatLng = _reflection.GeneratedProtocolMessageType('LatLng', (_message.Message,), dict(
  DESCRIPTOR = _LATLNG,
  __module__ = 'topos.points.v1.points_pb2'
  # @@protoc_insertion_point(class_scope:topos.points.v1.LatLng)
  ))
_sym_db.RegisterMessage(LatLng)

Point = _reflection.GeneratedProtocolMessageType('Point', (_message.Message,), dict(
  DESCRIPTOR = _POINT,
  __module__ = 'topos.points.v1.points_pb2'
  # @@protoc_insertion_point(class_scope:topos.points.v1.Point)
  ))
_sym_db.RegisterMessage(Point)

PointSource = _reflection.GeneratedProtocolMessageType('PointSource', (_message.Message,), dict(
  DESCRIPTOR = _POINTSOURCE,
  __module__ = 'topos.points.v1.points_pb2'
  # @@protoc_insertion_point(class_scope:topos.points.v1.PointSource)
  ))
_sym_db.RegisterMessage(PointSource)

GetBrandRequest = _reflection.GeneratedProtocolMessageType('GetBrandRequest', (_message.Message,), dict(
  DESCRIPTOR = _GETBRANDREQUEST,
  __module__ = 'topos.points.v1.points_pb2'
  # @@protoc_insertion_point(class_scope:topos.points.v1.GetBrandRequest)
  ))
_sym_db.RegisterMessage(GetBrandRequest)

CreateBrandRequest = _reflection.GeneratedProtocolMessageType('CreateBrandRequest', (_message.Message,), dict(
  DESCRIPTOR = _CREATEBRANDREQUEST,
  __module__ = 'topos.points.v1.points_pb2'
  # @@protoc_insertion_point(class_scope:topos.points.v1.CreateBrandRequest)
  ))
_sym_db.RegisterMessage(CreateBrandRequest)

UpdateBrandRequest = _reflection.GeneratedProtocolMessageType('UpdateBrandRequest', (_message.Message,), dict(
  DESCRIPTOR = _UPDATEBRANDREQUEST,
  __module__ = 'topos.points.v1.points_pb2'
  # @@protoc_insertion_point(class_scope:topos.points.v1.UpdateBrandRequest)
  ))
_sym_db.RegisterMessage(UpdateBrandRequest)

DeleteBrandRequest = _reflection.GeneratedProtocolMessageType('DeleteBrandRequest', (_message.Message,), dict(
  DESCRIPTOR = _DELETEBRANDREQUEST,
  __module__ = 'topos.points.v1.points_pb2'
  # @@protoc_insertion_point(class_scope:topos.points.v1.DeleteBrandRequest)
  ))
_sym_db.RegisterMessage(DeleteBrandRequest)

GetPointRequest = _reflection.GeneratedProtocolMessageType('GetPointRequest', (_message.Message,), dict(
  DESCRIPTOR = _GETPOINTREQUEST,
  __module__ = 'topos.points.v1.points_pb2'
  # @@protoc_insertion_point(class_scope:topos.points.v1.GetPointRequest)
  ))
_sym_db.RegisterMessage(GetPointRequest)

CreatePointRequest = _reflection.GeneratedProtocolMessageType('CreatePointRequest', (_message.Message,), dict(
  DESCRIPTOR = _CREATEPOINTREQUEST,
  __module__ = 'topos.points.v1.points_pb2'
  # @@protoc_insertion_point(class_scope:topos.points.v1.CreatePointRequest)
  ))
_sym_db.RegisterMessage(CreatePointRequest)

UpdatePointRequest = _reflection.GeneratedProtocolMessageType('UpdatePointRequest', (_message.Message,), dict(
  DESCRIPTOR = _UPDATEPOINTREQUEST,
  __module__ = 'topos.points.v1.points_pb2'
  # @@protoc_insertion_point(class_scope:topos.points.v1.UpdatePointRequest)
  ))
_sym_db.RegisterMessage(UpdatePointRequest)

DeletePointRequest = _reflection.GeneratedProtocolMessageType('DeletePointRequest', (_message.Message,), dict(
  DESCRIPTOR = _DELETEPOINTREQUEST,
  __module__ = 'topos.points.v1.points_pb2'
  # @@protoc_insertion_point(class_scope:topos.points.v1.DeletePointRequest)
  ))
_sym_db.RegisterMessage(DeletePointRequest)

GetPointSourceRequest = _reflection.GeneratedProtocolMessageType('GetPointSourceRequest', (_message.Message,), dict(
  DESCRIPTOR = _GETPOINTSOURCEREQUEST,
  __module__ = 'topos.points.v1.points_pb2'
  # @@protoc_insertion_point(class_scope:topos.points.v1.GetPointSourceRequest)
  ))
_sym_db.RegisterMessage(GetPointSourceRequest)

CreatePointSourceRequest = _reflection.GeneratedProtocolMessageType('CreatePointSourceRequest', (_message.Message,), dict(
  DESCRIPTOR = _CREATEPOINTSOURCEREQUEST,
  __module__ = 'topos.points.v1.points_pb2'
  # @@protoc_insertion_point(class_scope:topos.points.v1.CreatePointSourceRequest)
  ))
_sym_db.RegisterMessage(CreatePointSourceRequest)

UpdatePointSourceRequest = _reflection.GeneratedProtocolMessageType('UpdatePointSourceRequest', (_message.Message,), dict(
  DESCRIPTOR = _UPDATEPOINTSOURCEREQUEST,
  __module__ = 'topos.points.v1.points_pb2'
  # @@protoc_insertion_point(class_scope:topos.points.v1.UpdatePointSourceRequest)
  ))
_sym_db.RegisterMessage(UpdatePointSourceRequest)

DeletePointSourceRequest = _reflection.GeneratedProtocolMessageType('DeletePointSourceRequest', (_message.Message,), dict(
  DESCRIPTOR = _DELETEPOINTSOURCEREQUEST,
  __module__ = 'topos.points.v1.points_pb2'
  # @@protoc_insertion_point(class_scope:topos.points.v1.DeletePointSourceRequest)
  ))
_sym_db.RegisterMessage(DeletePointSourceRequest)

LinkPointSourceRequest = _reflection.GeneratedProtocolMessageType('LinkPointSourceRequest', (_message.Message,), dict(
  DESCRIPTOR = _LINKPOINTSOURCEREQUEST,
  __module__ = 'topos.points.v1.points_pb2'
  # @@protoc_insertion_point(class_scope:topos.points.v1.LinkPointSourceRequest)
  ))
_sym_db.RegisterMessage(LinkPointSourceRequest)


DESCRIPTOR._options = None

_POINTS = _descriptor.ServiceDescriptor(
  name='Points',
  full_name='topos.points.v1.Points',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=1512,
  serialized_end=2999,
  methods=[
  _descriptor.MethodDescriptor(
    name='GetBrand',
    full_name='topos.points.v1.Points.GetBrand',
    index=0,
    containing_service=None,
    input_type=_GETBRANDREQUEST,
    output_type=_BRAND,
    serialized_options=_b('\202\323\344\223\002\025\022\023/v1/{name=brands/*}'),
  ),
  _descriptor.MethodDescriptor(
    name='CreateBrand',
    full_name='topos.points.v1.Points.CreateBrand',
    index=1,
    containing_service=None,
    input_type=_CREATEBRANDREQUEST,
    output_type=_BRAND,
    serialized_options=_b('\202\323\344\223\002\023\"\n/v1/brands:\005brand'),
  ),
  _descriptor.MethodDescriptor(
    name='UpdateBrand',
    full_name='topos.points.v1.Points.UpdateBrand',
    index=2,
    containing_service=None,
    input_type=_UPDATEBRANDREQUEST,
    output_type=_BRAND,
    serialized_options=_b('\202\323\344\223\002\"\032\031/v1/{brand.name=brands/*}:\005brand'),
  ),
  _descriptor.MethodDescriptor(
    name='DeleteBrand',
    full_name='topos.points.v1.Points.DeleteBrand',
    index=3,
    containing_service=None,
    input_type=_DELETEBRANDREQUEST,
    output_type=google_dot_protobuf_dot_empty__pb2._EMPTY,
    serialized_options=_b('\202\323\344\223\002\025*\023/v1/{name=brands/*}'),
  ),
  _descriptor.MethodDescriptor(
    name='GetPoint',
    full_name='topos.points.v1.Points.GetPoint',
    index=4,
    containing_service=None,
    input_type=_GETPOINTREQUEST,
    output_type=_POINT,
    serialized_options=_b('\202\323\344\223\002\025\022\023/v1/{name=points/*}'),
  ),
  _descriptor.MethodDescriptor(
    name='UpdatePoint',
    full_name='topos.points.v1.Points.UpdatePoint',
    index=5,
    containing_service=None,
    input_type=_UPDATEPOINTREQUEST,
    output_type=_POINT,
    serialized_options=_b('\202\323\344\223\002\"\032\031/v1/{point.name=points/*}:\005point'),
  ),
  _descriptor.MethodDescriptor(
    name='DeletePoint',
    full_name='topos.points.v1.Points.DeletePoint',
    index=6,
    containing_service=None,
    input_type=_DELETEPOINTREQUEST,
    output_type=google_dot_protobuf_dot_empty__pb2._EMPTY,
    serialized_options=_b('\202\323\344\223\002\025*\023/v1/{name=points/*}'),
  ),
  _descriptor.MethodDescriptor(
    name='GetPointSource',
    full_name='topos.points.v1.Points.GetPointSource',
    index=7,
    containing_service=None,
    input_type=_GETPOINTSOURCEREQUEST,
    output_type=_POINTSOURCE,
    serialized_options=_b('\202\323\344\223\002&\022$/v1/{name=sources/*/point_sources/*}'),
  ),
  _descriptor.MethodDescriptor(
    name='CreatePointSource',
    full_name='topos.points.v1.Points.CreatePointSource',
    index=8,
    containing_service=None,
    input_type=_CREATEPOINTSOURCEREQUEST,
    output_type=_POINTSOURCE,
    serialized_options=_b('\202\323\344\223\0024\"$/v1/{source=sources/*}/point_sources:\014point_source'),
  ),
  _descriptor.MethodDescriptor(
    name='UpdatePointSource',
    full_name='topos.points.v1.Points.UpdatePointSource',
    index=9,
    containing_service=None,
    input_type=_UPDATEPOINTSOURCEREQUEST,
    output_type=_POINTSOURCE,
    serialized_options=_b('\202\323\344\223\002A\0321/v1/{point_source.name=sources/*/point_sources/*}:\014point_source'),
  ),
  _descriptor.MethodDescriptor(
    name='DeletePointSource',
    full_name='topos.points.v1.Points.DeletePointSource',
    index=10,
    containing_service=None,
    input_type=_DELETEPOINTSOURCEREQUEST,
    output_type=google_dot_protobuf_dot_empty__pb2._EMPTY,
    serialized_options=_b('\202\323\344\223\002&*$/v1/{name=sources/*/point_sources/*}'),
  ),
  _descriptor.MethodDescriptor(
    name='LinkPointSource',
    full_name='topos.points.v1.Points.LinkPointSource',
    index=11,
    containing_service=None,
    input_type=_LINKPOINTSOURCEREQUEST,
    output_type=_POINT,
    serialized_options=_b('\202\323\344\223\002.\")/v1/{name=sources/*/point_sources/*}:link:\001*'),
  ),
])
_sym_db.RegisterServiceDescriptor(_POINTS)

DESCRIPTOR.services_by_name['Points'] = _POINTS

# @@protoc_insertion_point(module_scope)
