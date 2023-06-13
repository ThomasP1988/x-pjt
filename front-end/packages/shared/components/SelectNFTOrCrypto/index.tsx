import Grid from '@mui/material/Grid';
import Card from '@mui/material/Card';
import CardContent from '@mui/material/CardContent';
import CardMedia from '@mui/material/CardMedia';
import Typography from '@mui/material/Typography';
import { CardActionArea, Tooltip } from '@mui/material';
import { useColors } from "../../lib/hooks/colors";
import HelpIcon from '@mui/icons-material/Help';

export type SelectNFTOrCryptoArgs = {
    onNFTClick?: () => void
}


export const SelectNFTOrCrypto = ({ onNFTClick }: SelectNFTOrCryptoArgs) => {

    const colors = useColors();

    return <><Typography variant="subtitle1" color={colors.grey} gutterBottom>Click on the type of token you wish to deposit</Typography>
        <Grid container direction="row">
            <Grid item xs={6} container justifyContent="center" onClick={onNFTClick} >
                <Card sx={{ width: 375 }}>
                    <CardActionArea>
                        <CardMedia
                            component="img"
                            height="140"
                            image="/img/nft.jpeg"
                            alt="bored ape"
                            sx={{ height: 200 }}
                        />
                        <CardContent>
                            <Typography gutterBottom variant="h5" component="div">
                                NFT <Tooltip title="A non-fungible token (NFT) is a unique and non-interchangeable unit of data stored on a blockchain, a form of digital ledger. NFTs can be associated with reproducible digital files such as photos, videos, and audio." placement="top">
                                    {/* <IconButton> */}
                                    <HelpIcon />
                                    {/* </IconButton> */}
                                </Tooltip>
                            </Typography>
                            <Typography variant="body2" color="text.secondary">
                                Deposit NFT from any collection
                            </Typography>
                        </CardContent>

                    </CardActionArea>
                </Card>
            </Grid>
            <Grid item xs={6} container justifyContent="center">
                <Card sx={{ width: 375 }}>
                    <CardActionArea>
                        <CardMedia
                            component="img"
                            height="140"
                            image="/img/dai.jpeg"
                            alt="bored ape"
                            sx={{ height: 200 }}
                        />
                        <CardContent>
                            <Typography gutterBottom variant="h5" component="div">
                                DAI Stable Coin <Tooltip title="Dai is a stablecoin cryptocurrency which aims to keep its value as close to one United States dollar as possible through an automated system of smart contracts on the Ethereum blockchain." placement="top">
                                    {/* <IconButton> */}
                                    <HelpIcon />
                                    {/* </IconButton> */}
                                </Tooltip>
                            </Typography>
                            <Typography variant="body2" color="text.secondary">
                                Deposit your DAI
                            </Typography>
                        </CardContent>
                    </CardActionArea>
                </Card>
            </Grid>
        </Grid>
    </>
}