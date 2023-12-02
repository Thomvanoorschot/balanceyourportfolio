// Original file: proto/main.proto

import type * as grpc from '@grpc/grpc-js'
import type { MethodDefinition } from '@grpc/proto-loader'
import type { FilterHoldingsRequest as _proto_FilterHoldingsRequest, FilterHoldingsRequest__Output as _proto_FilterHoldingsRequest__Output } from '../proto/FilterHoldingsRequest';
import type { FundDetailsResponse as _proto_FundDetailsResponse, FundDetailsResponse__Output as _proto_FundDetailsResponse__Output } from '../proto/FundDetailsResponse';
import type { GetFundDetailsRequest as _proto_GetFundDetailsRequest, GetFundDetailsRequest__Output as _proto_GetFundDetailsRequest__Output } from '../proto/GetFundDetailsRequest';
import type { HoldingsListResponse as _proto_HoldingsListResponse, HoldingsListResponse__Output as _proto_HoldingsListResponse__Output } from '../proto/HoldingsListResponse';
import type { SearchFundsRequest as _proto_SearchFundsRequest, SearchFundsRequest__Output as _proto_SearchFundsRequest__Output } from '../proto/SearchFundsRequest';
import type { SearchFundsResponse as _proto_SearchFundsResponse, SearchFundsResponse__Output as _proto_SearchFundsResponse__Output } from '../proto/SearchFundsResponse';

export interface FundServiceClient extends grpc.Client {
  FilterHoldings(argument: _proto_FilterHoldingsRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_HoldingsListResponse__Output>): grpc.ClientUnaryCall;
  FilterHoldings(argument: _proto_FilterHoldingsRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_proto_HoldingsListResponse__Output>): grpc.ClientUnaryCall;
  FilterHoldings(argument: _proto_FilterHoldingsRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_HoldingsListResponse__Output>): grpc.ClientUnaryCall;
  FilterHoldings(argument: _proto_FilterHoldingsRequest, callback: grpc.requestCallback<_proto_HoldingsListResponse__Output>): grpc.ClientUnaryCall;
  filterHoldings(argument: _proto_FilterHoldingsRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_HoldingsListResponse__Output>): grpc.ClientUnaryCall;
  filterHoldings(argument: _proto_FilterHoldingsRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_proto_HoldingsListResponse__Output>): grpc.ClientUnaryCall;
  filterHoldings(argument: _proto_FilterHoldingsRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_HoldingsListResponse__Output>): grpc.ClientUnaryCall;
  filterHoldings(argument: _proto_FilterHoldingsRequest, callback: grpc.requestCallback<_proto_HoldingsListResponse__Output>): grpc.ClientUnaryCall;
  
  GetDetails(argument: _proto_GetFundDetailsRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_FundDetailsResponse__Output>): grpc.ClientUnaryCall;
  GetDetails(argument: _proto_GetFundDetailsRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_proto_FundDetailsResponse__Output>): grpc.ClientUnaryCall;
  GetDetails(argument: _proto_GetFundDetailsRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_FundDetailsResponse__Output>): grpc.ClientUnaryCall;
  GetDetails(argument: _proto_GetFundDetailsRequest, callback: grpc.requestCallback<_proto_FundDetailsResponse__Output>): grpc.ClientUnaryCall;
  getDetails(argument: _proto_GetFundDetailsRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_FundDetailsResponse__Output>): grpc.ClientUnaryCall;
  getDetails(argument: _proto_GetFundDetailsRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_proto_FundDetailsResponse__Output>): grpc.ClientUnaryCall;
  getDetails(argument: _proto_GetFundDetailsRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_FundDetailsResponse__Output>): grpc.ClientUnaryCall;
  getDetails(argument: _proto_GetFundDetailsRequest, callback: grpc.requestCallback<_proto_FundDetailsResponse__Output>): grpc.ClientUnaryCall;
  
  SearchFunds(argument: _proto_SearchFundsRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_SearchFundsResponse__Output>): grpc.ClientUnaryCall;
  SearchFunds(argument: _proto_SearchFundsRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_proto_SearchFundsResponse__Output>): grpc.ClientUnaryCall;
  SearchFunds(argument: _proto_SearchFundsRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_SearchFundsResponse__Output>): grpc.ClientUnaryCall;
  SearchFunds(argument: _proto_SearchFundsRequest, callback: grpc.requestCallback<_proto_SearchFundsResponse__Output>): grpc.ClientUnaryCall;
  searchFunds(argument: _proto_SearchFundsRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_SearchFundsResponse__Output>): grpc.ClientUnaryCall;
  searchFunds(argument: _proto_SearchFundsRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_proto_SearchFundsResponse__Output>): grpc.ClientUnaryCall;
  searchFunds(argument: _proto_SearchFundsRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_SearchFundsResponse__Output>): grpc.ClientUnaryCall;
  searchFunds(argument: _proto_SearchFundsRequest, callback: grpc.requestCallback<_proto_SearchFundsResponse__Output>): grpc.ClientUnaryCall;
  
}

export interface FundServiceHandlers extends grpc.UntypedServiceImplementation {
  FilterHoldings: grpc.handleUnaryCall<_proto_FilterHoldingsRequest__Output, _proto_HoldingsListResponse>;
  
  GetDetails: grpc.handleUnaryCall<_proto_GetFundDetailsRequest__Output, _proto_FundDetailsResponse>;
  
  SearchFunds: grpc.handleUnaryCall<_proto_SearchFundsRequest__Output, _proto_SearchFundsResponse>;
  
}

export interface FundServiceDefinition extends grpc.ServiceDefinition {
  FilterHoldings: MethodDefinition<_proto_FilterHoldingsRequest, _proto_HoldingsListResponse, _proto_FilterHoldingsRequest__Output, _proto_HoldingsListResponse__Output>
  GetDetails: MethodDefinition<_proto_GetFundDetailsRequest, _proto_FundDetailsResponse, _proto_GetFundDetailsRequest__Output, _proto_FundDetailsResponse__Output>
  SearchFunds: MethodDefinition<_proto_SearchFundsRequest, _proto_SearchFundsResponse, _proto_SearchFundsRequest__Output, _proto_SearchFundsResponse__Output>
}
