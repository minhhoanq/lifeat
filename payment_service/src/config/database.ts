import { token } from 'brandi';

export class DatabaseConfig {
    public host = "127.0.0.1";
    public port = 5435;
    public user = "root";
    public password = "secret";
    public name = "";

    public static fromEnv(): DatabaseConfig {
        const config = new DatabaseConfig()
        if (process.env.DB_HOST !== undefined) {
            config.host = process.env.DB_HOST;
        }
        if (process.env.DB_PORT !== undefined) {
            config.port = +process.env.DB_PORT;
        }
        if (process.env.DB_USER !== undefined) {
            config.user = process.env.DB_USER;
        }
        if (process.env.DB_PASSWORD !== undefined) {
            config.password = process.env.DB_PASSWORD;
        }
        if (process.env.DB_NAME !== undefined) {
            config.name = process.env.DB_NAME;
        }
        return config;
    }
}

export const DATABASE_CONFIG_TOKEN = token<DatabaseConfig>("DatabaseConfig");
