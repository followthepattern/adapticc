import WithGraphQL from "@/components/withGraphQL"
import WithUserContext from "@/components/withUserContext"
import { useTokenStore } from "@/lib/store"
import { Outlet } from "react-router-dom";
import LoginExpired from "@/app/(auth)/login/components/loginLoggedOut";
import HeaderSidebar from "./components/headerSidebar/headerSidebar";

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
        <HeaderSidebar />
        <div className="lg:pl-72">
          <main className="py-10">
            <div className="px-1 xs:px-2 sm:px-6 lg:px-8">
              <Outlet />
            </div>
          </main>
        </div>
      </WithUserContext>
    </WithGraphQL>
  )
}
