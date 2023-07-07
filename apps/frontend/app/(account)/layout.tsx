'use client';

import WithGraphQL from "@/components/withGraphQL"
import WithUserContext from "@/components/withUserContext"
import { LOGIN_URL } from "@/lib/constants";
import { useTokenStore } from "@/lib/store"
import useHasMounted from "@/lib/useMounted";
import { redirect, useSelectedLayoutSegments } from "next/navigation";
import { useState } from "react";
import MobileSidebar from "./components/mobileSidebar";
import StaticSidebar from "./components/staticSidebar";
import Header from "./components/header";
import { navigationItems } from "./components/navigation";
import Breadcrumbs from "./components/breadcrumbs";

interface AccountLayoutProperties {
  children: React.ReactNode,
}

export default function AccountLayout({ children }: AccountLayoutProperties) {
  const [sidebarOpen, setSidebarOpen] = useState(false)

  const segments = useSelectedLayoutSegments()

  const { token } = useTokenStore();

  const hasMounted = useHasMounted();

  if (!hasMounted) {
    return null;
  }

  if (token == "") {
    redirect(LOGIN_URL);
  }

  return (
    <WithGraphQL token={token}>
      <WithUserContext>
        <MobileSidebar navigationItems={navigationItems} sidebarOpen={sidebarOpen} setSidebarOpen={setSidebarOpen} />
        <StaticSidebar navigationItems={navigationItems} />
        <div className="lg:pl-72">
          <Header setSidebarOpen={setSidebarOpen} />
          <Breadcrumbs pages={segments.map(segment => ({name: segment, href: segment}))}/>
          <main className="py-10">
            <div className="px-4 sm:px-6 lg:px-8">{children}</div>
          </main>
        </div>
      </WithUserContext>
    </WithGraphQL>
  )
}
