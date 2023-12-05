import FolderIcon from '@/app/icons/folder';
import classNames from 'classnames';
import { Link, useLocation } from 'react-router-dom';


export interface NavigationItem {
    name: string
    href: string
    icon: React.ComponentType<{className:string}>
}

export const navigationItems: NavigationItem[] = [
    { name: 'Users', href: '/users', icon: FolderIcon },
    { name: 'Products', href: '/products', icon: FolderIcon },
]

export function IsSelected(location: string, path: string): boolean {
    return location.split("/")[1] === path.split("/")[1]
}

interface NavbarProperties { }

export default function Navbar(props: NavbarProperties) {
    const pathname = useLocation().pathname;

    return (
        <nav className="flex flex-col flex-1">
            <div className="pb-4 text-sm font-semibold text-blue-200 uppercase">Resources</div>
            <ul role="list" className="flex flex-col flex-1">
                <li>
                    <ul role="list" className="-mx-2 space-y-2">
                        {navigationItems.map((item) => {
                            const current = IsSelected(pathname, item.href)
                            return (
                                <li key={item.name}>
                                    <Link
                                        to={item.href}
                                        className={classNames(
                                            current
                                                ? 'bg-blue-700 text-white'
                                                : 'text-blue-200 hover:text-white hover:bg-blue-700',
                                            'group flex items-center gap-x-3 rounded-lg p-3 text-sm font-semibold'
                                        )}
                                    >
                                        <item.icon
                                            className={classNames(
                                                current ? 'text-white' : 'text-blue-200 group-hover:text-white',
                                                'h-5 w-5 shrink-0'
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
    )
}