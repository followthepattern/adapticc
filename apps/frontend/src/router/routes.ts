import AccountLayout from "../layouts/AccountLayout/AccountLayout";
import LoginLayout from "../layouts/LoginLayout/LoginLayout";
// import PublicLayout from "../layouts/PublicLayout";
import Dashboard from "../modules/dashboard/views/Dashboard";
import Login from "../modules/login/views/Login";
import Users from "../modules/user/views/UserListPage";
import { ViewBoardsIcon, UsersIcon } from "@heroicons/react/outline";

export interface Route {
  path: string;
  exact: boolean;
  public: boolean;
  showNavbar?: boolean;
  Page: (props: any) => JSX.Element;
  Layout: (props: any) => JSX.Element;
  title: string;
  icon?: (props: any) => JSX.Element;
  subRoutes?: Route[];
}

export const routes: Route[] = [
  {
    path: "/dashboard",
    exact: true,
    public: false,
    Page: Dashboard,
    title: "Dashboard",
    icon: ViewBoardsIcon,
    showNavbar: true,
    Layout: AccountLayout,
  },
  {
    path: "/users",
    exact: true,
    public: false,
    showNavbar: true,
    title: "Users",
    icon: UsersIcon,
    Page: Users,
    Layout: AccountLayout,
  },
  {
    path: "/login",
    exact: true,
    public: true,
    title: "Login",
    Page: Login,
    Layout: LoginLayout,
  },
];
