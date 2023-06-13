export enum Blockchains {
    ETHEREUM = "ethereum",
    POLYGON = "polygon",
    PALM = "palm",
}

export type Config = {
    apiGateway: {
        REGION: string,
        HTTP: string,
        WS: string
    },
    appsync: {
        REGION: string,
        HTTP: string,
        WS: string
    },
    blockchain: {
        [Blockchains.ETHEREUM]: BlockchainConfig,
        [Blockchains.POLYGON]: BlockchainConfig,
        [Blockchains.PALM]: BlockchainConfig,
    }
}

export type BlockchainConfig = {
    name: string,
    chainId: number,
    network: string,
    explorer: string,
    currency?: {
        name: string,
        symbol: string,
        decimals: number,
    },
    svg: string
}

