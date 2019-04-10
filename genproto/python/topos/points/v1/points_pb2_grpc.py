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
    self.SetBrand = channel.unary_unary(
        '/topos.points.v1.Points/SetBrand',
        request_serializer=topos_dot_points_dot_v1_dot_points__pb2.SetBrandRequest.SerializeToString,
        response_deserializer=topos_dot_points_dot_v1_dot_points__pb2.Brand.FromString,
        )
    self.ListBrands = channel.unary_unary(
        '/topos.points.v1.Points/ListBrands',
        request_serializer=topos_dot_points_dot_v1_dot_points__pb2.ListBrandsRequest.SerializeToString,
        response_deserializer=topos_dot_points_dot_v1_dot_points__pb2.ListBrandsResponse.FromString,
        )
    self.DeleteBrand = channel.unary_unary(
        '/topos.points.v1.Points/DeleteBrand',
        request_serializer=topos_dot_points_dot_v1_dot_points__pb2.DeleteBrandRequest.SerializeToString,
        response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
        )
    self.SearchPoints = channel.unary_unary(
        '/topos.points.v1.Points/SearchPoints',
        request_serializer=topos_dot_points_dot_v1_dot_points__pb2.SearchPointsRequest.SerializeToString,
        response_deserializer=topos_dot_points_dot_v1_dot_points__pb2.SearchPointsResponse.FromString,
        )
    self.GetPoint = channel.unary_unary(
        '/topos.points.v1.Points/GetPoint',
        request_serializer=topos_dot_points_dot_v1_dot_points__pb2.GetPointRequest.SerializeToString,
        response_deserializer=topos_dot_points_dot_v1_dot_points__pb2.Point.FromString,
        )
    self.CreatePointWithPointSources = channel.unary_unary(
        '/topos.points.v1.Points/CreatePointWithPointSources',
        request_serializer=topos_dot_points_dot_v1_dot_points__pb2.CreatePointWithPointSourcesRequest.SerializeToString,
        response_deserializer=topos_dot_points_dot_v1_dot_points__pb2.CreatePointWithPointSourcesResponse.FromString,
        )
    self.UpdatePoint = channel.unary_unary(
        '/topos.points.v1.Points/UpdatePoint',
        request_serializer=topos_dot_points_dot_v1_dot_points__pb2.UpdatePointRequest.SerializeToString,
        response_deserializer=topos_dot_points_dot_v1_dot_points__pb2.Point.FromString,
        )
    self.DeletePoint = channel.unary_unary(
        '/topos.points.v1.Points/DeletePoint',
        request_serializer=topos_dot_points_dot_v1_dot_points__pb2.DeletePointRequest.SerializeToString,
        response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
        )
    self.GetTag = channel.unary_unary(
        '/topos.points.v1.Points/GetTag',
        request_serializer=topos_dot_points_dot_v1_dot_points__pb2.GetTagRequest.SerializeToString,
        response_deserializer=topos_dot_points_dot_v1_dot_points__pb2.Tag.FromString,
        )
    self.SetTag = channel.unary_unary(
        '/topos.points.v1.Points/SetTag',
        request_serializer=topos_dot_points_dot_v1_dot_points__pb2.SetTagRequest.SerializeToString,
        response_deserializer=topos_dot_points_dot_v1_dot_points__pb2.Tag.FromString,
        )
    self.ListTags = channel.unary_unary(
        '/topos.points.v1.Points/ListTags',
        request_serializer=topos_dot_points_dot_v1_dot_points__pb2.ListTagsRequest.SerializeToString,
        response_deserializer=topos_dot_points_dot_v1_dot_points__pb2.ListTagsResponse.FromString,
        )
    self.DeleteTag = channel.unary_unary(
        '/topos.points.v1.Points/DeleteTag',
        request_serializer=topos_dot_points_dot_v1_dot_points__pb2.DeleteTagRequest.SerializeToString,
        response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
        )
    self.GetPointSource = channel.unary_unary(
        '/topos.points.v1.Points/GetPointSource',
        request_serializer=topos_dot_points_dot_v1_dot_points__pb2.GetPointSourceRequest.SerializeToString,
        response_deserializer=topos_dot_points_dot_v1_dot_points__pb2.PointSource.FromString,
        )
    self.SetPointSource = channel.unary_unary(
        '/topos.points.v1.Points/SetPointSource',
        request_serializer=topos_dot_points_dot_v1_dot_points__pb2.SetPointSourceRequest.SerializeToString,
        response_deserializer=topos_dot_points_dot_v1_dot_points__pb2.PointSource.FromString,
        )
    self.DeletePointSource = channel.unary_unary(
        '/topos.points.v1.Points/DeletePointSource',
        request_serializer=topos_dot_points_dot_v1_dot_points__pb2.DeletePointSourceRequest.SerializeToString,
        response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
        )
    self.SetBrandRegionScore = channel.unary_unary(
        '/topos.points.v1.Points/SetBrandRegionScore',
        request_serializer=topos_dot_points_dot_v1_dot_points__pb2.SetBrandRegionScoreRequest.SerializeToString,
        response_deserializer=topos_dot_points_dot_v1_dot_points__pb2.BrandRegionScore.FromString,
        )
    self.SetTagRegionScore = channel.unary_unary(
        '/topos.points.v1.Points/SetTagRegionScore',
        request_serializer=topos_dot_points_dot_v1_dot_points__pb2.SetTagRegionScoreRequest.SerializeToString,
        response_deserializer=topos_dot_points_dot_v1_dot_points__pb2.TagRegionScore.FromString,
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

  def SetBrand(self, request, context):
    """Sets a brand.
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

  def DeleteBrand(self, request, context):
    """Deletes a brand.
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

  def GetPoint(self, request, context):
    """Gets a point.
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def CreatePointWithPointSources(self, request, context):
    """Creates a point with the given point sources.
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def UpdatePoint(self, request, context):
    """Updates an existing point.
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

  def GetTag(self, request, context):
    """Gets a tag.
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

  def ListTags(self, request, context):
    """Lists tags.
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

  def GetPointSource(self, request, context):
    """Gets a point source.
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def SetPointSource(self, request, context):
    """Sets a point source.
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def DeletePointSource(self, request, context):
    """Deletes a point source.
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def SetBrandRegionScore(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def SetTagRegionScore(self, request, context):
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
      'SetBrand': grpc.unary_unary_rpc_method_handler(
          servicer.SetBrand,
          request_deserializer=topos_dot_points_dot_v1_dot_points__pb2.SetBrandRequest.FromString,
          response_serializer=topos_dot_points_dot_v1_dot_points__pb2.Brand.SerializeToString,
      ),
      'ListBrands': grpc.unary_unary_rpc_method_handler(
          servicer.ListBrands,
          request_deserializer=topos_dot_points_dot_v1_dot_points__pb2.ListBrandsRequest.FromString,
          response_serializer=topos_dot_points_dot_v1_dot_points__pb2.ListBrandsResponse.SerializeToString,
      ),
      'DeleteBrand': grpc.unary_unary_rpc_method_handler(
          servicer.DeleteBrand,
          request_deserializer=topos_dot_points_dot_v1_dot_points__pb2.DeleteBrandRequest.FromString,
          response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
      ),
      'SearchPoints': grpc.unary_unary_rpc_method_handler(
          servicer.SearchPoints,
          request_deserializer=topos_dot_points_dot_v1_dot_points__pb2.SearchPointsRequest.FromString,
          response_serializer=topos_dot_points_dot_v1_dot_points__pb2.SearchPointsResponse.SerializeToString,
      ),
      'GetPoint': grpc.unary_unary_rpc_method_handler(
          servicer.GetPoint,
          request_deserializer=topos_dot_points_dot_v1_dot_points__pb2.GetPointRequest.FromString,
          response_serializer=topos_dot_points_dot_v1_dot_points__pb2.Point.SerializeToString,
      ),
      'CreatePointWithPointSources': grpc.unary_unary_rpc_method_handler(
          servicer.CreatePointWithPointSources,
          request_deserializer=topos_dot_points_dot_v1_dot_points__pb2.CreatePointWithPointSourcesRequest.FromString,
          response_serializer=topos_dot_points_dot_v1_dot_points__pb2.CreatePointWithPointSourcesResponse.SerializeToString,
      ),
      'UpdatePoint': grpc.unary_unary_rpc_method_handler(
          servicer.UpdatePoint,
          request_deserializer=topos_dot_points_dot_v1_dot_points__pb2.UpdatePointRequest.FromString,
          response_serializer=topos_dot_points_dot_v1_dot_points__pb2.Point.SerializeToString,
      ),
      'DeletePoint': grpc.unary_unary_rpc_method_handler(
          servicer.DeletePoint,
          request_deserializer=topos_dot_points_dot_v1_dot_points__pb2.DeletePointRequest.FromString,
          response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
      ),
      'GetTag': grpc.unary_unary_rpc_method_handler(
          servicer.GetTag,
          request_deserializer=topos_dot_points_dot_v1_dot_points__pb2.GetTagRequest.FromString,
          response_serializer=topos_dot_points_dot_v1_dot_points__pb2.Tag.SerializeToString,
      ),
      'SetTag': grpc.unary_unary_rpc_method_handler(
          servicer.SetTag,
          request_deserializer=topos_dot_points_dot_v1_dot_points__pb2.SetTagRequest.FromString,
          response_serializer=topos_dot_points_dot_v1_dot_points__pb2.Tag.SerializeToString,
      ),
      'ListTags': grpc.unary_unary_rpc_method_handler(
          servicer.ListTags,
          request_deserializer=topos_dot_points_dot_v1_dot_points__pb2.ListTagsRequest.FromString,
          response_serializer=topos_dot_points_dot_v1_dot_points__pb2.ListTagsResponse.SerializeToString,
      ),
      'DeleteTag': grpc.unary_unary_rpc_method_handler(
          servicer.DeleteTag,
          request_deserializer=topos_dot_points_dot_v1_dot_points__pb2.DeleteTagRequest.FromString,
          response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
      ),
      'GetPointSource': grpc.unary_unary_rpc_method_handler(
          servicer.GetPointSource,
          request_deserializer=topos_dot_points_dot_v1_dot_points__pb2.GetPointSourceRequest.FromString,
          response_serializer=topos_dot_points_dot_v1_dot_points__pb2.PointSource.SerializeToString,
      ),
      'SetPointSource': grpc.unary_unary_rpc_method_handler(
          servicer.SetPointSource,
          request_deserializer=topos_dot_points_dot_v1_dot_points__pb2.SetPointSourceRequest.FromString,
          response_serializer=topos_dot_points_dot_v1_dot_points__pb2.PointSource.SerializeToString,
      ),
      'DeletePointSource': grpc.unary_unary_rpc_method_handler(
          servicer.DeletePointSource,
          request_deserializer=topos_dot_points_dot_v1_dot_points__pb2.DeletePointSourceRequest.FromString,
          response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
      ),
      'SetBrandRegionScore': grpc.unary_unary_rpc_method_handler(
          servicer.SetBrandRegionScore,
          request_deserializer=topos_dot_points_dot_v1_dot_points__pb2.SetBrandRegionScoreRequest.FromString,
          response_serializer=topos_dot_points_dot_v1_dot_points__pb2.BrandRegionScore.SerializeToString,
      ),
      'SetTagRegionScore': grpc.unary_unary_rpc_method_handler(
          servicer.SetTagRegionScore,
          request_deserializer=topos_dot_points_dot_v1_dot_points__pb2.SetTagRegionScoreRequest.FromString,
          response_serializer=topos_dot_points_dot_v1_dot_points__pb2.TagRegionScore.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'topos.points.v1.Points', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))
