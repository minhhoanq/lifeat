import { injected, token } from "brandi";
import { DATABASE_CONFIG_TOKEN, DatabaseConfig } from "../../config";
import knex, { Knex } from "knex";

export function newKnexInstance(databaseConfig: DatabaseConfig): Knex {
    return knex({
        client: "pg",
        connection: {
            host: databaseConfig.host,
            port: databaseConfig.port,
            user: databaseConfig.user,
            password: databaseConfig.password,
            database: databaseConfig.name,
            ssl: false
        },
        migrations: {
            directory: "../../migrations",
            extension: "ts"
        },
        seeds: {
            directory: "",
            extension: "ts"
        }
    });
}

injected(newKnexInstance, DATABASE_CONFIG_TOKEN)

export const KNEX_INSTANCE_TOKEN = token<Knex>("Knex")
