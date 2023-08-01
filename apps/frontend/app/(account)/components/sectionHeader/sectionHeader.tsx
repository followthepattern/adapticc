import { BarsArrowUpIcon, ChevronDownIcon } from '@heroicons/react/20/solid'
import { Menu, Transition } from '@headlessui/react';
import { Fragment } from 'react';
import SearchInput from './components/searchInput';
import SectionHeadererMenu from './sectionHeadMenu';

function Header({ children }: { children?: any }) {
    return <div className="sm:flex-auto">
        <h1 className="text-base font-semibold leading-6 text-gray-900">{children}</h1>
    </div>
}

function SearchButton() {
    return (
        <Menu className="relative" as="div">
            <Menu.Button className="-ml-px inline-flex items-center gap-x-1.5 rounded-r-md px-3 py-2 text-sm font-semibold text-gray-900 ring-1 ring-inset ring-gray-300 hover:bg-gray-50">
                <span className="sr-only">Sort</span>
                <BarsArrowUpIcon className="-ml-0.5 h-5 w-5 text-gray-400" aria-hidden="true" />
                Sort
                <ChevronDownIcon className="-mr-1 h-5 w-5 text-gray-400" aria-hidden="true" />
            </Menu.Button>
            <Transition
                as={Fragment}
                enter="transition ease-out duration-100"
                enterFrom="transform opacity-0 scale-95"
                enterTo="transform opacity-100 scale-100"
                leave="transition ease-in duration-75"
                leaveFrom="transform opacity-100 scale-100"
                leaveTo="transform opacity-0 scale-95"
            >
                <Menu.Items className="absolute right-0 z-10 mt-2 w-56 origin-top-right rounded-md bg-white shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none">
                    <div className="py-1">
                        <Menu.Item>
                            <span
                                className="text-gray-700 flex justify-between px-4 py-2 text-sm hover:bg-gray-100 hover:text-gray-900"
                            >Create</span>
                        </Menu.Item>
                        <Menu.Item>
                            <span
                                className="text-gray-700 flex justify-between px-4 py-2 text-sm hover:bg-gray-100 hover:text-gray-900"
                            >Export</span>
                        </Menu.Item>
                    </div>
                </Menu.Items>
            </Transition>
        </Menu>
    )
}

interface SearchBarProperties {
    resourceUrl: string
    searchInputOnChange: (s: string) => void
    searchInput?: string
}

function SearchBar({ resourceUrl, searchInputOnChange, searchInput }: SearchBarProperties) {
    return <div className="mt-3 sm:ml-4 sm:mt-0">
        <label htmlFor="desktop-search-candidate" className="sr-only">
            Search
        </label>
        <div className="flex rounded-md shadow-sm">
            <SearchInput onChange={searchInputOnChange} search={searchInput} />
            <SearchButton />
            <SectionHeadererMenu resourceUrl={resourceUrl} />
        </div>
    </div>
}

interface SectionHeaderProperties extends SearchBarProperties {
    resourceName: string;
}

export default function SectionHeader(props: SectionHeaderProperties) {
    return (
        <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
            <div className="sm:flex sm:items-center">
                <Header>{props.resourceName}</Header>
                <SearchBar
                    resourceUrl={props.resourceUrl}
                    searchInputOnChange={props.searchInputOnChange}
                    searchInput={props.searchInput} />
            </div>
        </div>
    )
}