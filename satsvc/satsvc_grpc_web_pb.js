/**
 * @fileoverview gRPC-Web generated client stub for 
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');


var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js')

var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js')
const proto = require('./satsvc_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.PredictionClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'binary';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.PredictionPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'binary';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.SatsRequest,
 *   !proto.SatsInfosResponse>}
 */
const methodDescriptor_Prediction_SatsInfos = new grpc.web.MethodDescriptor(
  '/Prediction/SatsInfos',
  grpc.web.MethodType.UNARY,
  proto.SatsRequest,
  proto.SatsInfosResponse,
  /**
   * @param {!proto.SatsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.SatsInfosResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.SatsRequest,
 *   !proto.SatsInfosResponse>}
 */
const methodInfo_Prediction_SatsInfos = new grpc.web.AbstractClientBase.MethodInfo(
  proto.SatsInfosResponse,
  /**
   * @param {!proto.SatsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.SatsInfosResponse.deserializeBinary
);


/**
 * @param {!proto.SatsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.SatsInfosResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.SatsInfosResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.PredictionClient.prototype.satsInfos =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/Prediction/SatsInfos',
      request,
      metadata || {},
      methodDescriptor_Prediction_SatsInfos,
      callback);
};


/**
 * @param {!proto.SatsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.SatsInfosResponse>}
 *     A native promise that resolves to the response
 */
proto.PredictionPromiseClient.prototype.satsInfos =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/Prediction/SatsInfos',
      request,
      metadata || {},
      methodDescriptor_Prediction_SatsInfos);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.SatsRequest,
 *   !proto.SatsLocationsResponse>}
 */
const methodDescriptor_Prediction_SatsLocations = new grpc.web.MethodDescriptor(
  '/Prediction/SatsLocations',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.SatsRequest,
  proto.SatsLocationsResponse,
  /**
   * @param {!proto.SatsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.SatsLocationsResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.SatsRequest,
 *   !proto.SatsLocationsResponse>}
 */
const methodInfo_Prediction_SatsLocations = new grpc.web.AbstractClientBase.MethodInfo(
  proto.SatsLocationsResponse,
  /**
   * @param {!proto.SatsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.SatsLocationsResponse.deserializeBinary
);


/**
 * @param {!proto.SatsRequest} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.SatsLocationsResponse>}
 *     The XHR Node Readable Stream
 */
proto.PredictionClient.prototype.satsLocations =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/Prediction/SatsLocations',
      request,
      metadata || {},
      methodDescriptor_Prediction_SatsLocations);
};


/**
 * @param {!proto.SatsRequest} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.SatsLocationsResponse>}
 *     The XHR Node Readable Stream
 */
proto.PredictionPromiseClient.prototype.satsLocations =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/Prediction/SatsLocations',
      request,
      metadata || {},
      methodDescriptor_Prediction_SatsLocations);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.SatLocationFromObsRequest,
 *   !proto.Observation>}
 */
const methodDescriptor_Prediction_SatLocationFromObs = new grpc.web.MethodDescriptor(
  '/Prediction/SatLocationFromObs',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.SatLocationFromObsRequest,
  proto.Observation,
  /**
   * @param {!proto.SatLocationFromObsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.Observation.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.SatLocationFromObsRequest,
 *   !proto.Observation>}
 */
const methodInfo_Prediction_SatLocationFromObs = new grpc.web.AbstractClientBase.MethodInfo(
  proto.Observation,
  /**
   * @param {!proto.SatLocationFromObsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.Observation.deserializeBinary
);


/**
 * @param {!proto.SatLocationFromObsRequest} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.Observation>}
 *     The XHR Node Readable Stream
 */
proto.PredictionClient.prototype.satLocationFromObs =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/Prediction/SatLocationFromObs',
      request,
      metadata || {},
      methodDescriptor_Prediction_SatLocationFromObs);
};


/**
 * @param {!proto.SatLocationFromObsRequest} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.Observation>}
 *     The XHR Node Readable Stream
 */
proto.PredictionPromiseClient.prototype.satLocationFromObs =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/Prediction/SatLocationFromObs',
      request,
      metadata || {},
      methodDescriptor_Prediction_SatLocationFromObs);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.GenLocationsRequest,
 *   !proto.GenLocationsResponse>}
 */
const methodDescriptor_Prediction_GenLocations = new grpc.web.MethodDescriptor(
  '/Prediction/GenLocations',
  grpc.web.MethodType.UNARY,
  proto.GenLocationsRequest,
  proto.GenLocationsResponse,
  /**
   * @param {!proto.GenLocationsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.GenLocationsResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.GenLocationsRequest,
 *   !proto.GenLocationsResponse>}
 */
const methodInfo_Prediction_GenLocations = new grpc.web.AbstractClientBase.MethodInfo(
  proto.GenLocationsResponse,
  /**
   * @param {!proto.GenLocationsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.GenLocationsResponse.deserializeBinary
);


/**
 * @param {!proto.GenLocationsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.GenLocationsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.GenLocationsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.PredictionClient.prototype.genLocations =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/Prediction/GenLocations',
      request,
      metadata || {},
      methodDescriptor_Prediction_GenLocations,
      callback);
};


/**
 * @param {!proto.GenLocationsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.GenLocationsResponse>}
 *     A native promise that resolves to the response
 */
proto.PredictionPromiseClient.prototype.genLocations =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/Prediction/GenLocations',
      request,
      metadata || {},
      methodDescriptor_Prediction_GenLocations);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.GenPassesRequest,
 *   !proto.Passes>}
 */
const methodDescriptor_Prediction_GenPasses = new grpc.web.MethodDescriptor(
  '/Prediction/GenPasses',
  grpc.web.MethodType.UNARY,
  proto.GenPassesRequest,
  proto.Passes,
  /**
   * @param {!proto.GenPassesRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.Passes.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.GenPassesRequest,
 *   !proto.Passes>}
 */
const methodInfo_Prediction_GenPasses = new grpc.web.AbstractClientBase.MethodInfo(
  proto.Passes,
  /**
   * @param {!proto.GenPassesRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.Passes.deserializeBinary
);


/**
 * @param {!proto.GenPassesRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.Passes)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.Passes>|undefined}
 *     The XHR Node Readable Stream
 */
proto.PredictionClient.prototype.genPasses =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/Prediction/GenPasses',
      request,
      metadata || {},
      methodDescriptor_Prediction_GenPasses,
      callback);
};


/**
 * @param {!proto.GenPassesRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.Passes>}
 *     A native promise that resolves to the response
 */
proto.PredictionPromiseClient.prototype.genPasses =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/Prediction/GenPasses',
      request,
      metadata || {},
      methodDescriptor_Prediction_GenPasses);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.google.protobuf.Empty,
 *   !proto.CategoriesResponse>}
 */
const methodDescriptor_Prediction_Categories = new grpc.web.MethodDescriptor(
  '/Prediction/Categories',
  grpc.web.MethodType.UNARY,
  google_protobuf_empty_pb.Empty,
  proto.CategoriesResponse,
  /**
   * @param {!proto.google.protobuf.Empty} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.CategoriesResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.google.protobuf.Empty,
 *   !proto.CategoriesResponse>}
 */
const methodInfo_Prediction_Categories = new grpc.web.AbstractClientBase.MethodInfo(
  proto.CategoriesResponse,
  /**
   * @param {!proto.google.protobuf.Empty} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.CategoriesResponse.deserializeBinary
);


/**
 * @param {!proto.google.protobuf.Empty} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.CategoriesResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.CategoriesResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.PredictionClient.prototype.categories =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/Prediction/Categories',
      request,
      metadata || {},
      methodDescriptor_Prediction_Categories,
      callback);
};


/**
 * @param {!proto.google.protobuf.Empty} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.CategoriesResponse>}
 *     A native promise that resolves to the response
 */
proto.PredictionPromiseClient.prototype.categories =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/Prediction/Categories',
      request,
      metadata || {},
      methodDescriptor_Prediction_Categories);
};


module.exports = proto;

