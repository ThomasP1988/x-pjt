import { ChangeEvent, SyntheticEvent, MouseEvent, useState, useEffect } from 'react';
import { Search, SearchIconWrapper, StyledInputBase } from "../../../styles/theme/search";
import SearchIcon from '@mui/icons-material/Search';
import { useDebounce } from 'shared/lib/hooks/useDebounce';
import { ThemeProvider, Theme } from '@mui/material';

import { Suggestions } from '../Suggestions';
import { useLazyQuery } from '@apollo/client';
import { SEARCH_COLLECTION } from 'shared/repositories/collection/queries/search';
import { searchCollections, searchCollections_searchCollections, searchCollectionsVariables } from 'shared/repositories/collection/queries/__generated__/searchCollections';

export type Props = {
    theme?: Theme,
}

export const MenuSearch = ({ theme }: Props) => {
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

    const suggestions = <Suggestions search={debouncedSearch} anchor={anchorEl} error={error} loading={loading} data={(data?.searchCollections as searchCollections_searchCollections)} />;

    return <>
        <Search style={{ marginRight: 20 }}>
            <SearchIconWrapper>
                <SearchIcon />
            </SearchIconWrapper>
            <StyledInputBase
                placeholder="Search…"
                inputProps={{ 'aria-label': 'search' }}
                value={search}
                onChange={onChange}
                onSelect={(event: SyntheticEvent<HTMLInputElement>) => setAnchorEl(event.currentTarget)}
                onClick={(event: MouseEvent<HTMLInputElement>) => setAnchorEl(event.currentTarget)}
            />
        </Search>
        {
            theme ? <ThemeProvider theme={theme}>
                {
                    suggestions
                }
            </ThemeProvider> :
                suggestions
        }

    </>

}