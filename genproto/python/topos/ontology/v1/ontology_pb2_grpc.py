# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from topos.ontology.v1 import ontology_pb2 as topos_dot_ontology_dot_v1_dot_ontology__pb2


class OntologyStub(object):
  # missing associated documentation comment in .proto file
  pass

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.GetBrand = channel.unary_unary(
        '/topos.ontology.v1.Ontology/GetBrand',
        request_serializer=topos_dot_ontology_dot_v1_dot_ontology__pb2.GetBrandRequest.SerializeToString,
        response_deserializer=topos_dot_ontology_dot_v1_dot_ontology__pb2.Brand.FromString,
        )
    self.GetBrandFeatures = channel.unary_unary(
        '/topos.ontology.v1.Ontology/GetBrandFeatures',
        request_serializer=topos_dot_ontology_dot_v1_dot_ontology__pb2.GetBrandFeaturesRequest.SerializeToString,
        response_deserializer=topos_dot_ontology_dot_v1_dot_ontology__pb2.GetBrandFeaturesResponse.FromString,
        )
    self.GetBrandContacts = channel.unary_unary(
        '/topos.ontology.v1.Ontology/GetBrandContacts',
        request_serializer=topos_dot_ontology_dot_v1_dot_ontology__pb2.GetBrandContactsRequest.SerializeToString,
        response_deserializer=topos_dot_ontology_dot_v1_dot_ontology__pb2.GetBrandContactsResponse.FromString,
        )
    self.GetBrandsMinDistance = channel.unary_unary(
        '/topos.ontology.v1.Ontology/GetBrandsMinDistance',
        request_serializer=topos_dot_ontology_dot_v1_dot_ontology__pb2.GetBrandsMinDistanceRequest.SerializeToString,
        response_deserializer=topos_dot_ontology_dot_v1_dot_ontology__pb2.GetBrandsMinDistanceResponse.FromString,
        )
    self.GetBrandsBulk = channel.unary_unary(
        '/topos.ontology.v1.Ontology/GetBrandsBulk',
        request_serializer=topos_dot_ontology_dot_v1_dot_ontology__pb2.GetBrandsBulkRequest.SerializeToString,
        response_deserializer=topos_dot_ontology_dot_v1_dot_ontology__pb2.GetBrandsBulkResponse.FromString,
        )
    self.ListBrands = channel.unary_unary(
        '/topos.ontology.v1.Ontology/ListBrands',
        request_serializer=topos_dot_ontology_dot_v1_dot_ontology__pb2.ListBrandsRequest.SerializeToString,
        response_deserializer=topos_dot_ontology_dot_v1_dot_ontology__pb2.ListBrandsResponse.FromString,
        )
    self.SearchBrands = channel.unary_unary(
        '/topos.ontology.v1.Ontology/SearchBrands',
        request_serializer=topos_dot_ontology_dot_v1_dot_ontology__pb2.SearchBrandsRequest.SerializeToString,
        response_deserializer=topos_dot_ontology_dot_v1_dot_ontology__pb2.SearchBrandsResponse.FromString,
        )
    self.SetBrand = channel.unary_unary(
        '/topos.ontology.v1.Ontology/SetBrand',
        request_serializer=topos_dot_ontology_dot_v1_dot_ontology__pb2.SetBrandRequest.SerializeToString,
        response_deserializer=topos_dot_ontology_dot_v1_dot_ontology__pb2.Brand.FromString,
        )
    self.SetBrandsBulk = channel.unary_unary(
        '/topos.ontology.v1.Ontology/SetBrandsBulk',
        request_serializer=topos_dot_ontology_dot_v1_dot_ontology__pb2.SetBrandsBulkRequest.SerializeToString,
        response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
        )
    self.DeleteBrand = channel.unary_unary(
        '/topos.ontology.v1.Ontology/DeleteBrand',
        request_serializer=topos_dot_ontology_dot_v1_dot_ontology__pb2.DeleteBrandRequest.SerializeToString,
        response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
        )
    self.GetTag = channel.unary_unary(
        '/topos.ontology.v1.Ontology/GetTag',
        request_serializer=topos_dot_ontology_dot_v1_dot_ontology__pb2.GetTagRequest.SerializeToString,
        response_deserializer=topos_dot_ontology_dot_v1_dot_ontology__pb2.Tag.FromString,
        )
    self.ListTags = channel.unary_unary(
        '/topos.ontology.v1.Ontology/ListTags',
        request_serializer=topos_dot_ontology_dot_v1_dot_ontology__pb2.ListTagsRequest.SerializeToString,
        response_deserializer=topos_dot_ontology_dot_v1_dot_ontology__pb2.ListTagsResponse.FromString,
        )
    self.ListResolvedTags = channel.unary_unary(
        '/topos.ontology.v1.Ontology/ListResolvedTags',
        request_serializer=topos_dot_ontology_dot_v1_dot_ontology__pb2.ListResolvedTagsRequest.SerializeToString,
        response_deserializer=topos_dot_ontology_dot_v1_dot_ontology__pb2.ListResolvedTagsResponse.FromString,
        )
    self.SetTag = channel.unary_unary(
        '/topos.ontology.v1.Ontology/SetTag',
        request_serializer=topos_dot_ontology_dot_v1_dot_ontology__pb2.SetTagRequest.SerializeToString,
        response_deserializer=topos_dot_ontology_dot_v1_dot_ontology__pb2.Tag.FromString,
        )
    self.DeleteTag = channel.unary_unary(
        '/topos.ontology.v1.Ontology/DeleteTag',
        request_serializer=topos_dot_ontology_dot_v1_dot_ontology__pb2.DeleteTagRequest.SerializeToString,
        response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
        )


class OntologyServicer(object):
  # missing associated documentation comment in .proto file
  pass

  def GetBrand(self, request, context):
    """Gets a brand.
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def GetBrandFeatures(self, request, context):
    """Gets features for a brand.
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def GetBrandContacts(self, request, context):
    """Gets features for a brand.
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def GetBrandsMinDistance(self, request, context):
    """Gets min distanct for a number brand.
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def GetBrandsBulk(self, request, context):
    """Gets a number of brands.
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

  def SearchBrands(self, request, context):
    """Searches brands.
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

  def SetBrandsBulk(self, request, context):
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

  def ListResolvedTags(self, request, context):
    """Lists resolved tags.
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


def add_OntologyServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'GetBrand': grpc.unary_unary_rpc_method_handler(
          servicer.GetBrand,
          request_deserializer=topos_dot_ontology_dot_v1_dot_ontology__pb2.GetBrandRequest.FromString,
          response_serializer=topos_dot_ontology_dot_v1_dot_ontology__pb2.Brand.SerializeToString,
      ),
      'GetBrandFeatures': grpc.unary_unary_rpc_method_handler(
          servicer.GetBrandFeatures,
          request_deserializer=topos_dot_ontology_dot_v1_dot_ontology__pb2.GetBrandFeaturesRequest.FromString,
          response_serializer=topos_dot_ontology_dot_v1_dot_ontology__pb2.GetBrandFeaturesResponse.SerializeToString,
      ),
      'GetBrandContacts': grpc.unary_unary_rpc_method_handler(
          servicer.GetBrandContacts,
          request_deserializer=topos_dot_ontology_dot_v1_dot_ontology__pb2.GetBrandContactsRequest.FromString,
          response_serializer=topos_dot_ontology_dot_v1_dot_ontology__pb2.GetBrandContactsResponse.SerializeToString,
      ),
      'GetBrandsMinDistance': grpc.unary_unary_rpc_method_handler(
          servicer.GetBrandsMinDistance,
          request_deserializer=topos_dot_ontology_dot_v1_dot_ontology__pb2.GetBrandsMinDistanceRequest.FromString,
          response_serializer=topos_dot_ontology_dot_v1_dot_ontology__pb2.GetBrandsMinDistanceResponse.SerializeToString,
      ),
      'GetBrandsBulk': grpc.unary_unary_rpc_method_handler(
          servicer.GetBrandsBulk,
          request_deserializer=topos_dot_ontology_dot_v1_dot_ontology__pb2.GetBrandsBulkRequest.FromString,
          response_serializer=topos_dot_ontology_dot_v1_dot_ontology__pb2.GetBrandsBulkResponse.SerializeToString,
      ),
      'ListBrands': grpc.unary_unary_rpc_method_handler(
          servicer.ListBrands,
          request_deserializer=topos_dot_ontology_dot_v1_dot_ontology__pb2.ListBrandsRequest.FromString,
          response_serializer=topos_dot_ontology_dot_v1_dot_ontology__pb2.ListBrandsResponse.SerializeToString,
      ),
      'SearchBrands': grpc.unary_unary_rpc_method_handler(
          servicer.SearchBrands,
          request_deserializer=topos_dot_ontology_dot_v1_dot_ontology__pb2.SearchBrandsRequest.FromString,
          response_serializer=topos_dot_ontology_dot_v1_dot_ontology__pb2.SearchBrandsResponse.SerializeToString,
      ),
      'SetBrand': grpc.unary_unary_rpc_method_handler(
          servicer.SetBrand,
          request_deserializer=topos_dot_ontology_dot_v1_dot_ontology__pb2.SetBrandRequest.FromString,
          response_serializer=topos_dot_ontology_dot_v1_dot_ontology__pb2.Brand.SerializeToString,
      ),
      'SetBrandsBulk': grpc.unary_unary_rpc_method_handler(
          servicer.SetBrandsBulk,
          request_deserializer=topos_dot_ontology_dot_v1_dot_ontology__pb2.SetBrandsBulkRequest.FromString,
          response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
      ),
      'DeleteBrand': grpc.unary_unary_rpc_method_handler(
          servicer.DeleteBrand,
          request_deserializer=topos_dot_ontology_dot_v1_dot_ontology__pb2.DeleteBrandRequest.FromString,
          response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
      ),
      'GetTag': grpc.unary_unary_rpc_method_handler(
          servicer.GetTag,
          request_deserializer=topos_dot_ontology_dot_v1_dot_ontology__pb2.GetTagRequest.FromString,
          response_serializer=topos_dot_ontology_dot_v1_dot_ontology__pb2.Tag.SerializeToString,
      ),
      'ListTags': grpc.unary_unary_rpc_method_handler(
          servicer.ListTags,
          request_deserializer=topos_dot_ontology_dot_v1_dot_ontology__pb2.ListTagsRequest.FromString,
          response_serializer=topos_dot_ontology_dot_v1_dot_ontology__pb2.ListTagsResponse.SerializeToString,
      ),
      'ListResolvedTags': grpc.unary_unary_rpc_method_handler(
          servicer.ListResolvedTags,
          request_deserializer=topos_dot_ontology_dot_v1_dot_ontology__pb2.ListResolvedTagsRequest.FromString,
          response_serializer=topos_dot_ontology_dot_v1_dot_ontology__pb2.ListResolvedTagsResponse.SerializeToString,
      ),
      'SetTag': grpc.unary_unary_rpc_method_handler(
          servicer.SetTag,
          request_deserializer=topos_dot_ontology_dot_v1_dot_ontology__pb2.SetTagRequest.FromString,
          response_serializer=topos_dot_ontology_dot_v1_dot_ontology__pb2.Tag.SerializeToString,
      ),
      'DeleteTag': grpc.unary_unary_rpc_method_handler(
          servicer.DeleteTag,
          request_deserializer=topos_dot_ontology_dot_v1_dot_ontology__pb2.DeleteTagRequest.FromString,
          response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'topos.ontology.v1.Ontology', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))
