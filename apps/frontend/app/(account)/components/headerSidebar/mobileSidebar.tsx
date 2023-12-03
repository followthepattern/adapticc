import { Dispatch, Fragment, SetStateAction } from 'react'

import { Transition, Dialog } from "@headlessui/react"
import {
    Cog6ToothIcon,
    XMarkIcon,
} from '@heroicons/react/24/outline'
import classNames from 'classnames'
import {Link, useLocation} from "react-router-dom"
import { IsSelected, NavigationItem } from './navigation'

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
                                        <span className="sr-only">Close sidebar</span>
                                        <XMarkIcon className="w-6 h-6 text-white" aria-hidden="true" />
                                    </button>
                                </div>
                            </Transition.Child>
                            {/* Sidebar component, swap this element with another sidebar if you like */}
                            <div className="flex flex-col px-6 pb-4 overflow-y-auto bg-indigo-600 grow gap-y-5">
                                <div className="flex items-center h-16 shrink-0">
                                </div>
                                <nav className="flex flex-col flex-1">
                                    <ul role="list" className="flex flex-col flex-1 gap-y-7">
                                        <li>
                                            <ul role="list" className="-mx-2 space-y-1">
                                                {props.navigationItems.map((item) => {
                                                    const current = IsSelected(pathname, item.href)
                                                    return (
                                                        <li key={item.name}>
                                                            <Link
                                                                to={item.href}
                                                                className={classNames(
                                                                    current
                                                                        ? 'bg-indigo-700 text-white'
                                                                        : 'text-indigo-200 hover:text-white hover:bg-indigo-700',
                                                                    'group flex gap-x-3 rounded-md p-2 text-sm leading-6 font-semibold'
                                                                )}
                                                            >
                                                                <item.icon
                                                                    className={classNames(
                                                                        current ? 'text-white' : 'text-indigo-200 group-hover:text-white',
                                                                        'h-6 w-6 shrink-0'
                                                                    )}
                                                                    aria-hidden="true"
                                                                />
                                                                {item.name}
                                                            </Link>
                                                        </li>
                                                    )
                                                })}
                                            </ul>
                                        </li>
                                    </ul>
                                </nav>
                            </div>
                        </Dialog.Panel>
                    </Transition.Child>
                </div>
            </Dialog>
        </Transition.Root>
    )
}