import { token } from "brandi";

export class GRPCServerConfig {
    public port = 9003;

    public static fromEnv(): GRPCServerConfig {
        const config = new GRPCServerConfig();
        const portEnv = process.env.GRPC_SERVER_PORT;
        if (portEnv !== undefined) {
            config.port = +portEnv
        }
        return config
    }
}

export const GRPC_SERVER_CONFIG_TOKEN = token<GRPCServerConfig>("GRPCServerConfig")
