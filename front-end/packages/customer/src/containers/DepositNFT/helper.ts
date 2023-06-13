import { Contract } from 'web3-eth-contract';

export type GetBatchOfTokenArgs = {
    batchSize: number,
    page: number,
    total: number,
    contract: Contract,
    account: string,
}

export const GetBatchOfToken = async ({ batchSize, total, contract, account, page }: GetBatchOfTokenArgs): Promise<number[]> => {
    const promises: Promise<number>[] = [];
    let start: number = page * batchSize;
    let end: number = start + batchSize;

    if (end > total) {
        end = total
    }

    for (let i = start; i < end; i++) {
        promises.push(contract.methods.tokenOfOwnerByIndex(account, i).call());
    }

    try {
        return await Promise.all(promises);
    } catch(e) {
        console.log("error getting token ids")
        return Promise.reject(e);
    }
}