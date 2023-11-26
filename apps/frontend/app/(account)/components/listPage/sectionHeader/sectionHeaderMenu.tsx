'use client'

import { Fragment } from 'react'
import { Menu, Transition } from '@headlessui/react'
import { Link } from 'react-router-dom'

interface SectionHeaderMenuProperties {
    resourceUrl: string;
}

export default function SectionHeaderMenu({ resourceUrl }: SectionHeaderMenuProperties) {

    return (
        <Menu as="div" className="relative justify-center inline-block">
            <Menu.Button className="h-full px-4 py-1 border border-gray-300 rounded-lg hover:bg-gray-100 hover:text-gray-700">
                Actions
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
                <Menu.Items className="absolute right-0 z-10 w-56 shadow-lg ring-0">
                    <div className="py-1">
                        <Menu.Item>
                            <Link
                                to={`${resourceUrl}/new`}
                                className="flex px-4 py-2 text-sm hover:bg-gray-100 hover:text-gray-900"
                            >
                                <span>Create</span>
                            </Link>
                        </Menu.Item>
                    </div>
                </Menu.Items>
            </Transition>
        </Menu>
    )
}
