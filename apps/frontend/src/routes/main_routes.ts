import AccountLayout from "../layouts/AccountLayout/AccountLayout";
import PublicLayout from "../layouts/PublicLayout";
import Dashboard from "../pages/Dashboard";
import Users from "../pages/Users/Users";

export interface Route {
  path: string;
  exact: boolean;
  public: boolean;
  Page: (props: any) => JSX.Element;
  Layout: (props: any) => JSX.Element;
  title: string;
  icon: string;
  subRoutes?: Route[];
}

export const mainRoutes: Route[] = [
  {
    path: "/",
    exact: true,
    public: true,
    Page: Dashboard,
    title: "Dashboard",
    icon: "D",
    Layout: AccountLayout,
  },
  {
    path: "/users",
    exact: true,
    public: true,
    title: "Users",
    icon: "U",
    Page: Users,
    Layout: AccountLayout,
  },
];
