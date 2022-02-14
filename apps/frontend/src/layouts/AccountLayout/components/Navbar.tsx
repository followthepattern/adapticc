import NavbarItem from "../../../components/headless/Navbar/NavbarItem";
import HeadlessNavbar from "../../../components/headless/Navbar/Navbar";
import { UserCircleIcon } from "@heroicons/react/outline";

export interface Route {
  path: string;
  title: string;
  icon: string;
  subRoutes?: Route[];
  showNavbar?: boolean;
}

const Navbar = (items: Route[]) => {
  return (
    <HeadlessNavbar className="h-full flex-col">
      <HeadlessNavbar.Header className="flex h-[50px] w-full justify-center items-center">
        <div>Navbar title</div>
      </HeadlessNavbar.Header>
      {items
        .filter((item) => item.showNavbar)
        .map((item) => (
          <NavbarItem
            className="flex w-full h-[40px] items-center justify-center items-center pl-8 rounded bg-gray-300"
            path={item.path}
            key={item.path}
          >
            <NavbarItem.Icon className="w-14 items-center">
              {item.icon}
            </NavbarItem.Icon>
            <NavbarItem.Title className="w-full items-center">
              {item.title}
            </NavbarItem.Title>
          </NavbarItem>
        ))}
      <HeadlessNavbar.AccountProfile
        className="flex absolute bottom-0 w-full"
        accountProfilePath="/users"
        icon={<UserCircleIcon className="h-5 w-5" />}
      >
        View Profile
      </HeadlessNavbar.AccountProfile>
    </HeadlessNavbar>
  );
};

export default Navbar;
