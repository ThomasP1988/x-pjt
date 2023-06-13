
import { useNavigate, Link as LinkRouter } from 'react-router-dom';
import Breadcrumbs from '@mui/material/Breadcrumbs';
import Link from '@mui/material/Link';
import Typography from '@mui/material/Typography';
import { Alert, AlertTitle } from '@mui/material';
import { SelectNFTOrCrypto } from "shared/components/SelectNFTOrCrypto";

export const Deposit = () => {
    const navigate = useNavigate()

    const gotoDepositNFT = () => {
        navigate('/deposit/nft')
    }

    return (<>
        <Breadcrumbs aria-label="breadcrumb">
            <Link underline="hover" color="inherit" component={LinkRouter} to="/">
                Home
            </Link>
            <Typography color="text.primary">Deposit</Typography>
        </Breadcrumbs>
        <Typography variant="h4" sx={{ marginTop: 3 }} gutterBottom>Add tokens to your account</Typography>
        <Alert severity="info" sx={{ marginBottom: 3 }}>
            <AlertTitle>The collection of your NFT is not on the marketplace yet?</AlertTitle>
            We dynamically accept collection, most of them will be added <strong>without validation</strong> from our part. In order to keep quality on the marketplace, some will be subject to validation.
        </Alert>
        <SelectNFTOrCrypto onNFTClick={gotoDepositNFT} />
    </>)
}