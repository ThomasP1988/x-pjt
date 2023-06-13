import { useState, ChangeEventHandler, ChangeEvent, useEffect, KeyboardEventHandler, KeyboardEvent, FC } from "react"
import { Typography, OutlinedInput, InputAdornment, IconButton, Stack, Link, OutlinedInputProps } from "@mui/material"
import CheckCircleIcon from '@mui/icons-material/CheckCircle';
import CancelIcon from '@mui/icons-material/Cancel';
import EditIcon from '@mui/icons-material/Edit';
import ClickAwayListener from '@mui/base/ClickAwayListener';

type PropertyFieldProps<T> = Omit<OutlinedInputProps, "onChange"> & {
    title: string;
    keyProperty: string;
    value: T;
    onChange: ({ keyProperty, newValue }: { keyProperty: string, newValue: T }) => void
}

export function PropertyToField<T extends string = string>({ title, value, keyProperty, onChange, ...inputProps }: PropertyFieldProps<T>) {
    const [showField, setShowField] = useState<boolean>(false);
    const [localValue, setLocalValue] = useState<T>(value);

    useEffect(() => {
        setLocalValue(value);
    }, [value])


    const onChangeHandler: ChangeEventHandler<HTMLInputElement> = (event: ChangeEvent<HTMLInputElement>) => {
        setLocalValue(event.target.value as T);
    }

    const validate = () => {
        console.log("validatekey", keyProperty);
        onChange({
            keyProperty,
            newValue: localValue
        })
        setShowField(false);
    }

    const cancel = () => {
        setLocalValue(value);
        setShowField(false);
    }

    const onPressEnter: KeyboardEventHandler<HTMLDivElement> = (e: KeyboardEvent<HTMLDivElement>) => {
        e.preventDefault;
        if (e.key === 'Enter') {
            validate();
        }
    }

    const hasValue: boolean = Boolean(value) && value !== "";

    return <Stack sx={{ width: "100%" }}>
        <Typography component={Stack} direction="row" alignItems="center"
            variant="body1"
            sx={{ marginBotton: 1 }}
            onClick={() => setShowField(!showField)}
        >
            <EditIcon fontSize="inherit" sx={{ marginRight: 1 }} />
            {title}
        </Typography>
        <Stack alignContent="center" alignItems="center" sx={{ minHeight: 35, marginBottom: 1 }} direction="row">
            {
                showField ?
                    <ClickAwayListener onClickAway={() => setShowField(false)}>
                        <OutlinedInput type="text" size="small" fullWidth value={localValue} onChange={onChangeHandler}
                            {...inputProps}
                            onKeyPress={onPressEnter}
                            id={keyProperty} name={keyProperty}
                            endAdornment={
                                <InputAdornment position="end">
                                    <IconButton
                                        aria-label="toggle password visibility"
                                        onClick={validate}
                                        edge="end"
                                        
                                        color="success"
                                    >
                                        <CheckCircleIcon />
                                    </IconButton>
                                    <IconButton
                                        aria-label="toggle password visibility"
                                        onClick={cancel}
                                        edge="end"
                                        color="error"
                                    >
                                        <CancelIcon />
                                    </IconButton>
                                </InputAdornment>
                            }
                        />
                    </ClickAwayListener>
                    : <Typography variant="body2" color={hasValue ? "default" : "secondary"} onClick={() => setShowField(!showField)} sx={{ width: "100%" }}>
                        {
                            hasValue ? value : "click here to edit"
                        }
                    </Typography>
            }
        </Stack>
    </Stack>
}