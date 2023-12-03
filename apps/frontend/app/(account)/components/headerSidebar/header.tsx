import { Dispatch, SetStateAction } from 'react'
import BarsIcon from '@/app/icons/bars';
import ProfileMenu from "./profileMenu";

interface HeaderProperties {
    setSidebarOpen: Dispatch<SetStateAction<boolean>>
    sidebarOpen: boolean
}

export default function Header(props: HeaderProperties) {

    return (
        <div className="flex items-center h-16 px-6 bg-white border-b border-gray-200 shadow-sm shrink-0 sm:px-6 lg:px-8">
            <button type="button" className="text-gray-700 lg:hidden" onClick={() => { props.setSidebarOpen(!props.sidebarOpen) }}>
                <BarsIcon className="w-6 h-6" aria-hidden="true" />
            </button>

            <div className="flex items-center justify-end w-full gap-x-4">
                <div className="flex w-px h-8 lg:bg-gray-400" aria-hidden="true" />
                <ProfileMenu/>
            </div>
        </div>
    )
}