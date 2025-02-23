import { Container } from "brandi";
import { PAYMENT_SERVICE_GRPC_SERVER_TOKEN, PaymentServiceGRPCServer } from "./server";
import { PAYMENT_SERVICE_HANDLER_FACTORY_TOKEN, PaymentServiceHandlerFactory } from "./handler";

export * from "./server"

export function bindToContainer(container: Container): void {
    container.bind(PAYMENT_SERVICE_GRPC_SERVER_TOKEN).toInstance(PaymentServiceGRPCServer).inSingletonScope();
    container.bind(PAYMENT_SERVICE_HANDLER_FACTORY_TOKEN).toInstance(PaymentServiceHandlerFactory).inSingletonScope();
}
