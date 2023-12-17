import { config as PrdConfig } from './config.prod'
import { config as DevConfig } from './config.dev'
import type { Config } from './types';

const MODE = import.meta.env.MODE

export function getConfig(): Config {
    switch (MODE) {
        case 'development':
            return DevConfig;
        case 'production':
            return PrdConfig;
        default:
            return DevConfig;
    }
}