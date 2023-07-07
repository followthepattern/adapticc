import { BarsArrowUpIcon, ChevronDownIcon, MagnifyingGlassIcon } from '@heroicons/react/20/solid'
import SectionHeadingMenu from "./sectionHeadingMenu";
import classNames from 'classnames';
import { Menu, Transition } from '@headlessui/react';
import { Fragment } from 'react';

interface SectionHeadingProperties {
    resourceName: string;
    resourceUrl: string;
}

function Header({ children }: { children?: any }) {
    return <div className="sm:flex-auto">
        <h1 className="text-base font-semibold leading-6 text-gray-900">{children}</h1>
    </div>
}

function SearchInput() {
    const commonClasses = "w-full rounded-none rounded-l-md border-0 py-1.5 pl-10 text-gray-900 ring-1 ring-gray-300 ring-inset placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600";

    return (
        <div className="relative flex-grow focus-within:z-10">
            <div className="pointer-events-none absolute inset-y-0 left-0 flex items-center pl-3">
                <MagnifyingGlassIcon className="h-5 w-5 text-gray-400" aria-hidden="true" />
            </div>
            <input
                type="text"
                className={classNames("block ring-inset sm:hidden", commonClasses)}
                placeholder="Search"
            />
            <input
                type="text"
                className={classNames("hidden sm:block text-sm leading-6", commonClasses)}
                placeholder="Search"
            />
        </div>
    )
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
}

function SearchBar({resourceUrl}: SearchBarProperties) {
    return <div className="mt-3 sm:ml-4 sm:mt-0">
        <label htmlFor="desktop-search-candidate" className="sr-only">
            Search
        </label>
        <div className="flex rounded-md shadow-sm">
            <SearchInput />
            <SearchButton />
            <SectionHeadingMenu resourceUrl={resourceUrl} />
        </div>
    </div>
}

export default function SectionHeading(props: SectionHeadingProperties) {
    return (
        <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
            <div className="sm:flex sm:items-center">
                <Header>{props.resourceName}</Header>
                <SearchBar resourceUrl={props.resourceUrl} />
            </div>
        </div>
    )
}