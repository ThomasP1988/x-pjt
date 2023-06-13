
import { ChangeEvent, SyntheticEvent, MouseEvent, useState, useEffect } from 'react';
import SearchIcon from '@mui/icons-material/Search';
import { useDebounce } from 'shared/lib/hooks/useDebounce';
import { ThemeProvider, Theme, TextField, TextFieldProps } from '@mui/material';

import { Suggestions } from '../search/Suggestions';
import { useLazyQuery } from '@apollo/client';
import { SEARCH_COLLECTION } from 'shared/repositories/collection/queries/search';
import { searchCollections, searchCollections_searchCollections, searchCollectionsVariables } from 'shared/repositories/collection/queries/__generated__/searchCollections';
import { Collection } from 'shared/repositories/collection/__generated__/Collection';

type Props = {
    onSelect: (item: Collection) => void
} & Omit<TextFieldProps, "onSelect">

export const SelectCollection = ({ onSelect, ...textFieldProps }: Props) => {
    const [search, setSearch] = useState<string>("");
    const [anchorEl, setAnchorEl] = useState<HTMLInputElement | null>(null);
    const debouncedSearch = useDebounce(search, 500);
    const [triggerSearch, { data, error, loading }] = useLazyQuery<searchCollections, searchCollectionsVariables>(SEARCH_COLLECTION);

    useEffect(() => {
        console.log("debouncedSearch", debouncedSearch);
        if (debouncedSearch && debouncedSearch !== "") {
            triggerSearch({
                variables: {
                    text: debouncedSearch
                }
            })
        }
    }, [debouncedSearch, triggerSearch])

    function onChange(event: ChangeEvent<HTMLInputElement>): void {
        const { value } = event.target;
        setSearch(value)
    }

    return <>
        <TextField size="small" variant="outlined" label="Select collection"
            value={search} onChange={onChange}
            onSelect={(event: SyntheticEvent<HTMLInputElement>) => setAnchorEl(event.currentTarget)}
            onClick={(event: MouseEvent<HTMLInputElement>) => setAnchorEl(event.currentTarget)}
            {...textFieldProps}
        />
        <Suggestions search={debouncedSearch} anchor={anchorEl}
            error={error} loading={loading}
            data={(data?.searchCollections as searchCollections_searchCollections)}
            onClick={onSelect}
        />
    </>
}
