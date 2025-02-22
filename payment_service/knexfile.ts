import type { Knex } from "knex";
import dotenv from "dotenv"
// Update with your config settings.

dotenv.config()

const POSTGRES_HOST = process.env.DB_HOST;
const POSTGRES_PORT = +(process.env.DB_PORT || 5435);
const POSTGRES_DB = process.env.DB_NAME;
const POSTGRES_USER = process.env.DB_USER;
const POSTGRES_PASSWORD = process.env.DB_PASSWORD;

const config: { [key: string]: Knex.Config } = {
  development: {
    client: "postgresql",
    connection: {
      host: POSTGRES_HOST,
      port: POSTGRES_PORT,
      user: POSTGRES_USER,
      password: POSTGRES_PASSWORD,
      database: POSTGRES_DB,
    },
    pool: {
      min: 2,
      max: 4
    },
    migrations: {
      tableName: "knex_migrations",
      loadExtensions: [".ts"]
    }
  },

  staging: {
    client: "postgresql",
    connection: {
      host: POSTGRES_HOST,
      port: POSTGRES_PORT,
      user: POSTGRES_USER,
      password: POSTGRES_PASSWORD,
      database: POSTGRES_DB,
    },
    pool: {
      min: 2,
      max: 4
    },
    migrations: {
      tableName: "knex_migrations",
      loadExtensions: [".ts"]
    }
  },

  production: {
    client: "postgresql",
    connection: {
      host: POSTGRES_HOST,
      port: POSTGRES_PORT,
      user: POSTGRES_USER,
      password: POSTGRES_PASSWORD,
      database: POSTGRES_DB,
    },
    pool: {
      min: 2,
      max: 4
    },
    migrations: {
      tableName: "knex_migrations",
      loadExtensions: [".ts"]
    }
  }

};

module.exports = config;
