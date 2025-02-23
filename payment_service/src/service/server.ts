import { loadSync } from "@grpc/proto-loader";
import { GRPC_SERVER_CONFIG_TOKEN, GRPCServerConfig } from "../config/grpc_server";
import { ProtoGrpcType } from "../proto/gen/payment_service";
import { LOGGER_WINSTON_TOKEN, LoggerWinston } from "../utils/logger";
import { loadPackageDefinition, Server, ServerCredentials } from "@grpc/grpc-js";
import { injected, token } from "brandi";
import { PAYMENT_SERVICE_HANDLER_FACTORY_TOKEN, PaymentServiceHandlerFactory } from "./handler";
import { ReflectionService } from '@grpc/reflection';

export class PaymentServiceGRPCServer {

    constructor(
        private readonly logger: LoggerWinston,
        private readonly grpcConfig: GRPCServerConfig,
        private readonly handlerFactory: PaymentServiceHandlerFactory,
    ) {
        this.logger = new LoggerWinston();
    }

    public async loadProtoAndStartServer(protoPath: string) {
        // load proto
        const paymentServiceProtoGrpc = this.loadPaymentServiceProtoGRPC(protoPath);

        const server = new Server({
            "grpc.max_send_message_length": -1,
            "grpc.max_receive_message_length": -1
        })

        server.addService(paymentServiceProtoGrpc.protoGrpcType.payment_service.PaymentService.service, this.handlerFactory.handlers());
        const reflection = new ReflectionService(paymentServiceProtoGrpc.packageDefinition)
        reflection.addToServer(server)
        server.bindAsync(`0.0.0.0:${this.grpcConfig.port}`, ServerCredentials.createInsecure(), (error, port) => {
            if (error) {
                this.logger.error("failed to start server gRPC.")
            }

            this.logger.info(`start server gRPC on port: ${port}`)
        })
    }

    public loadPaymentServiceProtoGRPC(protoPath: string): { protoGrpcType: ProtoGrpcType, packageDefinition: any } {
        const packageDefinition = loadSync(protoPath, {
            enums: Number,
            keepCase: false,
            defaults: false,
            oneofs: true,
            longs: Number
        });


        const packageObject = loadPackageDefinition(packageDefinition) as unknown;
        return { protoGrpcType: packageObject as ProtoGrpcType, packageDefinition: packageDefinition }
    }
}

injected(PaymentServiceGRPCServer, LOGGER_WINSTON_TOKEN, GRPC_SERVER_CONFIG_TOKEN, PAYMENT_SERVICE_HANDLER_FACTORY_TOKEN);

export const PAYMENT_SERVICE_GRPC_SERVER_TOKEN = token<PaymentServiceGRPCServer>("PaymentServiceGRPCServer");
