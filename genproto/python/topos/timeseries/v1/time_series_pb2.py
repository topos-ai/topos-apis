# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: topos/timeseries/v1/time-series.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='topos/timeseries/v1/time-series.proto',
  package='topos.timeseries.v1',
  syntax='proto3',
  serialized_options=_b('ZIgithub.com/topos-ai/topos-apis/genproto/go/topos/timeseries/v1;timeseries'),
  serialized_pb=_b('\n%topos/timeseries/v1/time-series.proto\x12\x13topos.timeseries.v1\x1a\x1cgoogle/api/annotations.proto\x1a\x1bgoogle/protobuf/empty.proto\"w\n\x11TimeSeriesFeature\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x14\n\x0c\x64isplay_name\x18\x02 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x03 \x01(\t\x12\x0e\n\x06source\x18\x04 \x01(\t\x12\x19\n\x11sampling_interval\x18\x05 \x01(\x03\"\x82\x01\n\rTimeSeriesBin\x12\x11\n\tregion_id\x18\x01 \x01(\t\x12\x14\n\x0c\x66\x65\x61ture_name\x18\x02 \x01(\t\x12%\n\x03\x64\x61y\x18\x03 \x01(\x0e\x32\x18.topos.timeseries.v1.Day\x12\x12\n\nstart_time\x18\x04 \x01(\t\x12\r\n\x05value\x18\x06 \x01(\x03\"\"\n\x10TimeSeriesValues\x12\x0e\n\x06values\x18\x01 \x03(\x03\"\x91\x01\n\x17GetTimeSeriesBinRequest\x12\x13\n\x0bregion_name\x18\x01 \x01(\t\x12\x14\n\x0c\x66\x65\x61ture_name\x18\x02 \x01(\t\x12%\n\x03\x64\x61y\x18\x03 \x01(\x0e\x32\x18.topos.timeseries.v1.Day\x12\x12\n\nstart_time\x18\x04 \x01(\t\x12\x10\n\x08\x65nd_time\x18\x05 \x01(\t\"\x96\x01\n\x1bGetTimeSeriesBinBulkRequest\x12\x14\n\x0c\x66\x65\x61ture_name\x18\x01 \x01(\t\x12\x14\n\x0cregion_names\x18\x02 \x03(\t\x12%\n\x03\x64\x61y\x18\x03 \x01(\x0e\x32\x18.topos.timeseries.v1.Day\x12\x12\n\nstart_time\x18\x04 \x01(\t\x12\x10\n\x08\x65nd_time\x18\x05 \x01(\t\"\x7f\n GetTimeSeriesAveragesBulkRequest\x12\x14\n\x0c\x66\x65\x61ture_name\x18\x01 \x01(\t\x12\x14\n\x0cregion_names\x18\x02 \x03(\t\x12/\n\x06period\x18\x03 \x01(\x0e\x32\x1f.topos.timeseries.v1.TimePeriod\"\xd0\x01\n\x1cGetTimeSeriesBinBulkResponse\x12V\n\x0btime_series\x18\x01 \x03(\x0b\x32\x41.topos.timeseries.v1.GetTimeSeriesBinBulkResponse.TimeSeriesEntry\x1aX\n\x0fTimeSeriesEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\x34\n\x05value\x18\x02 \x01(\x0b\x32%.topos.timeseries.v1.TimeSeriesValues:\x02\x38\x01\"V\n\x1bSetTimeSeriesFeatureRequest\x12\x37\n\x07\x66\x65\x61ture\x18\x01 \x01(\x0b\x32&.topos.timeseries.v1.TimeSeriesFeature\"J\n\x17SetTimeSeriesBinRequest\x12/\n\x03\x62in\x18\x01 \x01(\x0b\x32\".topos.timeseries.v1.TimeSeriesBin\"M\n\x1a\x44\x65leteTimeSeriesBinRequest\x12/\n\x03\x62in\x18\x01 \x01(\x0b\x32\".topos.timeseries.v1.TimeSeriesBin\".\n\x1e\x44\x65leteTimeSeriesFeatureRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\"F\n\x1dListTimeSeriesFeaturesRequest\x12\x11\n\tpage_size\x18\x01 \x01(\x05\x12\x12\n\npage_token\x18\x02 \x01(\t\"s\n\x1eListTimeSeriesFeaturesResponse\x12\x38\n\x08\x66\x65\x61tures\x18\x01 \x03(\x0b\x32&.topos.timeseries.v1.TimeSeriesFeature\x12\x17\n\x0fnext_page_token\x18\x02 \x01(\t*E\n\x03\x44\x61y\x12\x07\n\x03MON\x10\x00\x12\x08\n\x04TUES\x10\x01\x12\x07\n\x03WED\x10\x02\x12\x07\n\x03THU\x10\x03\x12\x07\n\x03\x46RI\x10\x04\x12\x07\n\x03SAT\x10\x05\x12\x07\n\x03SUN\x10\x06*&\n\nTimePeriod\x12\x0b\n\x07WEEKDAY\x10\x00\x12\x0b\n\x07WEEKEND\x10\x01\x32\xd5\t\n\nTimeSeries\x12\x84\x01\n\x10GetTimeSeriesBin\x12,.topos.timeseries.v1.GetTimeSeriesBinRequest\x1a\".topos.timeseries.v1.TimeSeriesBin\"\x1e\x82\xd3\xe4\x93\x02\x18\"\x13/v1/time-series-bin:\x01*\x12\x9c\x01\n\x14GetTimeSeriesBinBulk\x12\x30.topos.timeseries.v1.GetTimeSeriesBinBulkRequest\x1a\x31.topos.timeseries.v1.GetTimeSeriesBinBulkResponse\"\x1f\x82\xd3\xe4\x93\x02\x19\"\x14/v1/time-series/bulk:\x01*\x12\xae\x01\n\x19GetTimeSeriesAveragesBulk\x12\x35.topos.timeseries.v1.GetTimeSeriesAveragesBulkRequest\x1a\x31.topos.timeseries.v1.GetTimeSeriesBinBulkResponse\"\'\x82\xd3\xe4\x93\x02!\"\x1c/v1/time-series/bulk/average:\x01*\x12\xa3\x01\n\x16ListTimeSeriesFeatures\x12\x32.topos.timeseries.v1.ListTimeSeriesFeaturesRequest\x1a\x33.topos.timeseries.v1.ListTimeSeriesFeaturesResponse\" \x82\xd3\xe4\x93\x02\x1a\x12\x18/v1/time-series-features\x12\x90\x01\n\x10SetTimeSeriesBin\x12,.topos.timeseries.v1.SetTimeSeriesBinRequest\x1a\".topos.timeseries.v1.TimeSeriesBin\"*\x82\xd3\xe4\x93\x02$\"\x1d/v1/{bin.region_id=bin/*}:set:\x03\x62in\x12\x9f\x01\n\x14SetTimeSeriesFeature\x12\x30.topos.timeseries.v1.SetTimeSeriesFeatureRequest\x1a&.topos.timeseries.v1.TimeSeriesFeature\"-\x82\xd3\xe4\x93\x02\'\"\x1c/v1/{feature.name=bin/*}:set:\x07\x66\x65\x61ture\x12\x8d\x01\n\x13\x44\x65leteTimeSeriesBin\x12/.topos.timeseries.v1.DeleteTimeSeriesBinRequest\x1a\x16.google.protobuf.Empty\"-\x82\xd3\xe4\x93\x02\'\" /v1/{bin.region_id=bin/*}:delete:\x03\x62in\x12\x84\x01\n\x17\x44\x65leteTimeSeriesFeature\x12\x33.topos.timeseries.v1.DeleteTimeSeriesFeatureRequest\x1a\x16.google.protobuf.Empty\"\x1c\x82\xd3\xe4\x93\x02\x16*\x14/v1/{name=feature/*}BKZIgithub.com/topos-ai/topos-apis/genproto/go/topos/timeseries/v1;timeseriesb\x06proto3')
  ,
  dependencies=[google_dot_api_dot_annotations__pb2.DESCRIPTOR,google_dot_protobuf_dot_empty__pb2.DESCRIPTOR,])

_DAY = _descriptor.EnumDescriptor(
  name='Day',
  full_name='topos.timeseries.v1.Day',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='MON', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='TUES', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='WED', index=2, number=2,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='THU', index=3, number=3,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='FRI', index=4, number=4,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SAT', index=5, number=5,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SUN', index=6, number=6,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=1532,
  serialized_end=1601,
)
_sym_db.RegisterEnumDescriptor(_DAY)

Day = enum_type_wrapper.EnumTypeWrapper(_DAY)
_TIMEPERIOD = _descriptor.EnumDescriptor(
  name='TimePeriod',
  full_name='topos.timeseries.v1.TimePeriod',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='WEEKDAY', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='WEEKEND', index=1, number=1,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=1603,
  serialized_end=1641,
)
_sym_db.RegisterEnumDescriptor(_TIMEPERIOD)

TimePeriod = enum_type_wrapper.EnumTypeWrapper(_TIMEPERIOD)
MON = 0
TUES = 1
WED = 2
THU = 3
FRI = 4
SAT = 5
SUN = 6
WEEKDAY = 0
WEEKEND = 1



_TIMESERIESFEATURE = _descriptor.Descriptor(
  name='TimeSeriesFeature',
  full_name='topos.timeseries.v1.TimeSeriesFeature',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='topos.timeseries.v1.TimeSeriesFeature.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='display_name', full_name='topos.timeseries.v1.TimeSeriesFeature.display_name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='description', full_name='topos.timeseries.v1.TimeSeriesFeature.description', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='source', full_name='topos.timeseries.v1.TimeSeriesFeature.source', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='sampling_interval', full_name='topos.timeseries.v1.TimeSeriesFeature.sampling_interval', index=4,
      number=5, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
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
  serialized_start=121,
  serialized_end=240,
)


_TIMESERIESBIN = _descriptor.Descriptor(
  name='TimeSeriesBin',
  full_name='topos.timeseries.v1.TimeSeriesBin',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='region_id', full_name='topos.timeseries.v1.TimeSeriesBin.region_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='feature_name', full_name='topos.timeseries.v1.TimeSeriesBin.feature_name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='day', full_name='topos.timeseries.v1.TimeSeriesBin.day', index=2,
      number=3, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='start_time', full_name='topos.timeseries.v1.TimeSeriesBin.start_time', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='topos.timeseries.v1.TimeSeriesBin.value', index=4,
      number=6, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
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
  serialized_start=243,
  serialized_end=373,
)


_TIMESERIESVALUES = _descriptor.Descriptor(
  name='TimeSeriesValues',
  full_name='topos.timeseries.v1.TimeSeriesValues',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='values', full_name='topos.timeseries.v1.TimeSeriesValues.values', index=0,
      number=1, type=3, cpp_type=2, label=3,
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
  serialized_start=375,
  serialized_end=409,
)


_GETTIMESERIESBINREQUEST = _descriptor.Descriptor(
  name='GetTimeSeriesBinRequest',
  full_name='topos.timeseries.v1.GetTimeSeriesBinRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='region_name', full_name='topos.timeseries.v1.GetTimeSeriesBinRequest.region_name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='feature_name', full_name='topos.timeseries.v1.GetTimeSeriesBinRequest.feature_name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='day', full_name='topos.timeseries.v1.GetTimeSeriesBinRequest.day', index=2,
      number=3, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='start_time', full_name='topos.timeseries.v1.GetTimeSeriesBinRequest.start_time', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='end_time', full_name='topos.timeseries.v1.GetTimeSeriesBinRequest.end_time', index=4,
      number=5, type=9, cpp_type=9, label=1,
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
  serialized_start=412,
  serialized_end=557,
)


_GETTIMESERIESBINBULKREQUEST = _descriptor.Descriptor(
  name='GetTimeSeriesBinBulkRequest',
  full_name='topos.timeseries.v1.GetTimeSeriesBinBulkRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='feature_name', full_name='topos.timeseries.v1.GetTimeSeriesBinBulkRequest.feature_name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='region_names', full_name='topos.timeseries.v1.GetTimeSeriesBinBulkRequest.region_names', index=1,
      number=2, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='day', full_name='topos.timeseries.v1.GetTimeSeriesBinBulkRequest.day', index=2,
      number=3, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='start_time', full_name='topos.timeseries.v1.GetTimeSeriesBinBulkRequest.start_time', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='end_time', full_name='topos.timeseries.v1.GetTimeSeriesBinBulkRequest.end_time', index=4,
      number=5, type=9, cpp_type=9, label=1,
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
  serialized_start=560,
  serialized_end=710,
)


_GETTIMESERIESAVERAGESBULKREQUEST = _descriptor.Descriptor(
  name='GetTimeSeriesAveragesBulkRequest',
  full_name='topos.timeseries.v1.GetTimeSeriesAveragesBulkRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='feature_name', full_name='topos.timeseries.v1.GetTimeSeriesAveragesBulkRequest.feature_name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='region_names', full_name='topos.timeseries.v1.GetTimeSeriesAveragesBulkRequest.region_names', index=1,
      number=2, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='period', full_name='topos.timeseries.v1.GetTimeSeriesAveragesBulkRequest.period', index=2,
      number=3, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
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
  serialized_start=712,
  serialized_end=839,
)


_GETTIMESERIESBINBULKRESPONSE_TIMESERIESENTRY = _descriptor.Descriptor(
  name='TimeSeriesEntry',
  full_name='topos.timeseries.v1.GetTimeSeriesBinBulkResponse.TimeSeriesEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='topos.timeseries.v1.GetTimeSeriesBinBulkResponse.TimeSeriesEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='topos.timeseries.v1.GetTimeSeriesBinBulkResponse.TimeSeriesEntry.value', index=1,
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
  serialized_options=_b('8\001'),
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=962,
  serialized_end=1050,
)

_GETTIMESERIESBINBULKRESPONSE = _descriptor.Descriptor(
  name='GetTimeSeriesBinBulkResponse',
  full_name='topos.timeseries.v1.GetTimeSeriesBinBulkResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='time_series', full_name='topos.timeseries.v1.GetTimeSeriesBinBulkResponse.time_series', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_GETTIMESERIESBINBULKRESPONSE_TIMESERIESENTRY, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=842,
  serialized_end=1050,
)


_SETTIMESERIESFEATUREREQUEST = _descriptor.Descriptor(
  name='SetTimeSeriesFeatureRequest',
  full_name='topos.timeseries.v1.SetTimeSeriesFeatureRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='feature', full_name='topos.timeseries.v1.SetTimeSeriesFeatureRequest.feature', index=0,
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
  serialized_start=1052,
  serialized_end=1138,
)


_SETTIMESERIESBINREQUEST = _descriptor.Descriptor(
  name='SetTimeSeriesBinRequest',
  full_name='topos.timeseries.v1.SetTimeSeriesBinRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='bin', full_name='topos.timeseries.v1.SetTimeSeriesBinRequest.bin', index=0,
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
  serialized_start=1140,
  serialized_end=1214,
)


_DELETETIMESERIESBINREQUEST = _descriptor.Descriptor(
  name='DeleteTimeSeriesBinRequest',
  full_name='topos.timeseries.v1.DeleteTimeSeriesBinRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='bin', full_name='topos.timeseries.v1.DeleteTimeSeriesBinRequest.bin', index=0,
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
  serialized_start=1216,
  serialized_end=1293,
)


_DELETETIMESERIESFEATUREREQUEST = _descriptor.Descriptor(
  name='DeleteTimeSeriesFeatureRequest',
  full_name='topos.timeseries.v1.DeleteTimeSeriesFeatureRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='topos.timeseries.v1.DeleteTimeSeriesFeatureRequest.name', index=0,
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
  serialized_start=1295,
  serialized_end=1341,
)


_LISTTIMESERIESFEATURESREQUEST = _descriptor.Descriptor(
  name='ListTimeSeriesFeaturesRequest',
  full_name='topos.timeseries.v1.ListTimeSeriesFeaturesRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='page_size', full_name='topos.timeseries.v1.ListTimeSeriesFeaturesRequest.page_size', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='page_token', full_name='topos.timeseries.v1.ListTimeSeriesFeaturesRequest.page_token', index=1,
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
  serialized_start=1343,
  serialized_end=1413,
)


_LISTTIMESERIESFEATURESRESPONSE = _descriptor.Descriptor(
  name='ListTimeSeriesFeaturesResponse',
  full_name='topos.timeseries.v1.ListTimeSeriesFeaturesResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='features', full_name='topos.timeseries.v1.ListTimeSeriesFeaturesResponse.features', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='next_page_token', full_name='topos.timeseries.v1.ListTimeSeriesFeaturesResponse.next_page_token', index=1,
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
  serialized_start=1415,
  serialized_end=1530,
)

_TIMESERIESBIN.fields_by_name['day'].enum_type = _DAY
_GETTIMESERIESBINREQUEST.fields_by_name['day'].enum_type = _DAY
_GETTIMESERIESBINBULKREQUEST.fields_by_name['day'].enum_type = _DAY
_GETTIMESERIESAVERAGESBULKREQUEST.fields_by_name['period'].enum_type = _TIMEPERIOD
_GETTIMESERIESBINBULKRESPONSE_TIMESERIESENTRY.fields_by_name['value'].message_type = _TIMESERIESVALUES
_GETTIMESERIESBINBULKRESPONSE_TIMESERIESENTRY.containing_type = _GETTIMESERIESBINBULKRESPONSE
_GETTIMESERIESBINBULKRESPONSE.fields_by_name['time_series'].message_type = _GETTIMESERIESBINBULKRESPONSE_TIMESERIESENTRY
_SETTIMESERIESFEATUREREQUEST.fields_by_name['feature'].message_type = _TIMESERIESFEATURE
_SETTIMESERIESBINREQUEST.fields_by_name['bin'].message_type = _TIMESERIESBIN
_DELETETIMESERIESBINREQUEST.fields_by_name['bin'].message_type = _TIMESERIESBIN
_LISTTIMESERIESFEATURESRESPONSE.fields_by_name['features'].message_type = _TIMESERIESFEATURE
DESCRIPTOR.message_types_by_name['TimeSeriesFeature'] = _TIMESERIESFEATURE
DESCRIPTOR.message_types_by_name['TimeSeriesBin'] = _TIMESERIESBIN
DESCRIPTOR.message_types_by_name['TimeSeriesValues'] = _TIMESERIESVALUES
DESCRIPTOR.message_types_by_name['GetTimeSeriesBinRequest'] = _GETTIMESERIESBINREQUEST
DESCRIPTOR.message_types_by_name['GetTimeSeriesBinBulkRequest'] = _GETTIMESERIESBINBULKREQUEST
DESCRIPTOR.message_types_by_name['GetTimeSeriesAveragesBulkRequest'] = _GETTIMESERIESAVERAGESBULKREQUEST
DESCRIPTOR.message_types_by_name['GetTimeSeriesBinBulkResponse'] = _GETTIMESERIESBINBULKRESPONSE
DESCRIPTOR.message_types_by_name['SetTimeSeriesFeatureRequest'] = _SETTIMESERIESFEATUREREQUEST
DESCRIPTOR.message_types_by_name['SetTimeSeriesBinRequest'] = _SETTIMESERIESBINREQUEST
DESCRIPTOR.message_types_by_name['DeleteTimeSeriesBinRequest'] = _DELETETIMESERIESBINREQUEST
DESCRIPTOR.message_types_by_name['DeleteTimeSeriesFeatureRequest'] = _DELETETIMESERIESFEATUREREQUEST
DESCRIPTOR.message_types_by_name['ListTimeSeriesFeaturesRequest'] = _LISTTIMESERIESFEATURESREQUEST
DESCRIPTOR.message_types_by_name['ListTimeSeriesFeaturesResponse'] = _LISTTIMESERIESFEATURESRESPONSE
DESCRIPTOR.enum_types_by_name['Day'] = _DAY
DESCRIPTOR.enum_types_by_name['TimePeriod'] = _TIMEPERIOD
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

TimeSeriesFeature = _reflection.GeneratedProtocolMessageType('TimeSeriesFeature', (_message.Message,), {
  'DESCRIPTOR' : _TIMESERIESFEATURE,
  '__module__' : 'topos.timeseries.v1.time_series_pb2'
  # @@protoc_insertion_point(class_scope:topos.timeseries.v1.TimeSeriesFeature)
  })
_sym_db.RegisterMessage(TimeSeriesFeature)

TimeSeriesBin = _reflection.GeneratedProtocolMessageType('TimeSeriesBin', (_message.Message,), {
  'DESCRIPTOR' : _TIMESERIESBIN,
  '__module__' : 'topos.timeseries.v1.time_series_pb2'
  # @@protoc_insertion_point(class_scope:topos.timeseries.v1.TimeSeriesBin)
  })
_sym_db.RegisterMessage(TimeSeriesBin)

TimeSeriesValues = _reflection.GeneratedProtocolMessageType('TimeSeriesValues', (_message.Message,), {
  'DESCRIPTOR' : _TIMESERIESVALUES,
  '__module__' : 'topos.timeseries.v1.time_series_pb2'
  # @@protoc_insertion_point(class_scope:topos.timeseries.v1.TimeSeriesValues)
  })
_sym_db.RegisterMessage(TimeSeriesValues)

GetTimeSeriesBinRequest = _reflection.GeneratedProtocolMessageType('GetTimeSeriesBinRequest', (_message.Message,), {
  'DESCRIPTOR' : _GETTIMESERIESBINREQUEST,
  '__module__' : 'topos.timeseries.v1.time_series_pb2'
  # @@protoc_insertion_point(class_scope:topos.timeseries.v1.GetTimeSeriesBinRequest)
  })
_sym_db.RegisterMessage(GetTimeSeriesBinRequest)

GetTimeSeriesBinBulkRequest = _reflection.GeneratedProtocolMessageType('GetTimeSeriesBinBulkRequest', (_message.Message,), {
  'DESCRIPTOR' : _GETTIMESERIESBINBULKREQUEST,
  '__module__' : 'topos.timeseries.v1.time_series_pb2'
  # @@protoc_insertion_point(class_scope:topos.timeseries.v1.GetTimeSeriesBinBulkRequest)
  })
_sym_db.RegisterMessage(GetTimeSeriesBinBulkRequest)

GetTimeSeriesAveragesBulkRequest = _reflection.GeneratedProtocolMessageType('GetTimeSeriesAveragesBulkRequest', (_message.Message,), {
  'DESCRIPTOR' : _GETTIMESERIESAVERAGESBULKREQUEST,
  '__module__' : 'topos.timeseries.v1.time_series_pb2'
  # @@protoc_insertion_point(class_scope:topos.timeseries.v1.GetTimeSeriesAveragesBulkRequest)
  })
_sym_db.RegisterMessage(GetTimeSeriesAveragesBulkRequest)

GetTimeSeriesBinBulkResponse = _reflection.GeneratedProtocolMessageType('GetTimeSeriesBinBulkResponse', (_message.Message,), {

  'TimeSeriesEntry' : _reflection.GeneratedProtocolMessageType('TimeSeriesEntry', (_message.Message,), {
    'DESCRIPTOR' : _GETTIMESERIESBINBULKRESPONSE_TIMESERIESENTRY,
    '__module__' : 'topos.timeseries.v1.time_series_pb2'
    # @@protoc_insertion_point(class_scope:topos.timeseries.v1.GetTimeSeriesBinBulkResponse.TimeSeriesEntry)
    })
  ,
  'DESCRIPTOR' : _GETTIMESERIESBINBULKRESPONSE,
  '__module__' : 'topos.timeseries.v1.time_series_pb2'
  # @@protoc_insertion_point(class_scope:topos.timeseries.v1.GetTimeSeriesBinBulkResponse)
  })
_sym_db.RegisterMessage(GetTimeSeriesBinBulkResponse)
_sym_db.RegisterMessage(GetTimeSeriesBinBulkResponse.TimeSeriesEntry)

SetTimeSeriesFeatureRequest = _reflection.GeneratedProtocolMessageType('SetTimeSeriesFeatureRequest', (_message.Message,), {
  'DESCRIPTOR' : _SETTIMESERIESFEATUREREQUEST,
  '__module__' : 'topos.timeseries.v1.time_series_pb2'
  # @@protoc_insertion_point(class_scope:topos.timeseries.v1.SetTimeSeriesFeatureRequest)
  })
_sym_db.RegisterMessage(SetTimeSeriesFeatureRequest)

SetTimeSeriesBinRequest = _reflection.GeneratedProtocolMessageType('SetTimeSeriesBinRequest', (_message.Message,), {
  'DESCRIPTOR' : _SETTIMESERIESBINREQUEST,
  '__module__' : 'topos.timeseries.v1.time_series_pb2'
  # @@protoc_insertion_point(class_scope:topos.timeseries.v1.SetTimeSeriesBinRequest)
  })
_sym_db.RegisterMessage(SetTimeSeriesBinRequest)

DeleteTimeSeriesBinRequest = _reflection.GeneratedProtocolMessageType('DeleteTimeSeriesBinRequest', (_message.Message,), {
  'DESCRIPTOR' : _DELETETIMESERIESBINREQUEST,
  '__module__' : 'topos.timeseries.v1.time_series_pb2'
  # @@protoc_insertion_point(class_scope:topos.timeseries.v1.DeleteTimeSeriesBinRequest)
  })
_sym_db.RegisterMessage(DeleteTimeSeriesBinRequest)

DeleteTimeSeriesFeatureRequest = _reflection.GeneratedProtocolMessageType('DeleteTimeSeriesFeatureRequest', (_message.Message,), {
  'DESCRIPTOR' : _DELETETIMESERIESFEATUREREQUEST,
  '__module__' : 'topos.timeseries.v1.time_series_pb2'
  # @@protoc_insertion_point(class_scope:topos.timeseries.v1.DeleteTimeSeriesFeatureRequest)
  })
_sym_db.RegisterMessage(DeleteTimeSeriesFeatureRequest)

ListTimeSeriesFeaturesRequest = _reflection.GeneratedProtocolMessageType('ListTimeSeriesFeaturesRequest', (_message.Message,), {
  'DESCRIPTOR' : _LISTTIMESERIESFEATURESREQUEST,
  '__module__' : 'topos.timeseries.v1.time_series_pb2'
  # @@protoc_insertion_point(class_scope:topos.timeseries.v1.ListTimeSeriesFeaturesRequest)
  })
_sym_db.RegisterMessage(ListTimeSeriesFeaturesRequest)

ListTimeSeriesFeaturesResponse = _reflection.GeneratedProtocolMessageType('ListTimeSeriesFeaturesResponse', (_message.Message,), {
  'DESCRIPTOR' : _LISTTIMESERIESFEATURESRESPONSE,
  '__module__' : 'topos.timeseries.v1.time_series_pb2'
  # @@protoc_insertion_point(class_scope:topos.timeseries.v1.ListTimeSeriesFeaturesResponse)
  })
_sym_db.RegisterMessage(ListTimeSeriesFeaturesResponse)


DESCRIPTOR._options = None
_GETTIMESERIESBINBULKRESPONSE_TIMESERIESENTRY._options = None

_TIMESERIES = _descriptor.ServiceDescriptor(
  name='TimeSeries',
  full_name='topos.timeseries.v1.TimeSeries',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=1644,
  serialized_end=2881,
  methods=[
  _descriptor.MethodDescriptor(
    name='GetTimeSeriesBin',
    full_name='topos.timeseries.v1.TimeSeries.GetTimeSeriesBin',
    index=0,
    containing_service=None,
    input_type=_GETTIMESERIESBINREQUEST,
    output_type=_TIMESERIESBIN,
    serialized_options=_b('\202\323\344\223\002\030\"\023/v1/time-series-bin:\001*'),
  ),
  _descriptor.MethodDescriptor(
    name='GetTimeSeriesBinBulk',
    full_name='topos.timeseries.v1.TimeSeries.GetTimeSeriesBinBulk',
    index=1,
    containing_service=None,
    input_type=_GETTIMESERIESBINBULKREQUEST,
    output_type=_GETTIMESERIESBINBULKRESPONSE,
    serialized_options=_b('\202\323\344\223\002\031\"\024/v1/time-series/bulk:\001*'),
  ),
  _descriptor.MethodDescriptor(
    name='GetTimeSeriesAveragesBulk',
    full_name='topos.timeseries.v1.TimeSeries.GetTimeSeriesAveragesBulk',
    index=2,
    containing_service=None,
    input_type=_GETTIMESERIESAVERAGESBULKREQUEST,
    output_type=_GETTIMESERIESBINBULKRESPONSE,
    serialized_options=_b('\202\323\344\223\002!\"\034/v1/time-series/bulk/average:\001*'),
  ),
  _descriptor.MethodDescriptor(
    name='ListTimeSeriesFeatures',
    full_name='topos.timeseries.v1.TimeSeries.ListTimeSeriesFeatures',
    index=3,
    containing_service=None,
    input_type=_LISTTIMESERIESFEATURESREQUEST,
    output_type=_LISTTIMESERIESFEATURESRESPONSE,
    serialized_options=_b('\202\323\344\223\002\032\022\030/v1/time-series-features'),
  ),
  _descriptor.MethodDescriptor(
    name='SetTimeSeriesBin',
    full_name='topos.timeseries.v1.TimeSeries.SetTimeSeriesBin',
    index=4,
    containing_service=None,
    input_type=_SETTIMESERIESBINREQUEST,
    output_type=_TIMESERIESBIN,
    serialized_options=_b('\202\323\344\223\002$\"\035/v1/{bin.region_id=bin/*}:set:\003bin'),
  ),
  _descriptor.MethodDescriptor(
    name='SetTimeSeriesFeature',
    full_name='topos.timeseries.v1.TimeSeries.SetTimeSeriesFeature',
    index=5,
    containing_service=None,
    input_type=_SETTIMESERIESFEATUREREQUEST,
    output_type=_TIMESERIESFEATURE,
    serialized_options=_b('\202\323\344\223\002\'\"\034/v1/{feature.name=bin/*}:set:\007feature'),
  ),
  _descriptor.MethodDescriptor(
    name='DeleteTimeSeriesBin',
    full_name='topos.timeseries.v1.TimeSeries.DeleteTimeSeriesBin',
    index=6,
    containing_service=None,
    input_type=_DELETETIMESERIESBINREQUEST,
    output_type=google_dot_protobuf_dot_empty__pb2._EMPTY,
    serialized_options=_b('\202\323\344\223\002\'\" /v1/{bin.region_id=bin/*}:delete:\003bin'),
  ),
  _descriptor.MethodDescriptor(
    name='DeleteTimeSeriesFeature',
    full_name='topos.timeseries.v1.TimeSeries.DeleteTimeSeriesFeature',
    index=7,
    containing_service=None,
    input_type=_DELETETIMESERIESFEATUREREQUEST,
    output_type=google_dot_protobuf_dot_empty__pb2._EMPTY,
    serialized_options=_b('\202\323\344\223\002\026*\024/v1/{name=feature/*}'),
  ),
])
_sym_db.RegisterServiceDescriptor(_TIMESERIES)

DESCRIPTOR.services_by_name['TimeSeries'] = _TIMESERIES

# @@protoc_insertion_point(module_scope)
