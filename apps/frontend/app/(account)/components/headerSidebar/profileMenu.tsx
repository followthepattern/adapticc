import { Fragment, useContext } from 'react'

import { Menu, Transition } from '@headlessui/react'
import ChevronDownIcon from '@/app/icons/chevronDown';
import UserContext from '@/components/userContext'
import { Link } from 'react-router-dom'
import { useTokenStore } from '@/lib/store'

interface ProfileMenuProperties { }

export default function ProfileMenu(props: ProfileMenuProperties) {
    const userProfile = useContext(UserContext);
    const { removeToken } = useTokenStore();

    const menuItemClassName = "block px-3 py-1 text-sm hover:bg-gray-50";

    return (
        <Menu as="div" className="relative">
            <Menu.Button className="flex items-center">
                <span className="flex items-center">
                    <span className="text-sm font-semibold" aria-hidden="true">
                        {userProfile.firstName} {userProfile.lastName}
                    </span>
                    <ChevronDownIcon className="w-3 h-3 ml-2 text-gray-400" aria-hidden="true" />
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
                <Menu.Items className="absolute right-0 z-10 w-32 py-2 mt-5 bg-white border border-gray-200 rounded-md shadow-lg">
                    <Menu.Item key="your-profile">
                        <Link
                            to="/profile"
                            className={menuItemClassName}
                        >
                            Your Profile
                        </Link>
                    </Menu.Item>
                    <Menu.Item key="sign-out">
                        <button onClick={() => removeToken()}
                            className={menuItemClassName}
                        >
                            Sign Out
                        </button>
                    </Menu.Item>
                </Menu.Items>
            </Transition>
        </Menu>
    )
}