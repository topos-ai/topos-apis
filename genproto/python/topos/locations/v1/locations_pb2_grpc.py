# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

from topos.locations.v1 import locations_pb2 as topos_dot_locations_dot_v1_dot_locations__pb2


class LocationsStub(object):
  # missing associated documentation comment in .proto file
  pass

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.GetRegion = channel.unary_unary(
        '/topos.locations.v1.Locations/GetRegion',
        request_serializer=topos_dot_locations_dot_v1_dot_locations__pb2.GetRegionRequest.SerializeToString,
        response_deserializer=topos_dot_locations_dot_v1_dot_locations__pb2.Region.FromString,
        )
    self.SearchRegions = channel.unary_unary(
        '/topos.locations.v1.Locations/SearchRegions',
        request_serializer=topos_dot_locations_dot_v1_dot_locations__pb2.SearchRegionsRequest.SerializeToString,
        response_deserializer=topos_dot_locations_dot_v1_dot_locations__pb2.SearchRegionsResponse.FromString,
        )


class LocationsServicer(object):
  # missing associated documentation comment in .proto file
  pass

  def GetRegion(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def SearchRegions(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_LocationsServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'GetRegion': grpc.unary_unary_rpc_method_handler(
          servicer.GetRegion,
          request_deserializer=topos_dot_locations_dot_v1_dot_locations__pb2.GetRegionRequest.FromString,
          response_serializer=topos_dot_locations_dot_v1_dot_locations__pb2.Region.SerializeToString,
      ),
      'SearchRegions': grpc.unary_unary_rpc_method_handler(
          servicer.SearchRegions,
          request_deserializer=topos_dot_locations_dot_v1_dot_locations__pb2.SearchRegionsRequest.FromString,
          response_serializer=topos_dot_locations_dot_v1_dot_locations__pb2.SearchRegionsResponse.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'topos.locations.v1.Locations', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))
