import { token } from "brandi";
import { DatabaseConfig } from "./database";

export class Config {
    public databaseConfig = new DatabaseConfig();

    public static fromEnv(): Config {
        const config = new Config()
        config.databaseConfig = DatabaseConfig.fromEnv();

        return config;
    }
}

export const CONFIG_TOKEN = token<Config>("Config")
