import { useState } from "react";
import Header from "./header";
import MobileSidebar from "./mobileSidebar";
import { navigationItems } from "./navbar/navigation";
import StaticSidebar from "./staticSidebar";


export default function HeaderSidebar() {
    const [sidebarOpen, setSidebarOpen] = useState(false);

    return (
        <>
            <MobileSidebar navigationItems={navigationItems} sidebarOpen={sidebarOpen} setSidebarOpen={setSidebarOpen} />
            <StaticSidebar navigationItems={navigationItems} />
            <div className="sticky top-0 z-40 lg:pl-72">
                <Header sidebarOpen={sidebarOpen} setSidebarOpen={setSidebarOpen} />
            </div>
        </>
    )
}