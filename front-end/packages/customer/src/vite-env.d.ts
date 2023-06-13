declare module 'supports-webp-sync';

/// <reference types="vite/client" />

interface WindowChain {
    ethereum?: {
        isMetaMask: true;
        request: (...args: any[]) => Promise<any>;
        chainId: string;
    };
}
