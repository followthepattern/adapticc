import {
    GiftTopIcon,
    UsersIcon,
} from '@heroicons/react/24/outline'
import { ForwardRefExoticComponent, SVGProps } from 'react'

export interface NavigationItem {
    name: string
    href: string
    icon: ForwardRefExoticComponent<Omit<SVGProps<SVGSVGElement>, "ref"> & {
        title?: string | undefined;
        titleId?: string | undefined;
    } & React.RefAttributes<SVGSVGElement>>
}

export const navigationItems: NavigationItem[] = [
    { name: 'Users', href: '/users', icon: UsersIcon },
    { name: 'Products', href: '/products', icon: GiftTopIcon },
]

export function IsSelected(location: string, path: string): boolean {
    return location.split("/")[1] === path.split("/")[1]
}