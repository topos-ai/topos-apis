# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from topos.scores.v1 import scores_pb2 as topos_dot_scores_dot_v1_dot_scores__pb2


class ScoresStub(object):
  # missing associated documentation comment in .proto file
  pass

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.GetScore = channel.unary_unary(
        '/topos.scores.v1.Scores/GetScore',
        request_serializer=topos_dot_scores_dot_v1_dot_scores__pb2.GetScoreRequest.SerializeToString,
        response_deserializer=topos_dot_scores_dot_v1_dot_scores__pb2.Score.FromString,
        )
    self.SetScore = channel.unary_unary(
        '/topos.scores.v1.Scores/SetScore',
        request_serializer=topos_dot_scores_dot_v1_dot_scores__pb2.SetScoreRequest.SerializeToString,
        response_deserializer=topos_dot_scores_dot_v1_dot_scores__pb2.Score.FromString,
        )
    self.BatchSetScore = channel.unary_unary(
        '/topos.scores.v1.Scores/BatchSetScore',
        request_serializer=topos_dot_scores_dot_v1_dot_scores__pb2.BatchSetScoreRequest.SerializeToString,
        response_deserializer=topos_dot_scores_dot_v1_dot_scores__pb2.BatchSetScoreResponse.FromString,
        )
    self.DeleteScore = channel.unary_unary(
        '/topos.scores.v1.Scores/DeleteScore',
        request_serializer=topos_dot_scores_dot_v1_dot_scores__pb2.DeleteScoreRequest.SerializeToString,
        response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
        )
    self.SearchScores = channel.unary_unary(
        '/topos.scores.v1.Scores/SearchScores',
        request_serializer=topos_dot_scores_dot_v1_dot_scores__pb2.SearchScoresRequest.SerializeToString,
        response_deserializer=topos_dot_scores_dot_v1_dot_scores__pb2.SearchScoresResponse.FromString,
        )


class ScoresServicer(object):
  # missing associated documentation comment in .proto file
  pass

  def GetScore(self, request, context):
    """Gets a score.
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def SetScore(self, request, context):
    """Sets a score.
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def BatchSetScore(self, request, context):
    """Sets a batch of scores.
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def DeleteScore(self, request, context):
    """Deletes a score.
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def SearchScores(self, request, context):
    """Search scores.
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_ScoresServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'GetScore': grpc.unary_unary_rpc_method_handler(
          servicer.GetScore,
          request_deserializer=topos_dot_scores_dot_v1_dot_scores__pb2.GetScoreRequest.FromString,
          response_serializer=topos_dot_scores_dot_v1_dot_scores__pb2.Score.SerializeToString,
      ),
      'SetScore': grpc.unary_unary_rpc_method_handler(
          servicer.SetScore,
          request_deserializer=topos_dot_scores_dot_v1_dot_scores__pb2.SetScoreRequest.FromString,
          response_serializer=topos_dot_scores_dot_v1_dot_scores__pb2.Score.SerializeToString,
      ),
      'BatchSetScore': grpc.unary_unary_rpc_method_handler(
          servicer.BatchSetScore,
          request_deserializer=topos_dot_scores_dot_v1_dot_scores__pb2.BatchSetScoreRequest.FromString,
          response_serializer=topos_dot_scores_dot_v1_dot_scores__pb2.BatchSetScoreResponse.SerializeToString,
      ),
      'DeleteScore': grpc.unary_unary_rpc_method_handler(
          servicer.DeleteScore,
          request_deserializer=topos_dot_scores_dot_v1_dot_scores__pb2.DeleteScoreRequest.FromString,
          response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
      ),
      'SearchScores': grpc.unary_unary_rpc_method_handler(
          servicer.SearchScores,
          request_deserializer=topos_dot_scores_dot_v1_dot_scores__pb2.SearchScoresRequest.FromString,
          response_serializer=topos_dot_scores_dot_v1_dot_scores__pb2.SearchScoresResponse.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'topos.scores.v1.Scores', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))
