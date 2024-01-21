// Original file: proto/main.proto

import type * as grpc from '@grpc/grpc-js'
import type { MethodDefinition } from '@grpc/proto-loader'
import type { FilterFundHoldingsRequest as _proto_FilterFundHoldingsRequest, FilterFundHoldingsRequest__Output as _proto_FilterFundHoldingsRequest__Output } from '../proto/FilterFundHoldingsRequest';
import type { FilterFundHoldingsResponse as _proto_FilterFundHoldingsResponse, FilterFundHoldingsResponse__Output as _proto_FilterFundHoldingsResponse__Output } from '../proto/FilterFundHoldingsResponse';
import type { FilterFundsRequest as _proto_FilterFundsRequest, FilterFundsRequest__Output as _proto_FilterFundsRequest__Output } from '../proto/FilterFundsRequest';
import type { FilterFundsResponse as _proto_FilterFundsResponse, FilterFundsResponse__Output as _proto_FilterFundsResponse__Output } from '../proto/FilterFundsResponse';
import type { FundDetailsRequest as _proto_FundDetailsRequest, FundDetailsRequest__Output as _proto_FundDetailsRequest__Output } from '../proto/FundDetailsRequest';
import type { FundDetailsResponse as _proto_FundDetailsResponse, FundDetailsResponse__Output as _proto_FundDetailsResponse__Output } from '../proto/FundDetailsResponse';
import type { SearchFundsRequest as _proto_SearchFundsRequest, SearchFundsRequest__Output as _proto_SearchFundsRequest__Output } from '../proto/SearchFundsRequest';
import type { SearchFundsResponse as _proto_SearchFundsResponse, SearchFundsResponse__Output as _proto_SearchFundsResponse__Output } from '../proto/SearchFundsResponse';

export interface FundServiceClient extends grpc.Client {
  FilterFunds(argument: _proto_FilterFundsRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_FilterFundsResponse__Output>): grpc.ClientUnaryCall;
  FilterFunds(argument: _proto_FilterFundsRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_proto_FilterFundsResponse__Output>): grpc.ClientUnaryCall;
  FilterFunds(argument: _proto_FilterFundsRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_FilterFundsResponse__Output>): grpc.ClientUnaryCall;
  FilterFunds(argument: _proto_FilterFundsRequest, callback: grpc.requestCallback<_proto_FilterFundsResponse__Output>): grpc.ClientUnaryCall;
  filterFunds(argument: _proto_FilterFundsRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_FilterFundsResponse__Output>): grpc.ClientUnaryCall;
  filterFunds(argument: _proto_FilterFundsRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_proto_FilterFundsResponse__Output>): grpc.ClientUnaryCall;
  filterFunds(argument: _proto_FilterFundsRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_FilterFundsResponse__Output>): grpc.ClientUnaryCall;
  filterFunds(argument: _proto_FilterFundsRequest, callback: grpc.requestCallback<_proto_FilterFundsResponse__Output>): grpc.ClientUnaryCall;
  
  FilterHoldings(argument: _proto_FilterFundHoldingsRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_FilterFundHoldingsResponse__Output>): grpc.ClientUnaryCall;
  FilterHoldings(argument: _proto_FilterFundHoldingsRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_proto_FilterFundHoldingsResponse__Output>): grpc.ClientUnaryCall;
  FilterHoldings(argument: _proto_FilterFundHoldingsRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_FilterFundHoldingsResponse__Output>): grpc.ClientUnaryCall;
  FilterHoldings(argument: _proto_FilterFundHoldingsRequest, callback: grpc.requestCallback<_proto_FilterFundHoldingsResponse__Output>): grpc.ClientUnaryCall;
  filterHoldings(argument: _proto_FilterFundHoldingsRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_FilterFundHoldingsResponse__Output>): grpc.ClientUnaryCall;
  filterHoldings(argument: _proto_FilterFundHoldingsRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_proto_FilterFundHoldingsResponse__Output>): grpc.ClientUnaryCall;
  filterHoldings(argument: _proto_FilterFundHoldingsRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_FilterFundHoldingsResponse__Output>): grpc.ClientUnaryCall;
  filterHoldings(argument: _proto_FilterFundHoldingsRequest, callback: grpc.requestCallback<_proto_FilterFundHoldingsResponse__Output>): grpc.ClientUnaryCall;
  
  GetDetails(argument: _proto_FundDetailsRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_FundDetailsResponse__Output>): grpc.ClientUnaryCall;
  GetDetails(argument: _proto_FundDetailsRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_proto_FundDetailsResponse__Output>): grpc.ClientUnaryCall;
  GetDetails(argument: _proto_FundDetailsRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_FundDetailsResponse__Output>): grpc.ClientUnaryCall;
  GetDetails(argument: _proto_FundDetailsRequest, callback: grpc.requestCallback<_proto_FundDetailsResponse__Output>): grpc.ClientUnaryCall;
  getDetails(argument: _proto_FundDetailsRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_FundDetailsResponse__Output>): grpc.ClientUnaryCall;
  getDetails(argument: _proto_FundDetailsRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_proto_FundDetailsResponse__Output>): grpc.ClientUnaryCall;
  getDetails(argument: _proto_FundDetailsRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_FundDetailsResponse__Output>): grpc.ClientUnaryCall;
  getDetails(argument: _proto_FundDetailsRequest, callback: grpc.requestCallback<_proto_FundDetailsResponse__Output>): grpc.ClientUnaryCall;
  
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
  FilterFunds: grpc.handleUnaryCall<_proto_FilterFundsRequest__Output, _proto_FilterFundsResponse>;
  
  FilterHoldings: grpc.handleUnaryCall<_proto_FilterFundHoldingsRequest__Output, _proto_FilterFundHoldingsResponse>;
  
  GetDetails: grpc.handleUnaryCall<_proto_FundDetailsRequest__Output, _proto_FundDetailsResponse>;
  
  SearchFunds: grpc.handleUnaryCall<_proto_SearchFundsRequest__Output, _proto_SearchFundsResponse>;
  
}

export interface FundServiceDefinition extends grpc.ServiceDefinition {
  FilterFunds: MethodDefinition<_proto_FilterFundsRequest, _proto_FilterFundsResponse, _proto_FilterFundsRequest__Output, _proto_FilterFundsResponse__Output>
  FilterHoldings: MethodDefinition<_proto_FilterFundHoldingsRequest, _proto_FilterFundHoldingsResponse, _proto_FilterFundHoldingsRequest__Output, _proto_FilterFundHoldingsResponse__Output>
  GetDetails: MethodDefinition<_proto_FundDetailsRequest, _proto_FundDetailsResponse, _proto_FundDetailsRequest__Output, _proto_FundDetailsResponse__Output>
  SearchFunds: MethodDefinition<_proto_SearchFundsRequest, _proto_SearchFundsResponse, _proto_SearchFundsRequest__Output, _proto_SearchFundsResponse__Output>
}
