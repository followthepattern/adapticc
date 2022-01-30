import PublicLayout from "../layouts/PublicLayout";

export interface Route {
  path: string;
  exact: boolean;
  public: boolean;
//   Component: (props: any) => JSX.Element;
  Layout: (props: any) => JSX.Element;
  subRoutes?: Route[];
}

export const mainRoutes: Route[] = [
    {
      path: "/",
      exact: true,
      public: true,
    //   Component: Home,
      Layout: PublicLayout,
    },
  ];