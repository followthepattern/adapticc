import NewResourceLink from './components/newResourceLink';
import SearchInput from './components/searchInput';
import SortButton, { SortButtonProperties } from './components/sortButton';

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

interface SectionHeaderProperties extends SearchBarProperties, SortButtonProperties {
    resourceName: string;
}

export default function SectionHeader(props: SectionHeaderProperties) {
    return (
        <div className="mx-auto">
            <Header>{props.resourceName}</Header>
            <div className="grid mt-4 sm:justify-between sm:flex gap-x-4 gap-y-2">
                <SearchInput className="sm:max-w-3xl" onChange={props.searchInputOnChange} />
                <div className="grid grid-cols-2 gap-x-2">
                    <NewResourceLink className="grow" {...props} />
                    <SortButton className="grow" {...props} />
                </div>
            </div>
        </div>
    )
}