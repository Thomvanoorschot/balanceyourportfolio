// Original file: proto/main.proto

import type * as grpc from '@grpc/grpc-js';
import type { MethodDefinition } from '@grpc/proto-loader';
import type { Empty as _proto_Empty, Empty__Output as _proto_Empty__Output } from '../proto/Empty';
import type {
	FilterPortfolioFundHoldingsRequest as _proto_FilterPortfolioFundHoldingsRequest,
	FilterPortfolioFundHoldingsRequest__Output as _proto_FilterPortfolioFundHoldingsRequest__Output
} from '../proto/FilterPortfolioFundHoldingsRequest';
import type {
	FilterPortfolioFundHoldingsResponse as _proto_FilterPortfolioFundHoldingsResponse,
	FilterPortfolioFundHoldingsResponse__Output as _proto_FilterPortfolioFundHoldingsResponse__Output
} from '../proto/FilterPortfolioFundHoldingsResponse';
import type {
	PortfolioDetailsRequest as _proto_PortfolioDetailsRequest,
	PortfolioDetailsRequest__Output as _proto_PortfolioDetailsRequest__Output
} from '../proto/PortfolioDetailsRequest';
import type {
	PortfolioDetailsResponse as _proto_PortfolioDetailsResponse,
	PortfolioDetailsResponse__Output as _proto_PortfolioDetailsResponse__Output
} from '../proto/PortfolioDetailsResponse';
import type {
	PortfoliosRequest as _proto_PortfoliosRequest,
	PortfoliosRequest__Output as _proto_PortfoliosRequest__Output
} from '../proto/PortfoliosRequest';
import type {
	PortfoliosResponse as _proto_PortfoliosResponse,
	PortfoliosResponse__Output as _proto_PortfoliosResponse__Output
} from '../proto/PortfoliosResponse';
import type {
	UpdatePortfolioFundAmountRequest as _proto_UpdatePortfolioFundAmountRequest,
	UpdatePortfolioFundAmountRequest__Output as _proto_UpdatePortfolioFundAmountRequest__Output
} from '../proto/UpdatePortfolioFundAmountRequest';
import type {
	UpsertPortfolioRequest as _proto_UpsertPortfolioRequest,
	UpsertPortfolioRequest__Output as _proto_UpsertPortfolioRequest__Output
} from '../proto/UpsertPortfolioRequest';
import type {
	UpsertPortfolioResponse as _proto_UpsertPortfolioResponse,
	UpsertPortfolioResponse__Output as _proto_UpsertPortfolioResponse__Output
} from '../proto/UpsertPortfolioResponse';

export interface PortfolioServiceClient extends grpc.Client {
	FilterPortfolioHoldings(
		argument: _proto_FilterPortfolioFundHoldingsRequest,
		metadata: grpc.Metadata,
		options: grpc.CallOptions,
		callback: grpc.requestCallback<_proto_FilterPortfolioFundHoldingsResponse__Output>
	): grpc.ClientUnaryCall;
	FilterPortfolioHoldings(
		argument: _proto_FilterPortfolioFundHoldingsRequest,
		metadata: grpc.Metadata,
		callback: grpc.requestCallback<_proto_FilterPortfolioFundHoldingsResponse__Output>
	): grpc.ClientUnaryCall;
	FilterPortfolioHoldings(
		argument: _proto_FilterPortfolioFundHoldingsRequest,
		options: grpc.CallOptions,
		callback: grpc.requestCallback<_proto_FilterPortfolioFundHoldingsResponse__Output>
	): grpc.ClientUnaryCall;
	FilterPortfolioHoldings(
		argument: _proto_FilterPortfolioFundHoldingsRequest,
		callback: grpc.requestCallback<_proto_FilterPortfolioFundHoldingsResponse__Output>
	): grpc.ClientUnaryCall;
	filterPortfolioHoldings(
		argument: _proto_FilterPortfolioFundHoldingsRequest,
		metadata: grpc.Metadata,
		options: grpc.CallOptions,
		callback: grpc.requestCallback<_proto_FilterPortfolioFundHoldingsResponse__Output>
	): grpc.ClientUnaryCall;
	filterPortfolioHoldings(
		argument: _proto_FilterPortfolioFundHoldingsRequest,
		metadata: grpc.Metadata,
		callback: grpc.requestCallback<_proto_FilterPortfolioFundHoldingsResponse__Output>
	): grpc.ClientUnaryCall;
	filterPortfolioHoldings(
		argument: _proto_FilterPortfolioFundHoldingsRequest,
		options: grpc.CallOptions,
		callback: grpc.requestCallback<_proto_FilterPortfolioFundHoldingsResponse__Output>
	): grpc.ClientUnaryCall;
	filterPortfolioHoldings(
		argument: _proto_FilterPortfolioFundHoldingsRequest,
		callback: grpc.requestCallback<_proto_FilterPortfolioFundHoldingsResponse__Output>
	): grpc.ClientUnaryCall;

	GetPortfolioDetails(
		argument: _proto_PortfolioDetailsRequest,
		metadata: grpc.Metadata,
		options: grpc.CallOptions,
		callback: grpc.requestCallback<_proto_PortfolioDetailsResponse__Output>
	): grpc.ClientUnaryCall;
	GetPortfolioDetails(
		argument: _proto_PortfolioDetailsRequest,
		metadata: grpc.Metadata,
		callback: grpc.requestCallback<_proto_PortfolioDetailsResponse__Output>
	): grpc.ClientUnaryCall;
	GetPortfolioDetails(
		argument: _proto_PortfolioDetailsRequest,
		options: grpc.CallOptions,
		callback: grpc.requestCallback<_proto_PortfolioDetailsResponse__Output>
	): grpc.ClientUnaryCall;
	GetPortfolioDetails(
		argument: _proto_PortfolioDetailsRequest,
		callback: grpc.requestCallback<_proto_PortfolioDetailsResponse__Output>
	): grpc.ClientUnaryCall;
	getPortfolioDetails(
		argument: _proto_PortfolioDetailsRequest,
		metadata: grpc.Metadata,
		options: grpc.CallOptions,
		callback: grpc.requestCallback<_proto_PortfolioDetailsResponse__Output>
	): grpc.ClientUnaryCall;
	getPortfolioDetails(
		argument: _proto_PortfolioDetailsRequest,
		metadata: grpc.Metadata,
		callback: grpc.requestCallback<_proto_PortfolioDetailsResponse__Output>
	): grpc.ClientUnaryCall;
	getPortfolioDetails(
		argument: _proto_PortfolioDetailsRequest,
		options: grpc.CallOptions,
		callback: grpc.requestCallback<_proto_PortfolioDetailsResponse__Output>
	): grpc.ClientUnaryCall;
	getPortfolioDetails(
		argument: _proto_PortfolioDetailsRequest,
		callback: grpc.requestCallback<_proto_PortfolioDetailsResponse__Output>
	): grpc.ClientUnaryCall;

	GetPortfolios(
		argument: _proto_PortfoliosRequest,
		metadata: grpc.Metadata,
		options: grpc.CallOptions,
		callback: grpc.requestCallback<_proto_PortfoliosResponse__Output>
	): grpc.ClientUnaryCall;
	GetPortfolios(
		argument: _proto_PortfoliosRequest,
		metadata: grpc.Metadata,
		callback: grpc.requestCallback<_proto_PortfoliosResponse__Output>
	): grpc.ClientUnaryCall;
	GetPortfolios(
		argument: _proto_PortfoliosRequest,
		options: grpc.CallOptions,
		callback: grpc.requestCallback<_proto_PortfoliosResponse__Output>
	): grpc.ClientUnaryCall;
	GetPortfolios(
		argument: _proto_PortfoliosRequest,
		callback: grpc.requestCallback<_proto_PortfoliosResponse__Output>
	): grpc.ClientUnaryCall;
	getPortfolios(
		argument: _proto_PortfoliosRequest,
		metadata: grpc.Metadata,
		options: grpc.CallOptions,
		callback: grpc.requestCallback<_proto_PortfoliosResponse__Output>
	): grpc.ClientUnaryCall;
	getPortfolios(
		argument: _proto_PortfoliosRequest,
		metadata: grpc.Metadata,
		callback: grpc.requestCallback<_proto_PortfoliosResponse__Output>
	): grpc.ClientUnaryCall;
	getPortfolios(
		argument: _proto_PortfoliosRequest,
		options: grpc.CallOptions,
		callback: grpc.requestCallback<_proto_PortfoliosResponse__Output>
	): grpc.ClientUnaryCall;
	getPortfolios(
		argument: _proto_PortfoliosRequest,
		callback: grpc.requestCallback<_proto_PortfoliosResponse__Output>
	): grpc.ClientUnaryCall;

	UpdatePortfolioFundAmount(
		argument: _proto_UpdatePortfolioFundAmountRequest,
		metadata: grpc.Metadata,
		options: grpc.CallOptions,
		callback: grpc.requestCallback<_proto_Empty__Output>
	): grpc.ClientUnaryCall;
	UpdatePortfolioFundAmount(
		argument: _proto_UpdatePortfolioFundAmountRequest,
		metadata: grpc.Metadata,
		callback: grpc.requestCallback<_proto_Empty__Output>
	): grpc.ClientUnaryCall;
	UpdatePortfolioFundAmount(
		argument: _proto_UpdatePortfolioFundAmountRequest,
		options: grpc.CallOptions,
		callback: grpc.requestCallback<_proto_Empty__Output>
	): grpc.ClientUnaryCall;
	UpdatePortfolioFundAmount(
		argument: _proto_UpdatePortfolioFundAmountRequest,
		callback: grpc.requestCallback<_proto_Empty__Output>
	): grpc.ClientUnaryCall;
	updatePortfolioFundAmount(
		argument: _proto_UpdatePortfolioFundAmountRequest,
		metadata: grpc.Metadata,
		options: grpc.CallOptions,
		callback: grpc.requestCallback<_proto_Empty__Output>
	): grpc.ClientUnaryCall;
	updatePortfolioFundAmount(
		argument: _proto_UpdatePortfolioFundAmountRequest,
		metadata: grpc.Metadata,
		callback: grpc.requestCallback<_proto_Empty__Output>
	): grpc.ClientUnaryCall;
	updatePortfolioFundAmount(
		argument: _proto_UpdatePortfolioFundAmountRequest,
		options: grpc.CallOptions,
		callback: grpc.requestCallback<_proto_Empty__Output>
	): grpc.ClientUnaryCall;
	updatePortfolioFundAmount(
		argument: _proto_UpdatePortfolioFundAmountRequest,
		callback: grpc.requestCallback<_proto_Empty__Output>
	): grpc.ClientUnaryCall;

	UpsertPortfolio(
		argument: _proto_UpsertPortfolioRequest,
		metadata: grpc.Metadata,
		options: grpc.CallOptions,
		callback: grpc.requestCallback<_proto_UpsertPortfolioResponse__Output>
	): grpc.ClientUnaryCall;
	UpsertPortfolio(
		argument: _proto_UpsertPortfolioRequest,
		metadata: grpc.Metadata,
		callback: grpc.requestCallback<_proto_UpsertPortfolioResponse__Output>
	): grpc.ClientUnaryCall;
	UpsertPortfolio(
		argument: _proto_UpsertPortfolioRequest,
		options: grpc.CallOptions,
		callback: grpc.requestCallback<_proto_UpsertPortfolioResponse__Output>
	): grpc.ClientUnaryCall;
	UpsertPortfolio(
		argument: _proto_UpsertPortfolioRequest,
		callback: grpc.requestCallback<_proto_UpsertPortfolioResponse__Output>
	): grpc.ClientUnaryCall;
	upsertPortfolio(
		argument: _proto_UpsertPortfolioRequest,
		metadata: grpc.Metadata,
		options: grpc.CallOptions,
		callback: grpc.requestCallback<_proto_UpsertPortfolioResponse__Output>
	): grpc.ClientUnaryCall;
	upsertPortfolio(
		argument: _proto_UpsertPortfolioRequest,
		metadata: grpc.Metadata,
		callback: grpc.requestCallback<_proto_UpsertPortfolioResponse__Output>
	): grpc.ClientUnaryCall;
	upsertPortfolio(
		argument: _proto_UpsertPortfolioRequest,
		options: grpc.CallOptions,
		callback: grpc.requestCallback<_proto_UpsertPortfolioResponse__Output>
	): grpc.ClientUnaryCall;
	upsertPortfolio(
		argument: _proto_UpsertPortfolioRequest,
		callback: grpc.requestCallback<_proto_UpsertPortfolioResponse__Output>
	): grpc.ClientUnaryCall;
}

export interface PortfolioServiceHandlers extends grpc.UntypedServiceImplementation {
	FilterPortfolioHoldings: grpc.handleUnaryCall<
		_proto_FilterPortfolioFundHoldingsRequest__Output,
		_proto_FilterPortfolioFundHoldingsResponse
	>;

	GetPortfolioDetails: grpc.handleUnaryCall<
		_proto_PortfolioDetailsRequest__Output,
		_proto_PortfolioDetailsResponse
	>;

	GetPortfolios: grpc.handleUnaryCall<_proto_PortfoliosRequest__Output, _proto_PortfoliosResponse>;

	UpdatePortfolioFundAmount: grpc.handleUnaryCall<
		_proto_UpdatePortfolioFundAmountRequest__Output,
		_proto_Empty
	>;

	UpsertPortfolio: grpc.handleUnaryCall<
		_proto_UpsertPortfolioRequest__Output,
		_proto_UpsertPortfolioResponse
	>;
}

export interface PortfolioServiceDefinition extends grpc.ServiceDefinition {
	FilterPortfolioHoldings: MethodDefinition<
		_proto_FilterPortfolioFundHoldingsRequest,
		_proto_FilterPortfolioFundHoldingsResponse,
		_proto_FilterPortfolioFundHoldingsRequest__Output,
		_proto_FilterPortfolioFundHoldingsResponse__Output
	>;
	GetPortfolioDetails: MethodDefinition<
		_proto_PortfolioDetailsRequest,
		_proto_PortfolioDetailsResponse,
		_proto_PortfolioDetailsRequest__Output,
		_proto_PortfolioDetailsResponse__Output
	>;
	GetPortfolios: MethodDefinition<
		_proto_PortfoliosRequest,
		_proto_PortfoliosResponse,
		_proto_PortfoliosRequest__Output,
		_proto_PortfoliosResponse__Output
	>;
	UpdatePortfolioFundAmount: MethodDefinition<
		_proto_UpdatePortfolioFundAmountRequest,
		_proto_Empty,
		_proto_UpdatePortfolioFundAmountRequest__Output,
		_proto_Empty__Output
	>;
	UpsertPortfolio: MethodDefinition<
		_proto_UpsertPortfolioRequest,
		_proto_UpsertPortfolioResponse,
		_proto_UpsertPortfolioRequest__Output,
		_proto_UpsertPortfolioResponse__Output
	>;
}
