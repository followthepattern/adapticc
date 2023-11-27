import SearchInput from './components/searchInput';
import SortButton, { SortButtonProperties } from './components/sortButton';
import SectionHeaderMenu from './sectionHeaderMenu';

function Header({ children }: { children?: any }) {
    return <div className="sm:flex-auto">
        <h1 className="text-base font-semibold leading-6 text-gray-900">{children}</h1>
    </div>
}

interface SearchBarProperties extends SortButtonProperties {
    resourceUrl: string
    searchInputOnChange: (s: string) => void
    searchInput?: string
}

function SearchBar({
    resourceUrl,
    searchInputOnChange,
    searchInput,
    sortByLables,
    sortOnChange,
    selectedSortLabel,
}: SearchBarProperties) {
    return <div className="flex flex-col mt-3 sm:flex-row sm:flex-auto gap-y-2 sm:gap-x-2 sm:ml-4 sm:mt-0">
        <SearchInput onChange={searchInputOnChange} search={searchInput} />
        <div className="flex gap-x-2">
            <SortButton sortByLables={sortByLables} sortOnChange={sortOnChange} selectedSortLabel={selectedSortLabel} />
            <SectionHeaderMenu resourceUrl={resourceUrl} />
        </div>
    </div>
}

interface SectionHeaderProperties extends SearchBarProperties, SortButtonProperties {
    resourceName: string;
}

export default function SectionHeader(props: SectionHeaderProperties) {
    return (
        <div className="px-4 mx-auto max-w-7xl sm:px-6 lg:px-8">
            <div className="sm:flex">
                <Header>{props.resourceName}</Header>
                <SearchBar
                    resourceUrl={props.resourceUrl}
                    searchInputOnChange={props.searchInputOnChange}
                    searchInput={props.searchInput}
                    sortByLables={props.sortByLables}
                    sortOnChange={props.sortOnChange}
                    selectedSortLabel={props.selectedSortLabel}
                />
            </div>
        </div>
    )
}