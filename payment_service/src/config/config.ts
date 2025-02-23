import { token } from "brandi";
import { DatabaseConfig } from "./database";
import { GRPCServerConfig } from "./grpc_server";

export class Config {
    public databaseConfig = new DatabaseConfig();
    public grpcServerConfig = new GRPCServerConfig();

    public static fromEnv(): Config {
        const config = new Config()
        config.databaseConfig = DatabaseConfig.fromEnv();
        config.grpcServerConfig = GRPCServerConfig.fromEnv();

        return config;
    }
}

export const CONFIG_TOKEN = token<Config>("Config")
