import { Container } from "brandi";
import { KNEX_INSTANCE_TOKEN, newKnexInstance } from "./knex";
import { PAYMENT_DATA_ACCESSOR_TOKEN, PaymentDataAccessorImpl } from "./payment";

export function bindToContainer(container: Container): void {
    container.bind(KNEX_INSTANCE_TOKEN).toInstance(newKnexInstance).inSingletonScope();
    container.bind(PAYMENT_DATA_ACCESSOR_TOKEN).toInstance(PaymentDataAccessorImpl).inSingletonScope();
}