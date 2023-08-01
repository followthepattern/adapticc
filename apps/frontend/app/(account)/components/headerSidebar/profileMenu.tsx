import { Fragment, useContext } from 'react'

import { Menu, Transition } from '@headlessui/react'
import { ChevronDownIcon } from '@heroicons/react/24/outline'
import UserContext from '@/components/userContext'
import { Link } from 'react-router-dom'
import { useTokenStore } from '@/lib/store'

interface ProfileMenuProperties { }

export default function ProfileMenu(props: ProfileMenuProperties) {
    const userProfile = useContext(UserContext);
    const { removeToken } = useTokenStore();

    return (
        <Menu as="div" className="relative">
            <Menu.Button className="-m-1.5 flex items-center p-1.5">
                <span className="sr-only">Open user menu</span>
                <span className="hidden lg:flex lg:items-center">
                    <span className="text-sm font-semibold leading-6 text-gray-900" aria-hidden="true">
                        {userProfile.firstName} {userProfile.lastName}
                    </span>
                    <ChevronDownIcon className="ml-2 h-5 w-5 text-gray-400" aria-hidden="true" />
                </span>
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
                <Menu.Items className="absolute right-0 z-10 mt-2.5 w-32 origin-top-right rounded-md bg-white py-2 shadow-lg ring-1 ring-gray-900/5 focus:outline-none">
                    <Menu.Item key="your-profile">
                        <Link
                            to="/profile"
                            className="block px-3 py-1 text-sm leading-6 text-gray-900 hover:bg-gray-50"
                        >
                            Your Profile
                        </Link>
                    </Menu.Item>
                    <Menu.Item key="sign out">
                        <button onClick={() => removeToken()}
                            className="block px-3 py-1 text-sm leading-6 text-gray-900 hover:bg-gray-50"
                        >
                            Sign Out
                        </button>
                    </Menu.Item>
                </Menu.Items>
            </Transition>
        </Menu>
    )
}