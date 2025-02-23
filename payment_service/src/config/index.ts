import { Container } from "brandi";
import { Config, CONFIG_TOKEN } from "./config";
import { DATABASE_CONFIG_TOKEN } from "./database";
import { GRPC_SERVER_CONFIG_TOKEN } from "./grpc_server";

export * from "./config"
export * from "./database"
export * from "./grpc_server"

export function bindToContainer(container: Container): void {
    container.bind(CONFIG_TOKEN).toInstance(Config.fromEnv).inSingletonScope();
    container.bind(DATABASE_CONFIG_TOKEN).toInstance(() => Config.fromEnv().databaseConfig).inSingletonScope();
    container.bind(GRPC_SERVER_CONFIG_TOKEN).toInstance(() => Config.fromEnv().grpcServerConfig).inSingletonScope();
}
