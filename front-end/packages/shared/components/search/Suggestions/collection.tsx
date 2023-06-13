import { Avatar, ListItemAvatar, ListItemText } from "@mui/material"
import { searchCollections_searchCollections_results } from "shared/repositories/collection/queries/__generated__/searchCollections"
import { ImageFromBucket } from "../../ImageFromBucket";
type Props = {
    item: searchCollections_searchCollections_results
}

export const Collection = ({ item }: Props) => {
    return (<>
        <ListItemAvatar sx={{ marginRight: 1 }}>
           <ImageFromBucket path={item.thumbnailPath} maxHeight={80} maxWidth={80} style={{ verticalAlign: "middle"}} />
        </ListItemAvatar>
        <ListItemText primary={item.name} secondary={item.description} />
    </>
    )
}