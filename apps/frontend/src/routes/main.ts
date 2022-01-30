import PublicLayout from "../layouts/PublicLayout";
import Dashboard from "../pages/Dashboard";
import Users from "../pages/Users";

export interface Route {
  path: string;
  exact: boolean;
  public: boolean;
  Page: (props: any) => JSX.Element;
  Layout: (props: any) => JSX.Element;
  subRoutes?: Route[];
}

export const mainRoutes: Route[] = [
    {
      path: "/",
      exact: true,
      public: true,
      Page: Dashboard,
      Layout: PublicLayout,
    },
    {
        path: "/users",
        exact: true,
        public: true,
        Page: Users,
        Layout: PublicLayout,
      },
  ];