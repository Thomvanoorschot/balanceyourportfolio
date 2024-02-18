import protoLoader from '@grpc/proto-loader';
import { ChannelCredentials, credentials, loadPackageDefinition } from '@grpc/grpc-js';
import type { ProtoGrpcType } from '$lib/proto/main';
import { env } from '$env/dynamic/private'


export const packageDefinition = protoLoader.loadSync(env.PROTO_FILES_LOCATION, {
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

export const fundClient = new proto.proto.FundService(env.GRPC_API_URL, cr);
export const portfolioClient = new proto.proto.PortfolioService(env.GRPC_API_URL, cr);
