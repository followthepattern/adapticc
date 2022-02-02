import NavbarItem from "../headless/Navbar/NavbarItem";

export interface Route {
  path: string;
  title: string;
  icon: string;
  subRoutes?: Route[];
}

const Navbar = (items: Route[]) => {
  return (
    <>
      <div className="flex h-[50px] w-full">
        <div className="flex h-[50px] w-full justify-center items-center">
          {" "}
          Navbar - Header
        </div>
        <div className="w-[20px]">+</div>
      </div>
      {items.map((item) => (
        <NavbarItem
          className="flex w-full h-[40px] items-center justify-center items-center pl-8 rounded bg-gray-300"
          path={item.path}
          key={item.path}
        >
          <div className="w-14 items-center">{item.icon}</div>
          <div className="w-full items-center">{item.title}</div>
          </NavbarItem>
      ))}
    </>
  );
};

export default Navbar;
