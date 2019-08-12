# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

from topos.topiary.v1 import topiary_pb2 as topos_dot_topiary_dot_v1_dot_topiary__pb2


class TopiaryStub(object):
  # missing associated documentation comment in .proto file
  pass

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.ListIDs = channel.unary_unary(
        '/topos.topiary.v1.Topiary/ListIDs',
        request_serializer=topos_dot_topiary_dot_v1_dot_topiary__pb2.ListIDsRequest.SerializeToString,
        response_deserializer=topos_dot_topiary_dot_v1_dot_topiary__pb2.ListIDsResponse.FromString,
        )
    self.SetIDPosition = channel.unary_unary(
        '/topos.topiary.v1.Topiary/SetIDPosition',
        request_serializer=topos_dot_topiary_dot_v1_dot_topiary__pb2.SetIDPositionRequest.SerializeToString,
        response_deserializer=topos_dot_topiary_dot_v1_dot_topiary__pb2.SetIDPositionResponse.FromString,
        )
    self.DeleteID = channel.unary_unary(
        '/topos.topiary.v1.Topiary/DeleteID',
        request_serializer=topos_dot_topiary_dot_v1_dot_topiary__pb2.DeleteIDRequest.SerializeToString,
        response_deserializer=topos_dot_topiary_dot_v1_dot_topiary__pb2.DeleteIDResponse.FromString,
        )
    self.GetIDKeyValues = channel.unary_unary(
        '/topos.topiary.v1.Topiary/GetIDKeyValues',
        request_serializer=topos_dot_topiary_dot_v1_dot_topiary__pb2.GetIDKeyValuesRequest.SerializeToString,
        response_deserializer=topos_dot_topiary_dot_v1_dot_topiary__pb2.GetIDKeyValuesResponse.FromString,
        )
    self.SetIDKeyValue = channel.unary_unary(
        '/topos.topiary.v1.Topiary/SetIDKeyValue',
        request_serializer=topos_dot_topiary_dot_v1_dot_topiary__pb2.SetIDKeyValueRequest.SerializeToString,
        response_deserializer=topos_dot_topiary_dot_v1_dot_topiary__pb2.SetIDKeyValueResponse.FromString,
        )
    self.SearchIDs = channel.stream_unary(
        '/topos.topiary.v1.Topiary/SearchIDs',
        request_serializer=topos_dot_topiary_dot_v1_dot_topiary__pb2.SearchIDsRequest.SerializeToString,
        response_deserializer=topos_dot_topiary_dot_v1_dot_topiary__pb2.SearchIDsResponse.FromString,
        )
    self.CountIDs = channel.stream_unary(
        '/topos.topiary.v1.Topiary/CountIDs',
        request_serializer=topos_dot_topiary_dot_v1_dot_topiary__pb2.CountIDsRequest.SerializeToString,
        response_deserializer=topos_dot_topiary_dot_v1_dot_topiary__pb2.CountIDsResponse.FromString,
        )


class TopiaryServicer(object):
  # missing associated documentation comment in .proto file
  pass

  def ListIDs(self, request, context):
    """Lists IDs.
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def SetIDPosition(self, request, context):
    """Set a position key balue.
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def DeleteID(self, request, context):
    """Delete an ID.
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def GetIDKeyValues(self, request, context):
    """Set a position.
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def SetIDKeyValue(self, request, context):
    """Set a position.
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def SearchIDs(self, request_iterator, context):
    """Searches IDs.
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def CountIDs(self, request_iterator, context):
    """Counts IDs.
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_TopiaryServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'ListIDs': grpc.unary_unary_rpc_method_handler(
          servicer.ListIDs,
          request_deserializer=topos_dot_topiary_dot_v1_dot_topiary__pb2.ListIDsRequest.FromString,
          response_serializer=topos_dot_topiary_dot_v1_dot_topiary__pb2.ListIDsResponse.SerializeToString,
      ),
      'SetIDPosition': grpc.unary_unary_rpc_method_handler(
          servicer.SetIDPosition,
          request_deserializer=topos_dot_topiary_dot_v1_dot_topiary__pb2.SetIDPositionRequest.FromString,
          response_serializer=topos_dot_topiary_dot_v1_dot_topiary__pb2.SetIDPositionResponse.SerializeToString,
      ),
      'DeleteID': grpc.unary_unary_rpc_method_handler(
          servicer.DeleteID,
          request_deserializer=topos_dot_topiary_dot_v1_dot_topiary__pb2.DeleteIDRequest.FromString,
          response_serializer=topos_dot_topiary_dot_v1_dot_topiary__pb2.DeleteIDResponse.SerializeToString,
      ),
      'GetIDKeyValues': grpc.unary_unary_rpc_method_handler(
          servicer.GetIDKeyValues,
          request_deserializer=topos_dot_topiary_dot_v1_dot_topiary__pb2.GetIDKeyValuesRequest.FromString,
          response_serializer=topos_dot_topiary_dot_v1_dot_topiary__pb2.GetIDKeyValuesResponse.SerializeToString,
      ),
      'SetIDKeyValue': grpc.unary_unary_rpc_method_handler(
          servicer.SetIDKeyValue,
          request_deserializer=topos_dot_topiary_dot_v1_dot_topiary__pb2.SetIDKeyValueRequest.FromString,
          response_serializer=topos_dot_topiary_dot_v1_dot_topiary__pb2.SetIDKeyValueResponse.SerializeToString,
      ),
      'SearchIDs': grpc.stream_unary_rpc_method_handler(
          servicer.SearchIDs,
          request_deserializer=topos_dot_topiary_dot_v1_dot_topiary__pb2.SearchIDsRequest.FromString,
          response_serializer=topos_dot_topiary_dot_v1_dot_topiary__pb2.SearchIDsResponse.SerializeToString,
      ),
      'CountIDs': grpc.stream_unary_rpc_method_handler(
          servicer.CountIDs,
          request_deserializer=topos_dot_topiary_dot_v1_dot_topiary__pb2.CountIDsRequest.FromString,
          response_serializer=topos_dot_topiary_dot_v1_dot_topiary__pb2.CountIDsResponse.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'topos.topiary.v1.Topiary', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))
