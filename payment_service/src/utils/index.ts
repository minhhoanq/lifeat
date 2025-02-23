import { Container } from "brandi";
import { LOGGER_WINSTON_TOKEN, LoggerWinston } from "./logger";

export * from "./logger";
export * from "./errors";

export function bindToContainer(container: Container): void {
    container.bind(LOGGER_WINSTON_TOKEN).toInstance(LoggerWinston).inSingletonScope();
}
