'use client';

import WithGraphQL from "@/components/withGraphQL"
import WithUserContext from "@/components/withUserContext"
import { useTokenStore } from "@/lib/store"
import { useState } from "react";
import MobileSidebar from "./components/mobileSidebar";
import StaticSidebar from "./components/staticSidebar";
import Header from "./components/header";
import { navigationItems } from "./components/navigation";
import { Outlet } from "react-router-dom";
import LoginExpired from "@/app/(auth)/login/components/loginLoggedOut";

export default function AccountLayout() {
  // const segments = [];

  const { token } = useTokenStore();

  if (token == "") {
    return (
      <WithGraphQL>
        <LoginExpired />
      </WithGraphQL>
    )
  }

  return (
    <WithGraphQL token={token}>
      <WithUserContext>
        <MobileSidebar navigationItems={navigationItems} />
        <StaticSidebar navigationItems={navigationItems} />
        <div className="lg:pl-72">
          <Header />
          {/* <Breadcrumbs pages={segments.map(segment => ({name: segment, href: segment}))}/> */}
          <main className="py-10">
            <div className="px-4 sm:px-6 lg:px-8">
              <Outlet />
            </div>
          </main>
        </div>
      </WithUserContext>
    </WithGraphQL>
  )
}
