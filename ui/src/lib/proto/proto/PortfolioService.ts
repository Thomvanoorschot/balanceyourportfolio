// Original file: proto/main.proto

import type * as grpc from '@grpc/grpc-js'
import type { MethodDefinition } from '@grpc/proto-loader'
import type { PortfolioDetailsRequest as _proto_PortfolioDetailsRequest, PortfolioDetailsRequest__Output as _proto_PortfolioDetailsRequest__Output } from '../proto/PortfolioDetailsRequest';
import type { PortfolioDetailsResponse as _proto_PortfolioDetailsResponse, PortfolioDetailsResponse__Output as _proto_PortfolioDetailsResponse__Output } from '../proto/PortfolioDetailsResponse';
import type { PortfoliosRequest as _proto_PortfoliosRequest, PortfoliosRequest__Output as _proto_PortfoliosRequest__Output } from '../proto/PortfoliosRequest';
import type { PortfoliosResponse as _proto_PortfoliosResponse, PortfoliosResponse__Output as _proto_PortfoliosResponse__Output } from '../proto/PortfoliosResponse';
import type { UpsertPortfolioRequest as _proto_UpsertPortfolioRequest, UpsertPortfolioRequest__Output as _proto_UpsertPortfolioRequest__Output } from '../proto/UpsertPortfolioRequest';
import type { UpsertPortfolioResponse as _proto_UpsertPortfolioResponse, UpsertPortfolioResponse__Output as _proto_UpsertPortfolioResponse__Output } from '../proto/UpsertPortfolioResponse';

export interface PortfolioServiceClient extends grpc.Client {
  GetPortfolioDetails(argument: _proto_PortfolioDetailsRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_PortfolioDetailsResponse__Output>): grpc.ClientUnaryCall;
  GetPortfolioDetails(argument: _proto_PortfolioDetailsRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_proto_PortfolioDetailsResponse__Output>): grpc.ClientUnaryCall;
  GetPortfolioDetails(argument: _proto_PortfolioDetailsRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_PortfolioDetailsResponse__Output>): grpc.ClientUnaryCall;
  GetPortfolioDetails(argument: _proto_PortfolioDetailsRequest, callback: grpc.requestCallback<_proto_PortfolioDetailsResponse__Output>): grpc.ClientUnaryCall;
  getPortfolioDetails(argument: _proto_PortfolioDetailsRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_PortfolioDetailsResponse__Output>): grpc.ClientUnaryCall;
  getPortfolioDetails(argument: _proto_PortfolioDetailsRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_proto_PortfolioDetailsResponse__Output>): grpc.ClientUnaryCall;
  getPortfolioDetails(argument: _proto_PortfolioDetailsRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_PortfolioDetailsResponse__Output>): grpc.ClientUnaryCall;
  getPortfolioDetails(argument: _proto_PortfolioDetailsRequest, callback: grpc.requestCallback<_proto_PortfolioDetailsResponse__Output>): grpc.ClientUnaryCall;
  
  GetPortfolios(argument: _proto_PortfoliosRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_PortfoliosResponse__Output>): grpc.ClientUnaryCall;
  GetPortfolios(argument: _proto_PortfoliosRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_proto_PortfoliosResponse__Output>): grpc.ClientUnaryCall;
  GetPortfolios(argument: _proto_PortfoliosRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_PortfoliosResponse__Output>): grpc.ClientUnaryCall;
  GetPortfolios(argument: _proto_PortfoliosRequest, callback: grpc.requestCallback<_proto_PortfoliosResponse__Output>): grpc.ClientUnaryCall;
  getPortfolios(argument: _proto_PortfoliosRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_PortfoliosResponse__Output>): grpc.ClientUnaryCall;
  getPortfolios(argument: _proto_PortfoliosRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_proto_PortfoliosResponse__Output>): grpc.ClientUnaryCall;
  getPortfolios(argument: _proto_PortfoliosRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_PortfoliosResponse__Output>): grpc.ClientUnaryCall;
  getPortfolios(argument: _proto_PortfoliosRequest, callback: grpc.requestCallback<_proto_PortfoliosResponse__Output>): grpc.ClientUnaryCall;
  
  UpsertPortfolio(argument: _proto_UpsertPortfolioRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_UpsertPortfolioResponse__Output>): grpc.ClientUnaryCall;
  UpsertPortfolio(argument: _proto_UpsertPortfolioRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_proto_UpsertPortfolioResponse__Output>): grpc.ClientUnaryCall;
  UpsertPortfolio(argument: _proto_UpsertPortfolioRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_UpsertPortfolioResponse__Output>): grpc.ClientUnaryCall;
  UpsertPortfolio(argument: _proto_UpsertPortfolioRequest, callback: grpc.requestCallback<_proto_UpsertPortfolioResponse__Output>): grpc.ClientUnaryCall;
  upsertPortfolio(argument: _proto_UpsertPortfolioRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_UpsertPortfolioResponse__Output>): grpc.ClientUnaryCall;
  upsertPortfolio(argument: _proto_UpsertPortfolioRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_proto_UpsertPortfolioResponse__Output>): grpc.ClientUnaryCall;
  upsertPortfolio(argument: _proto_UpsertPortfolioRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_UpsertPortfolioResponse__Output>): grpc.ClientUnaryCall;
  upsertPortfolio(argument: _proto_UpsertPortfolioRequest, callback: grpc.requestCallback<_proto_UpsertPortfolioResponse__Output>): grpc.ClientUnaryCall;
  
}

export interface PortfolioServiceHandlers extends grpc.UntypedServiceImplementation {
  GetPortfolioDetails: grpc.handleUnaryCall<_proto_PortfolioDetailsRequest__Output, _proto_PortfolioDetailsResponse>;
  
  GetPortfolios: grpc.handleUnaryCall<_proto_PortfoliosRequest__Output, _proto_PortfoliosResponse>;
  
  UpsertPortfolio: grpc.handleUnaryCall<_proto_UpsertPortfolioRequest__Output, _proto_UpsertPortfolioResponse>;
  
}

export interface PortfolioServiceDefinition extends grpc.ServiceDefinition {
  GetPortfolioDetails: MethodDefinition<_proto_PortfolioDetailsRequest, _proto_PortfolioDetailsResponse, _proto_PortfolioDetailsRequest__Output, _proto_PortfolioDetailsResponse__Output>
  GetPortfolios: MethodDefinition<_proto_PortfoliosRequest, _proto_PortfoliosResponse, _proto_PortfoliosRequest__Output, _proto_PortfoliosResponse__Output>
  UpsertPortfolio: MethodDefinition<_proto_UpsertPortfolioRequest, _proto_UpsertPortfolioResponse, _proto_UpsertPortfolioRequest__Output, _proto_UpsertPortfolioResponse__Output>
}
