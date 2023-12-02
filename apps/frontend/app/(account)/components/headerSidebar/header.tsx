import { Dispatch, SetStateAction } from 'react'
import {
    Bars3Icon,
} from '@heroicons/react/24/outline'
import ProfileMenu from './profileMenu';

interface HeaderProperties {
    setSidebarOpen: Dispatch<SetStateAction<boolean>>
    sidebarOpen: boolean
}

export default function Header(props: HeaderProperties) {

    return (
        <div className="sticky top-0 z-40 flex items-center h-16 px-6 bg-white border-b border-gray-200 shadow-sm shrink-0 gap-x-4 sm:gap-x-6 sm:px-6 lg:px-8">
            <button type="button" className="-m-2.5 p-2.5 text-gray-700 lg:hidden" onClick={() => { props.setSidebarOpen(!props.sidebarOpen) }}>
                <span className="sr-only">Open sidebar</span>
                <Bars3Icon className="w-6 h-6" aria-hidden="true" />
            </button>

            <div className="flex justify-end w-full gap-x-4">
                <div className="flex w-px h-6 lg:bg-gray-400" aria-hidden="true" />
                <ProfileMenu />
            </div>
        </div>
    )
}