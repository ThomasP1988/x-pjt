import { Property } from 'csstype';
import TextField, {  TextFieldProps } from '@mui/material/TextField';

export type Props = TextFieldProps & {
    step?: string, 
    min?: string, 
    max?: string,
    textAlign?: Property.TextAlign,
}


export const InputLabelInside = (props:  Props) => {

    const {step, min, max, textAlign = "left", ...restProps} = props;

    return <TextField InputLabelProps={{ shrink: false,  }} size="small" {...restProps}
        fullWidth
        InputProps={{
            inputProps: {
                style: { 
                    textAlign,
                },
                ...step ? {
                    step
                } : {},
                ...min ? {
                    min
                } : {},
                ...max ? {
                    max
                } : {},
            },
        
        }}

    />
}