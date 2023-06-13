import { Skeleton } from "@mui/material";
import { useState, DetailedHTMLProps, ImgHTMLAttributes, useEffect } from "react";
import { GetImageURL } from "../../repositories/media/image";

type Args = {
    path?: string | null,
    maxHeight: number,
    maxWidth: number,
} & DetailedHTMLProps<ImgHTMLAttributes<HTMLImageElement>, HTMLImageElement>

export const ImageFromBucket = ({ path, maxHeight, maxWidth, style, ...imgProps }: Args) => {
    const [url, setUrl] = useState<string | null>(null)

    useEffect(() => {
        if (!url && path) {
            GetImageURL(path).then(setUrl)
        } else if(!path) {
            setUrl(null);
        }
    }, [path, url, setUrl]);
    return url ? <img  {...imgProps} style={{ maxWidth, maxHeight, ...style }} src={url} /> : <Skeleton variant="rectangular" width={maxWidth} height={maxHeight} style={{ maxWidth, maxHeight, ...style }}  />

}