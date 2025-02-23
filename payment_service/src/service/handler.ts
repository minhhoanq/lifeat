import { sendUnaryData, status } from "@grpc/grpc-js";
import { PaymentServiceHandlers } from "../proto/gen/payment_service/PaymentService";
import { ErrorWithStatus } from "../utils";
import { CreatePaymentResponse } from "../proto/gen/payment_service/CreatePaymentResponse";
import { token } from "brandi";

export class PaymentServiceHandlerFactory {
    constructor() { }

    public handlers(): PaymentServiceHandlers {
        const handler: PaymentServiceHandlers = {
            CreatePayment: async (call, callback) => {
                const req = call.request;

                console.log("req: ", req);

                const response: CreatePaymentResponse = {
                    payment: {}
                }

                callback(null, response)
            }
        }

        return handler
    }

    public handlerError(error: unknown, callback: sendUnaryData<any>) {
        if (error instanceof ErrorWithStatus) {
            callback({ message: error.message, code: error.status });
        } else if (error instanceof Error) {
            callback({ message: error.message, code: status.INTERNAL });
        } else {
            callback({ code: status.INTERNAL });
        }
    }
}

export const PAYMENT_SERVICE_HANDLER_FACTORY_TOKEN = token<PaymentServiceHandlerFactory>("PaymentServiceHandlerFactory");
