'use client'

import { Fragment } from 'react'
import { Menu, Transition } from '@headlessui/react'
import { EllipsisVerticalIcon } from '@heroicons/react/20/solid'
import { Link } from 'react-router-dom'

interface SectionHeadingMenu {
    resourceUrl: string;
}

export default function SectionHeadingMenu({resourceUrl}: SectionHeadingMenu) {

    return (
        <Menu as="div" className="relative sm:ml-3 px-3 py-2 inline-block text-left">
            <div>
                <Menu.Button className="-my-2 flex items-center rounded-full bg-white p-2 text-gray-400 hover:text-gray-600 focus:outline-none focus:ring-2 focus:ring-indigo-500">
                    <span className="sr-only">Open options</span>
                    <EllipsisVerticalIcon className="h-5 w-5" aria-hidden="true" />
                </Menu.Button>
            </div>

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
                            <Link
                                to={`${resourceUrl}/new`}
                                className="text-gray-700 flex justify-between px-4 py-2 text-sm hover:bg-gray-100 hover:text-gray-900"
                            >
                                <span>Create</span>
                            </Link>
                        </Menu.Item>
                        <Menu.Item>
                            <Link
                                to="#"
                                className="text-gray-700 flex justify-between px-4 py-2 text-sm hover:bg-gray-100 hover:text-gray-900"
                            >
                                <span>Export</span>
                            </Link>
                        </Menu.Item>
                        <Menu.Item>
                            <Link
                                to="#"
                                className="text-gray-700 flex justify-between px-4 py-2 text-sm hover:bg-gray-100 hover:text-gray-900"
                            >
                                <span>Delete</span>
                            </Link>
                        </Menu.Item>
                    </div>
                </Menu.Items>
            </Transition>
        </Menu>
    )
}
