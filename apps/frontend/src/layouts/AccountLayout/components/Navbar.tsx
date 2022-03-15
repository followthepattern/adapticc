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
    <HeadlessNavbar className="relative h-full flex-col overflow-y-auto py-4 px-3 bg-gray-50 rounded dark:bg-gray-800">
        <HeadlessNavbar.Header className="flex h-[50px] w-full self-center text-lg font-semibold whitespace-nowrap dark:text-white">
          <div>Adapticc</div>
        </HeadlessNavbar.Header>
        {items
          .filter((item) => item.showNavbar)
          .map((item) => (
            <NavbarItem
              className="flex items-center p-2 text-base font-normal text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700 cursor-pointer"
              path={item.path}
              key={item.path}
            >
              <NavbarItem.Icon className="w-14 items-center">
                <svg
                  className="w-6 h-6 text-gray-500 transition duration-75 dark:text-gray-400 group-hover:text-gray-900 dark:group-hover:text-white"
                  fill="currentColor"
                  viewBox="0 0 20 20"
                  xmlns="http://www.w3.org/2000/svg"
                >
                  <path d="M2 10a8 8 0 018-8v8h8a8 8 0 11-16 0z"></path>
                  <path d="M12 2.252A8.014 8.014 0 0117.748 8H12V2.252z"></path>
                </svg>
              </NavbarItem.Icon>
              <NavbarItem.Title className="w-full ml-3 items-center">
                {item.title}
              </NavbarItem.Title>
            </NavbarItem>
          ))}
        <HeadlessNavbar.AccountProfile
          className="flex items-center absolute bottom-0 p-2 text-base font-normal text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700 cursor-pointer"
          accountProfilePath="/users"
          icon={<UserCircleIcon className="h-7 w-7 p-1" />}
        >
          View Profile
        </HeadlessNavbar.AccountProfile>
    </HeadlessNavbar>
  );
};

export default Navbar;
