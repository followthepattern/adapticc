import { Dispatch, SetStateAction } from 'react'
import {
    Bars3Icon,
} from '@heroicons/react/24/outline'
import Searchbar from './searchbar';
import ProfileMenu from './profileMenu';

interface HeaderProperties {
    setSidebarOpen: Dispatch<SetStateAction<boolean>>
    sidebarOpen: boolean
}

export default function Header(props: HeaderProperties) {

    return (
        <div className="sticky top-0 z-40 flex items-center h-16 px-4 bg-white border-b border-gray-200 shadow-sm shrink-0 gap-x-4 sm:gap-x-6 sm:px-6 lg:px-8">
            <button type="button" className="-m-2.5 p-2.5 text-gray-700 lg:hidden" onClick={() => {  props.setSidebarOpen(!props.sidebarOpen) }}>
                <span className="sr-only">Open sidebar</span>
                <Bars3Icon className="w-6 h-6" aria-hidden="true" />
            </button>

            {/* Separator */}
            <div className="w-px h-6 bg-gray-900/10 lg:hidden" aria-hidden="true" />

            <div className="flex self-stretch flex-1 gap-x-4 lg:gap-x-6">
                <Searchbar />
                <div className="flex items-center gap-x-4 lg:gap-x-6">

                    {/* Separator */}
                    <div className="hidden lg:block lg:h-6 lg:w-px lg:bg-gray-900/10" aria-hidden="true" />

                    {/* Profile dropdown */}
                    <ProfileMenu />
                </div>
            </div>
        </div>
    )
}