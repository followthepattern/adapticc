export interface NavbarProperties {
    path: string;
    title: string;
    icon: string;
    subRoutes?: NavbarProperties[];
  }
  
  const Navbar = (items: NavbarProperties[]) => {
    return (
      <>
        <div className="flex h-[50px] w-full">
            <div className="flex h-[50px] w-full justify-center items-center"> Navbar - Header</div>
            <div className="w-[20px]">+</div>
        </div>
        {items.map((item) => (
          <a
            className="flex w-full h-[40px] items-center justify-center items-center pl-8 rounded bg-gray-300"
            href={item.path}
            key={item.path}
          >
            <div className="w-14 items-center">{item.icon}</div>
            <div className="w-full items-center">{item.title}</div>
          </a>
        ))}
      </>
    );
  };
  
  export default Navbar;
  