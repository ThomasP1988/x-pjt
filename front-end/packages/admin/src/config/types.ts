export type Config = {
    cognito: {
        REGION: string,
        USER_POOL_ID: string,
        APP_CLIENT_ID: string,
        IDENTITY_POOL_ID: string
    },
    s3: {
        REGION: string,
		BUCKET: string
    }
}


