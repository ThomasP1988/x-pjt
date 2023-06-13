import { useEffect, useState, ReactNode } from "react";
import Card from '@mui/material/Card';
import CardActions from '@mui/material/CardActions';
import CardContent from '@mui/material/CardContent';
import CardMedia from '@mui/material/CardMedia';
import Alert from '@mui/material/Alert';
import Typography from '@mui/material/Typography';
import { OnSubscriptionDataOptions, useQuery, useSubscription } from '@apollo/client';
import { GET_ITEM } from "../../repositories/item/queries/get";
import { Item } from "../../repositories/item/__generated__/item";
import { getItem, getItemVariables } from "../../repositories/item/queries/__generated__/getItem";
import { onUpdatedItem } from "../../repositories/item/subscriptions/update";
import { onUpdatedItem as onUpdatedItemType, onUpdatedItemVariables } from "../../repositories/item/subscriptions/__generated__/onUpdatedItem";
import { Skeleton } from "@mui/material";
import { ImageFromBucket } from "../ImageFromBucket";

type ListNFTItemProps = {
    collectionAddress: string,
    tokenId: number,
    actions?: ReactNode
}

export const ListNFTItem = ({ collectionAddress, tokenId, actions }: ListNFTItemProps) => {

    const [NFTItem, setNFTItem] = useState<Item | null>(null);

    const { data: createSubData, error: createSubError } = useSubscription<onUpdatedItemType, onUpdatedItemVariables>(onUpdatedItem, {
        shouldResubscribe: true,
        variables: {
            collectionAddress,
            tokenId,
        },
        onSubscriptionData: ({ subscriptionData: { data } }: OnSubscriptionDataOptions<onUpdatedItemType>) => {
            if (data?.onUpdatedItem) {
                setNFTItem(data.onUpdatedItem)
            }
        }
    });

    console.log("createSubData", createSubData);
    console.log("createSubError", createSubError);

    const { loading, error, refetch } = useQuery<getItem, getItemVariables>(GET_ITEM, {
        variables: {
            collectionAddress,
            tokenId,
        },
        onCompleted: (data: getItem): void => {
            if (data?.getItem?.id) {
                setNFTItem(data.getItem);
            }
        },
    });


    useEffect(() => {
        if (refetch) {
            refetch({
                collectionAddress,
                tokenId
            })
        }
    }, [collectionAddress, tokenId, refetch])

    return <Card sx={{ maxWidth: 345, marginTop: 2 }}>
        {
            loading ? <Skeleton /> :
                (NFTItem ? <><CardMedia
                    component={ImageFromBucket}
                    path={NFTItem.imagePath}
                    maxHeight={250}
                    maxWidth={250}
                    style={{
                        padding: 20,
                        minHeight: 250,
                        maxHeight: 250,
                        marginLeft: "auto",
                        marginRight: "auto"
                    }}
                    alt={NFTItem.name || ""}
                /></> : <></>)
        }
        <CardContent>
            {
                error && <Alert severity="error">{error.message}</Alert>
            }
            <Typography gutterBottom variant="h5" component="div">
                {
                    loading ? <Skeleton /> :
                        (NFTItem ? <>{NFTItem.id} {NFTItem.name}</> : <></>)
                }
            </Typography>
            {/* <Typography variant="body2" color="text.secondary">
                {
                    loading ? <Skeleton /> :
                        (NFTItem ? <>{NFTItem.description}</> : <></>)
                }
            </Typography> */}
        </CardContent>
        <CardActions>
            {
                actions
            }
        </CardActions>
    </Card>
}