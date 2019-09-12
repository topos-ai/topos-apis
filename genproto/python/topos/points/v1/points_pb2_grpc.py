# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from topos.points.v1 import points_pb2 as topos_dot_points_dot_v1_dot_points__pb2


class PointsStub(object):
  # missing associated documentation comment in .proto file
  pass

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.GetBrand = channel.unary_unary(
        '/topos.points.v1.Points/GetBrand',
        request_serializer=topos_dot_points_dot_v1_dot_points__pb2.GetBrandRequest.SerializeToString,
        response_deserializer=topos_dot_points_dot_v1_dot_points__pb2.Brand.FromString,
        )
    self.ListBrands = channel.unary_unary(
        '/topos.points.v1.Points/ListBrands',
        request_serializer=topos_dot_points_dot_v1_dot_points__pb2.ListBrandsRequest.SerializeToString,
        response_deserializer=topos_dot_points_dot_v1_dot_points__pb2.ListBrandsResponse.FromString,
        )
    self.SetBrand = channel.unary_unary(
        '/topos.points.v1.Points/SetBrand',
        request_serializer=topos_dot_points_dot_v1_dot_points__pb2.SetBrandRequest.SerializeToString,
        response_deserializer=topos_dot_points_dot_v1_dot_points__pb2.Brand.FromString,
        )
    self.DeleteBrand = channel.unary_unary(
        '/topos.points.v1.Points/DeleteBrand',
        request_serializer=topos_dot_points_dot_v1_dot_points__pb2.DeleteBrandRequest.SerializeToString,
        response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
        )
    self.GetTag = channel.unary_unary(
        '/topos.points.v1.Points/GetTag',
        request_serializer=topos_dot_points_dot_v1_dot_points__pb2.GetTagRequest.SerializeToString,
        response_deserializer=topos_dot_points_dot_v1_dot_points__pb2.Tag.FromString,
        )
    self.ListTags = channel.unary_unary(
        '/topos.points.v1.Points/ListTags',
        request_serializer=topos_dot_points_dot_v1_dot_points__pb2.ListTagsRequest.SerializeToString,
        response_deserializer=topos_dot_points_dot_v1_dot_points__pb2.ListTagsResponse.FromString,
        )
    self.SetTag = channel.unary_unary(
        '/topos.points.v1.Points/SetTag',
        request_serializer=topos_dot_points_dot_v1_dot_points__pb2.SetTagRequest.SerializeToString,
        response_deserializer=topos_dot_points_dot_v1_dot_points__pb2.Tag.FromString,
        )
    self.DeleteTag = channel.unary_unary(
        '/topos.points.v1.Points/DeleteTag',
        request_serializer=topos_dot_points_dot_v1_dot_points__pb2.DeleteTagRequest.SerializeToString,
        response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
        )
    self.GetPoint = channel.unary_unary(
        '/topos.points.v1.Points/GetPoint',
        request_serializer=topos_dot_points_dot_v1_dot_points__pb2.GetPointRequest.SerializeToString,
        response_deserializer=topos_dot_points_dot_v1_dot_points__pb2.Point.FromString,
        )
    self.SearchPoints = channel.unary_unary(
        '/topos.points.v1.Points/SearchPoints',
        request_serializer=topos_dot_points_dot_v1_dot_points__pb2.SearchPointsRequest.SerializeToString,
        response_deserializer=topos_dot_points_dot_v1_dot_points__pb2.SearchPointsResponse.FromString,
        )
    self.PolygonSearchPoints = channel.stream_unary(
        '/topos.points.v1.Points/PolygonSearchPoints',
        request_serializer=topos_dot_points_dot_v1_dot_points__pb2.PolygonSearchPointsRequest.SerializeToString,
        response_deserializer=topos_dot_points_dot_v1_dot_points__pb2.PolygonSearchPointsResponse.FromString,
        )
    self.SetPoint = channel.unary_unary(
        '/topos.points.v1.Points/SetPoint',
        request_serializer=topos_dot_points_dot_v1_dot_points__pb2.SetPointRequest.SerializeToString,
        response_deserializer=topos_dot_points_dot_v1_dot_points__pb2.Point.FromString,
        )
    self.DeletePoint = channel.unary_unary(
        '/topos.points.v1.Points/DeletePoint',
        request_serializer=topos_dot_points_dot_v1_dot_points__pb2.DeletePointRequest.SerializeToString,
        response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
        )
    self.CountBrandPoints = channel.unary_unary(
        '/topos.points.v1.Points/CountBrandPoints',
        request_serializer=topos_dot_points_dot_v1_dot_points__pb2.CountBrandPointsRequest.SerializeToString,
        response_deserializer=topos_dot_points_dot_v1_dot_points__pb2.CountBrandPointsResponse.FromString,
        )
    self.PolygonCountBrandPoints = channel.stream_unary(
        '/topos.points.v1.Points/PolygonCountBrandPoints',
        request_serializer=topos_dot_points_dot_v1_dot_points__pb2.PolygonCountBrandPointsRequest.SerializeToString,
        response_deserializer=topos_dot_points_dot_v1_dot_points__pb2.PolygonCountBrandPointsResponse.FromString,
        )
    self.CountTagPoints = channel.unary_unary(
        '/topos.points.v1.Points/CountTagPoints',
        request_serializer=topos_dot_points_dot_v1_dot_points__pb2.CountTagPointsRequest.SerializeToString,
        response_deserializer=topos_dot_points_dot_v1_dot_points__pb2.CountTagPointsResponse.FromString,
        )
    self.PolygonCountTagPoints = channel.stream_unary(
        '/topos.points.v1.Points/PolygonCountTagPoints',
        request_serializer=topos_dot_points_dot_v1_dot_points__pb2.PolygonCountTagPointsRequest.SerializeToString,
        response_deserializer=topos_dot_points_dot_v1_dot_points__pb2.PolygonCountTagPointsResponse.FromString,
        )
    self.SearchRegions = channel.unary_unary(
        '/topos.points.v1.Points/SearchRegions',
        request_serializer=topos_dot_points_dot_v1_dot_points__pb2.SearchRegionsRequest.SerializeToString,
        response_deserializer=topos_dot_points_dot_v1_dot_points__pb2.SearchRegionsResponse.FromString,
        )


class PointsServicer(object):
  # missing associated documentation comment in .proto file
  pass

  def GetBrand(self, request, context):
    """Gets a brand.
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def ListBrands(self, request, context):
    """Lists brands.
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def SetBrand(self, request, context):
    """Sets a brand.
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def DeleteBrand(self, request, context):
    """Deletes a brand.
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def GetTag(self, request, context):
    """Gets a tag.
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def ListTags(self, request, context):
    """Lists tags.
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def SetTag(self, request, context):
    """Sets a tag.
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def DeleteTag(self, request, context):
    """Deletes a tag.
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def GetPoint(self, request, context):
    """Gets a point.
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def SearchPoints(self, request, context):
    """Search points.
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def PolygonSearchPoints(self, request_iterator, context):
    """Search points.
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def SetPoint(self, request, context):
    """Sets a point.
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def DeletePoint(self, request, context):
    """Deletes a point.
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def CountBrandPoints(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def PolygonCountBrandPoints(self, request_iterator, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def CountTagPoints(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def PolygonCountTagPoints(self, request_iterator, context):
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


def add_PointsServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'GetBrand': grpc.unary_unary_rpc_method_handler(
          servicer.GetBrand,
          request_deserializer=topos_dot_points_dot_v1_dot_points__pb2.GetBrandRequest.FromString,
          response_serializer=topos_dot_points_dot_v1_dot_points__pb2.Brand.SerializeToString,
      ),
      'ListBrands': grpc.unary_unary_rpc_method_handler(
          servicer.ListBrands,
          request_deserializer=topos_dot_points_dot_v1_dot_points__pb2.ListBrandsRequest.FromString,
          response_serializer=topos_dot_points_dot_v1_dot_points__pb2.ListBrandsResponse.SerializeToString,
      ),
      'SetBrand': grpc.unary_unary_rpc_method_handler(
          servicer.SetBrand,
          request_deserializer=topos_dot_points_dot_v1_dot_points__pb2.SetBrandRequest.FromString,
          response_serializer=topos_dot_points_dot_v1_dot_points__pb2.Brand.SerializeToString,
      ),
      'DeleteBrand': grpc.unary_unary_rpc_method_handler(
          servicer.DeleteBrand,
          request_deserializer=topos_dot_points_dot_v1_dot_points__pb2.DeleteBrandRequest.FromString,
          response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
      ),
      'GetTag': grpc.unary_unary_rpc_method_handler(
          servicer.GetTag,
          request_deserializer=topos_dot_points_dot_v1_dot_points__pb2.GetTagRequest.FromString,
          response_serializer=topos_dot_points_dot_v1_dot_points__pb2.Tag.SerializeToString,
      ),
      'ListTags': grpc.unary_unary_rpc_method_handler(
          servicer.ListTags,
          request_deserializer=topos_dot_points_dot_v1_dot_points__pb2.ListTagsRequest.FromString,
          response_serializer=topos_dot_points_dot_v1_dot_points__pb2.ListTagsResponse.SerializeToString,
      ),
      'SetTag': grpc.unary_unary_rpc_method_handler(
          servicer.SetTag,
          request_deserializer=topos_dot_points_dot_v1_dot_points__pb2.SetTagRequest.FromString,
          response_serializer=topos_dot_points_dot_v1_dot_points__pb2.Tag.SerializeToString,
      ),
      'DeleteTag': grpc.unary_unary_rpc_method_handler(
          servicer.DeleteTag,
          request_deserializer=topos_dot_points_dot_v1_dot_points__pb2.DeleteTagRequest.FromString,
          response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
      ),
      'GetPoint': grpc.unary_unary_rpc_method_handler(
          servicer.GetPoint,
          request_deserializer=topos_dot_points_dot_v1_dot_points__pb2.GetPointRequest.FromString,
          response_serializer=topos_dot_points_dot_v1_dot_points__pb2.Point.SerializeToString,
      ),
      'SearchPoints': grpc.unary_unary_rpc_method_handler(
          servicer.SearchPoints,
          request_deserializer=topos_dot_points_dot_v1_dot_points__pb2.SearchPointsRequest.FromString,
          response_serializer=topos_dot_points_dot_v1_dot_points__pb2.SearchPointsResponse.SerializeToString,
      ),
      'PolygonSearchPoints': grpc.stream_unary_rpc_method_handler(
          servicer.PolygonSearchPoints,
          request_deserializer=topos_dot_points_dot_v1_dot_points__pb2.PolygonSearchPointsRequest.FromString,
          response_serializer=topos_dot_points_dot_v1_dot_points__pb2.PolygonSearchPointsResponse.SerializeToString,
      ),
      'SetPoint': grpc.unary_unary_rpc_method_handler(
          servicer.SetPoint,
          request_deserializer=topos_dot_points_dot_v1_dot_points__pb2.SetPointRequest.FromString,
          response_serializer=topos_dot_points_dot_v1_dot_points__pb2.Point.SerializeToString,
      ),
      'DeletePoint': grpc.unary_unary_rpc_method_handler(
          servicer.DeletePoint,
          request_deserializer=topos_dot_points_dot_v1_dot_points__pb2.DeletePointRequest.FromString,
          response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
      ),
      'CountBrandPoints': grpc.unary_unary_rpc_method_handler(
          servicer.CountBrandPoints,
          request_deserializer=topos_dot_points_dot_v1_dot_points__pb2.CountBrandPointsRequest.FromString,
          response_serializer=topos_dot_points_dot_v1_dot_points__pb2.CountBrandPointsResponse.SerializeToString,
      ),
      'PolygonCountBrandPoints': grpc.stream_unary_rpc_method_handler(
          servicer.PolygonCountBrandPoints,
          request_deserializer=topos_dot_points_dot_v1_dot_points__pb2.PolygonCountBrandPointsRequest.FromString,
          response_serializer=topos_dot_points_dot_v1_dot_points__pb2.PolygonCountBrandPointsResponse.SerializeToString,
      ),
      'CountTagPoints': grpc.unary_unary_rpc_method_handler(
          servicer.CountTagPoints,
          request_deserializer=topos_dot_points_dot_v1_dot_points__pb2.CountTagPointsRequest.FromString,
          response_serializer=topos_dot_points_dot_v1_dot_points__pb2.CountTagPointsResponse.SerializeToString,
      ),
      'PolygonCountTagPoints': grpc.stream_unary_rpc_method_handler(
          servicer.PolygonCountTagPoints,
          request_deserializer=topos_dot_points_dot_v1_dot_points__pb2.PolygonCountTagPointsRequest.FromString,
          response_serializer=topos_dot_points_dot_v1_dot_points__pb2.PolygonCountTagPointsResponse.SerializeToString,
      ),
      'SearchRegions': grpc.unary_unary_rpc_method_handler(
          servicer.SearchRegions,
          request_deserializer=topos_dot_points_dot_v1_dot_points__pb2.SearchRegionsRequest.FromString,
          response_serializer=topos_dot_points_dot_v1_dot_points__pb2.SearchRegionsResponse.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'topos.points.v1.Points', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))
