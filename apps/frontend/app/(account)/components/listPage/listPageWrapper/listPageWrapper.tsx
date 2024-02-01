import { SetURLSearchParams, useSearchParams } from "react-router-dom";
import { GetPageFromSearchParams, GetPageSizeFromSearchParams, GetSearch, GetSortLabel, SortLabel } from "./listingFunctions";

export interface ListPageComponentProperties {
    searchParams: URLSearchParams
    setSearchParams: SetURLSearchParams
    paginationProperties: PaginationProperties
    filterProps: FilterProperties
    sortProps: SortProperties
}

export interface SortProperties {
    sortLabel?: SortLabel
}

export interface FilterProperties {
    searchString: string
}

export interface PaginationProperties {
    page: number
    pageSize: number
}

interface ListPageWrapperProperties {
    Component: React.ComponentType<ListPageComponentProperties>
}

export const ListPageWrapper: React.FC<ListPageWrapperProperties> = ({ Component }: ListPageWrapperProperties, ...props) => {
    const [searchParams, setSearchParams] = useSearchParams();
    const searchString = GetSearch(searchParams);
    const sortLabel = GetSortLabel(searchParams);

    const page = GetPageFromSearchParams(searchParams);
    const pageSize = GetPageSizeFromSearchParams(searchParams);


    const filt = {searchString}
    const sort = {sortLabel}

    return (
        <Component
            searchParams={searchParams}
            setSearchParams={setSearchParams}
            filterProps={filt}
            sortProps={sort}
            paginationProperties={{page, pageSize}}
            {...props}
        />
    )
}