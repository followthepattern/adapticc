const AccountLayout = (props: any) => {
  return (
    <>
      <div className="h-full flex">
        <div className="flex-none w-[250px] h-full bg-yellow-200">navbar</div>
        <div className="grow h-full">
          <div className="w-full h-[50px] bg-red-200">header</div>
          <div className="grow w-full h-full bg-violet-300">content</div>
        </div>
      </div>
    </>
  );
};

export default AccountLayout;
