import NavbarItem from "../../../components/headless/Navbar/NavbarItem";
import HeadlessNavbar from "../../../components/headless/Navbar/Navbar";
import { UserCircleIcon } from "@heroicons/react/outline";

export interface INavbarItem {
  path: string;
  title: string;
  icon?: (props: any) => JSX.Element;
  subRoutes?: INavbarItem[];
  showNavbar?: boolean;
}

const Navbar = (items: INavbarItem[]) => {
  return (
    <HeadlessNavbar className="flex flex-col h-full overflow-y-auto py-4 px-3 bg-gray-50 rounded dark:bg-gray-800">
      <HeadlessNavbar.Header className="h-[50px] w-full self-center text-lg font-semibold whitespace-nowrap dark:text-white">
        <div>Adapticc</div>
      </HeadlessNavbar.Header>
      <div className="grow">
        {items
          .map((item) => (
            <NavbarItem
              className="flex items-center p-2 text-base font-normal text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700 cursor-pointer"
              path={item.path}
              key={item.path}
            >
              <NavbarItem.Icon className="w-14 items-center">
                {item.icon? <item.icon className="h-8 w-8"/> : ""}
              </NavbarItem.Icon>
              <NavbarItem.Title className="w-full ml-3 items-center">
                {item.title}
              </NavbarItem.Title>
            </NavbarItem>
          ))}
      </div>
      <HeadlessNavbar.AccountProfile
        className="flex items-center bottom-0 p-2 text-base font-normal text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700 cursor-pointer"
        accountProfilePath="/users"
        icon={<UserCircleIcon className="h-8 w-8 p-1" />}
      >
        View Profile
      </HeadlessNavbar.AccountProfile>
    </HeadlessNavbar>
  );
};

export default Navbar;
