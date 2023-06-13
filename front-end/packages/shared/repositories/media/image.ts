import { Storage } from 'aws-amplify';

export const GetImageURL = (path:string): Promise<string> => {
	return Storage.get(path, {
        customPrefix: {
            public: ''
        }
    })
}