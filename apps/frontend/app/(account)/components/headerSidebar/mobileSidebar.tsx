import { Dispatch, Fragment, SetStateAction } from 'react'

import { Transition, Dialog } from "@headlessui/react"
import {
    Cog6ToothIcon,
    XMarkIcon,
} from '@heroicons/react/24/outline'
import classNames from 'classnames'
import { Link, useLocation } from "react-router-dom"
import { IsSelected, NavigationItem } from './navbar/navigation'
import Navbar from './navbar/navbar'

interface MobileSidebarProperties {
    sidebarOpen: boolean
    setSidebarOpen: Dispatch<SetStateAction<boolean>>
    navigationItems: NavigationItem[]
}

export default function MobileSidebar(props: MobileSidebarProperties) {
    const pathname = useLocation().pathname;


    return (
        <Transition.Root show={props.sidebarOpen} as={Fragment}>
            <Dialog as="div" className="relative z-50 lg:hidden" onClose={props.setSidebarOpen}>
                <Transition.Child
                    as={Fragment}
                    enter="transition-opacity ease-linear duration-300"
                    enterFrom="opacity-0"
                    enterTo="opacity-100"
                    leave="transition-opacity ease-linear duration-300"
                    leaveFrom="opacity-100"
                    leaveTo="opacity-0"
                >
                    <div className="fixed inset-0 bg-gray-900/80" />
                </Transition.Child>

                <div className="fixed inset-0 flex">
                    <Transition.Child
                        as={Fragment}
                        enter="transition ease-in-out duration-300 transform"
                        enterFrom="-translate-x-full"
                        enterTo="translate-x-0"
                        leave="transition ease-in-out duration-300 transform"
                        leaveFrom="translate-x-0"
                        leaveTo="-translate-x-full"
                    >
                        <Dialog.Panel className="relative flex flex-1 w-full max-w-xs mr-16">
                            <Transition.Child
                                as={Fragment}
                                enter="ease-in-out duration-300"
                                enterFrom="opacity-0"
                                enterTo="opacity-100"
                                leave="ease-in-out duration-300"
                                leaveFrom="opacity-100"
                                leaveTo="opacity-0"
                            >
                                <div className="absolute top-0 flex justify-center w-16 pt-5 left-full">
                                    <button type="button" className="-m-2.5 p-2.5" onClick={() => props.setSidebarOpen(false)}>
                                        <XMarkIcon className="w-6 h-6 text-white" aria-hidden="true" />
                                    </button>
                                </div>
                            </Transition.Child>
                            <div className="flex flex-col px-6 pb-4 overflow-y-auto bg-blue-600 grow gap-y-5">
                                <div className="flex items-center h-16 shrink-0"></div>
                                <Navbar />
                            </div>
                        </Dialog.Panel>
                    </Transition.Child>
                </div>
            </Dialog>
        </Transition.Root>
    )
}