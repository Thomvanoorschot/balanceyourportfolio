import protoLoader from '@grpc/proto-loader';
import { ChannelCredentials, credentials, loadPackageDefinition } from '@grpc/grpc-js';
import type { ProtoGrpcType } from '$lib/proto/main';

const url = 'localhost:8080';

export const packageDefinition = protoLoader.loadSync('./src/lib/proto/main.proto', {
	keepCase: false,
	longs: String,
	defaults: true,
	oneofs: true,
	arrays: true
});
export const proto = loadPackageDefinition(packageDefinition) as unknown as ProtoGrpcType;

const cr: ChannelCredentials = credentials.createInsecure();
// const cr: ChannelCredentials =
//     ENV === "production"
//         ? credentials.createSsl()
//         : credentials.createInsecure();

export const fundClient = new proto.proto.FundService(url, cr);
export const portfolioClient = new proto.proto.PortfolioService(url, cr);
