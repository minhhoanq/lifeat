import { Container } from "brandi";
import { Config, CONFIG_TOKEN } from "./config";
import { DATABASE_CONFIG_TOKEN } from "./database";

export * from "./config"
export * from "./database"

export function bindToContainer(container: Container): void {
    container.bind(CONFIG_TOKEN).toInstance(Config.fromEnv).inSingletonScope();
    container.bind(DATABASE_CONFIG_TOKEN).toInstance(() => Config.fromEnv().databaseConfig).inSingletonScope();
}
