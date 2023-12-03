import {
    Cog6ToothIcon,
} from '@heroicons/react/24/outline'

import classNames from 'classnames'
import { IsSelected, NavigationItem } from './navigation'
import {Link, useLocation} from 'react-router-dom'
interface StaticSidebarPorperties {
    navigationItems: NavigationItem[]
}

export default function StaticSidebar(props: StaticSidebarPorperties) {
    const pathname = useLocation().pathname;

    return (
        <div className="hidden lg:fixed lg:inset-y-0 lg:z-50 lg:flex lg:w-72 lg:flex-col">
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
                                }
                                )}
                            </ul>
                        </li>
                    </ul>
                </nav>
            </div>
        </div>
    )
}