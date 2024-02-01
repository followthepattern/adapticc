import classNames from 'classnames';
import { IsSelected, NavigationItem } from './navbar/navigation';
import {Link, useLocation} from 'react-router-dom';
import Navbar from './navbar/navbar';

interface StaticSidebarPorperties {
    navigationItems: NavigationItem[]
}

export default function StaticSidebar(props: StaticSidebarPorperties) {
    const pathname = useLocation().pathname;

    return (
        <div className="hidden lg:fixed lg:inset-y-0 lg:z-50 lg:flex lg:w-72 lg:flex-col">
            <div className="flex flex-col px-5 pb-4 overflow-y-auto bg-blue-600 grow">
                <div className="flex items-center h-16 shrink-0"></div>
                <Navbar />
            </div>
        </div>
    )
}