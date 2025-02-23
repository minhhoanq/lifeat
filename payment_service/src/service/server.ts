import { GRPC_SERVER_CONFIG_TOKEN, GRPCServerConfig } from "../config/grpc_server";
import { LOGGER_WINSTON_TOKEN, LoggerWinston } from "../utils/logger";
import { Server, ServerCredentials } from "@grpc/grpc-js";
import { injected, token } from "brandi";

export class PaymentServiceGRPCServer {
    constructor(
        private readonly logger: LoggerWinston,
        private readonly grpcConfig: GRPCServerConfig
    ) {
        this.logger = new LoggerWinston();
    }

    public async loadProtoAndStartServer(protoPath: string) {
        // load proto

        const server = new Server({
            "grpc.max_send_message_length": -1,
            "grpc.max_receive_message_length": -1
        })

        server.bindAsync(`0.0.0.0:${this.grpcConfig.port}`, ServerCredentials.createInsecure(), (error, port) => {
            if (error) {
                this.logger.error("failed to start server gRPC.")
            }

            this.logger.info(`start server gRPC on port: ${port}`)
        })
    }
}

injected(PaymentServiceGRPCServer, LOGGER_WINSTON_TOKEN, GRPC_SERVER_CONFIG_TOKEN);

export const PAYMENT_SERVICE_GRPC_SERVER_TOKEN = token<PaymentServiceGRPCServer>("PaymentServiceGRPCServer");
