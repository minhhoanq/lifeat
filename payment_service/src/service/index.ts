import { Container } from "brandi";
import { PAYMENT_SERVICE_GRPC_SERVER_TOKEN, PaymentServiceGRPCServer } from "./server";

export * from "./server"

export function bindToContainer(container: Container): void {
    container.bind(PAYMENT_SERVICE_GRPC_SERVER_TOKEN).toInstance(PaymentServiceGRPCServer).inSingletonScope();
}
