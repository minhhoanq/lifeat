import winston from 'winston';
import path from 'path';
import { token } from 'brandi';

export class LoggerWinston {
    private logger: winston.Logger;

    constructor() {
        const logFormat = winston.format.combine(
            winston.format.timestamp({
                format: 'YYYY-MM-DD HH:mm:ss'
            }),
            winston.format.colorize(),
            winston.format.printf(({ timestamp, level, message }) => {
                return `[${timestamp}] ${level}: ${message}`;
            })
        );

        const logDir = 'logs';

        this.logger = winston.createLogger({
            level: 'info',
            format: logFormat,
            transports: [
                new winston.transports.Console(),
                new winston.transports.File({
                    filename: path.join(logDir, 'combined.log')
                }),
                new winston.transports.File({
                    filename: path.join(logDir, 'error.log'),
                    level: 'error'
                })
            ]
        });
    }

    info(message: string) {
        this.logger.info(message);
    }

    error(message: string) {
        this.logger.error(message);
    }

    warn(message: string) {
        this.logger.warn(message);
    }

    debug(message: string) {
        this.logger.debug(message);
    }
}

export const LOGGER_WINSTON_TOKEN = token<LoggerWinston>('LoggerWinston')
