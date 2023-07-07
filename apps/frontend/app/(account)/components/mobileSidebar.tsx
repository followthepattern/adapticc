import { Dispatch, Fragment, SetStateAction, useState } from 'react'

import { Transition, Dialog } from "@headlessui/react"
import {
    CalendarIcon,
    ChartPieIcon,
    Cog6ToothIcon,
    DocumentDuplicateIcon,
    FolderIcon,
    HomeIcon,
    UsersIcon,
    XMarkIcon,
} from '@heroicons/react/24/outline'
import classNames from 'classnames'
import Link from 'next/link'
import { IsSelected, NavigationItem } from './navigation'
import { usePathname } from 'next/navigation'

interface MobileSidebarProperties {
    sidebarOpen: boolean,
    setSidebarOpen: Dispatch<SetStateAction<boolean>>,
    navigationItems: NavigationItem[]
}

export default function MobileSidebar(props: MobileSidebarProperties) {
    const pathname = usePathname();

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
                        <Dialog.Panel className="relative mr-16 flex w-full max-w-xs flex-1">
                            <Transition.Child
                                as={Fragment}
                                enter="ease-in-out duration-300"
                                enterFrom="opacity-0"
                                enterTo="opacity-100"
                                leave="ease-in-out duration-300"
                                leaveFrom="opacity-100"
                                leaveTo="opacity-0"
                            >
                                <div className="absolute left-full top-0 flex w-16 justify-center pt-5">
                                    <button type="button" className="-m-2.5 p-2.5" onClick={() => props.setSidebarOpen(false)}>
                                        <span className="sr-only">Close sidebar</span>
                                        <XMarkIcon className="h-6 w-6 text-white" aria-hidden="true" />
                                    </button>
                                </div>
                            </Transition.Child>
                            {/* Sidebar component, swap this element with another sidebar if you like */}
                            <div className="flex grow flex-col gap-y-5 overflow-y-auto bg-indigo-600 px-6 pb-4">
                                <div className="flex h-16 shrink-0 items-center">
                                </div>
                                <nav className="flex flex-1 flex-col">
                                    <ul role="list" className="flex flex-1 flex-col gap-y-7">
                                        <li>
                                            <ul role="list" className="-mx-2 space-y-1">
                                                {props.navigationItems.map((item) => {
                                                    const current = IsSelected(pathname, item.href)
                                                    return (
                                                        <li key={item.name}>
                                                            <Link
                                                                href={item.href}
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
                                        <li className="mt-auto">
                                            <a
                                                href="/settings"
                                                className="group -mx-2 flex gap-x-3 rounded-md p-2 text-sm font-semibold leading-6 text-indigo-200 hover:bg-indigo-700 hover:text-white"
                                            >
                                                <Cog6ToothIcon
                                                    className="h-6 w-6 shrink-0 text-indigo-200 group-hover:text-white"
                                                    aria-hidden="true"
                                                />
                                                Settings
                                            </a>
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