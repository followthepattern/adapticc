const Users = (props: any) => {
  return (
    <>
      <div className="inline-block text-3xl px-1 py-4">Users</div>
      <div className="px-1 py-4">
        <label htmlFor="table-search" className="sr-only">
          Search
        </label>
        <div className="">
          <input
            type="text"
            id="table-search"
            className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-80 pl-10 p-2.5  dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
            placeholder="Search for items"
          />
        </div>
      </div>
      <div className="pb-6 sm:rounded-lg overflow-x-auto">
        <table className="min-w-full">
          <thead className="bg-gray-100 dark:bg-gray-700">
            <tr>
              <th
                scope="col"
                className="py-3 px-6 text-xs font-medium tracking-wider text-left text-gray-700 uppercase dark:text-gray-400"
              >
                Name
              </th>
              <th
                scope="col"
                className="py-3 px-6 text-xs font-medium tracking-wider text-left text-gray-700 uppercase dark:text-gray-400"
              >
                Color
              </th>
              <th
                scope="col"
                className="py-3 px-6 text-xs font-medium tracking-wider text-left text-gray-700 uppercase dark:text-gray-400"
              >
                Category
              </th>
              <th
                scope="col"
                className="py-3 px-6 text-xs font-medium tracking-wider text-left text-gray-700 uppercase dark:text-gray-400"
              >
                Price
              </th>
              <th scope="col" className="py-3 px-6"></th>
            </tr>
          </thead>
          <tbody>
            <tr className="border-b odd:bg-white even:bg-gray-50 odd:dark:bg-gray-800 even:dark:bg-gray-700 dark:border-gray-600">
              <td className="py-4 px-6 text-sm font-medium text-gray-900 whitespace-nowrap dark:text-white">
                Apple MacBook Pro 17"
              </td>
              <td className="py-4 px-6 text-sm text-gray-500 whitespace-nowrap dark:text-gray-400">
                Sliver
              </td>
              <td className="py-4 px-6 text-sm text-gray-500 whitespace-nowrap dark:text-gray-400">
                Laptop
              </td>
              <td className="py-4 px-6 text-sm text-gray-500 whitespace-nowrap dark:text-gray-400">
                $2999
              </td>
              <td className="py-4 px-6 text-sm font-medium text-right whitespace-nowrap">
                <a
                  href="#"
                  className="text-blue-600 dark:text-blue-500 hover:underline"
                >
                  Edit
                </a>
              </td>
            </tr>
            <tr className="border-b odd:bg-white even:bg-gray-50 odd:dark:bg-gray-800 even:dark:bg-gray-700 dark:border-gray-600">
              <td className="py-4 px-6 text-sm font-medium text-gray-900 whitespace-nowrap dark:text-white">
                Apple MacBook Pro 17"
              </td>
              <td className="py-4 px-6 text-sm text-gray-500 whitespace-nowrap dark:text-gray-400">
                Sliver
              </td>
              <td className="py-4 px-6 text-sm text-gray-500 whitespace-nowrap dark:text-gray-400">
                Laptop
              </td>
              <td className="py-4 px-6 text-sm text-gray-500 whitespace-nowrap dark:text-gray-400">
                $2999
              </td>
              <td className="py-4 px-6 text-sm font-medium text-right whitespace-nowrap">
                <a
                  href="#"
                  className="text-blue-600 dark:text-blue-500 hover:underline"
                >
                  Edit
                </a>
              </td>
            </tr>
            <tr className="border-b odd:bg-white even:bg-gray-50 odd:dark:bg-gray-800 even:dark:bg-gray-700 dark:border-gray-600">
              <td className="py-4 px-6 text-sm font-medium text-gray-900 whitespace-nowrap dark:text-white">
                Apple MacBook Pro 17"
              </td>
              <td className="py-4 px-6 text-sm text-gray-500 whitespace-nowrap dark:text-gray-400">
                Sliver
              </td>
              <td className="py-4 px-6 text-sm text-gray-500 whitespace-nowrap dark:text-gray-400">
                Laptop
              </td>
              <td className="py-4 px-6 text-sm text-gray-500 whitespace-nowrap dark:text-gray-400">
                $2999
              </td>
              <td className="py-4 px-6 text-sm font-medium text-right whitespace-nowrap">
                <a
                  href="#"
                  className="text-blue-600 dark:text-blue-500 hover:underline"
                >
                  Edit
                </a>
              </td>
            </tr>{" "}
            <tr className="border-b odd:bg-white even:bg-gray-50 odd:dark:bg-gray-800 even:dark:bg-gray-700 dark:border-gray-600">
              <td className="py-4 px-6 text-sm font-medium text-gray-900 whitespace-nowrap dark:text-white">
                Apple MacBook Pro 17"
              </td>
              <td className="py-4 px-6 text-sm text-gray-500 whitespace-nowrap dark:text-gray-400">
                Sliver
              </td>
              <td className="py-4 px-6 text-sm text-gray-500 whitespace-nowrap dark:text-gray-400">
                Laptop
              </td>
              <td className="py-4 px-6 text-sm text-gray-500 whitespace-nowrap dark:text-gray-400">
                $2999
              </td>
              <td className="py-4 px-6 text-sm font-medium text-right whitespace-nowrap">
                <a
                  href="#"
                  className="text-blue-600 dark:text-blue-500 hover:underline"
                >
                  Edit
                </a>
              </td>
            </tr>{" "}
            <tr className="border-b odd:bg-white even:bg-gray-50 odd:dark:bg-gray-800 even:dark:bg-gray-700 dark:border-gray-600">
              <td className="py-4 px-6 text-sm font-medium text-gray-900 whitespace-nowrap dark:text-white">
                Apple MacBook Pro 17"
              </td>
              <td className="py-4 px-6 text-sm text-gray-500 whitespace-nowrap dark:text-gray-400">
                Sliver
              </td>
              <td className="py-4 px-6 text-sm text-gray-500 whitespace-nowrap dark:text-gray-400">
                Laptop
              </td>
              <td className="py-4 px-6 text-sm text-gray-500 whitespace-nowrap dark:text-gray-400">
                $2999
              </td>
              <td className="py-4 px-6 text-sm font-medium text-right whitespace-nowrap">
                <a
                  href="#"
                  className="text-blue-600 dark:text-blue-500 hover:underline"
                >
                  Edit
                </a>
              </td>
            </tr>{" "}
            <tr className="border-b odd:bg-white even:bg-gray-50 odd:dark:bg-gray-800 even:dark:bg-gray-700 dark:border-gray-600">
              <td className="py-4 px-6 text-sm font-medium text-gray-900 whitespace-nowrap dark:text-white">
                Apple MacBook Pro 17"
              </td>
              <td className="py-4 px-6 text-sm text-gray-500 whitespace-nowrap dark:text-gray-400">
                Sliver
              </td>
              <td className="py-4 px-6 text-sm text-gray-500 whitespace-nowrap dark:text-gray-400">
                Laptop
              </td>
              <td className="py-4 px-6 text-sm text-gray-500 whitespace-nowrap dark:text-gray-400">
                $2999
              </td>
              <td className="py-4 px-6 text-sm font-medium text-right whitespace-nowrap">
                <a
                  href="#"
                  className="text-blue-600 dark:text-blue-500 hover:underline"
                >
                  Edit
                </a>
              </td>
            </tr>{" "}
            <tr className="border-b odd:bg-white even:bg-gray-50 odd:dark:bg-gray-800 even:dark:bg-gray-700 dark:border-gray-600">
              <td className="py-4 px-6 text-sm font-medium text-gray-900 whitespace-nowrap dark:text-white">
                Apple MacBook Pro 17"
              </td>
              <td className="py-4 px-6 text-sm text-gray-500 whitespace-nowrap dark:text-gray-400">
                Sliver
              </td>
              <td className="py-4 px-6 text-sm text-gray-500 whitespace-nowrap dark:text-gray-400">
                Laptop
              </td>
              <td className="py-4 px-6 text-sm text-gray-500 whitespace-nowrap dark:text-gray-400">
                $2999
              </td>
              <td className="py-4 px-6 text-sm font-medium text-right whitespace-nowrap">
                <a
                  href="#"
                  className="text-blue-600 dark:text-blue-500 hover:underline"
                >
                  Edit
                </a>
              </td>
            </tr>{" "}
            <tr className="border-b odd:bg-white even:bg-gray-50 odd:dark:bg-gray-800 even:dark:bg-gray-700 dark:border-gray-600">
              <td className="py-4 px-6 text-sm font-medium text-gray-900 whitespace-nowrap dark:text-white">
                Apple MacBook Pro 17"
              </td>
              <td className="py-4 px-6 text-sm text-gray-500 whitespace-nowrap dark:text-gray-400">
                Sliver
              </td>
              <td className="py-4 px-6 text-sm text-gray-500 whitespace-nowrap dark:text-gray-400">
                Laptop
              </td>
              <td className="py-4 px-6 text-sm text-gray-500 whitespace-nowrap dark:text-gray-400">
                $2999
              </td>
              <td className="py-4 px-6 text-sm font-medium text-right whitespace-nowrap">
                <a
                  href="#"
                  className="text-blue-600 dark:text-blue-500 hover:underline"
                >
                  Edit
                </a>
              </td>
            </tr>{" "}
            <tr className="border-b odd:bg-white even:bg-gray-50 odd:dark:bg-gray-800 even:dark:bg-gray-700 dark:border-gray-600">
              <td className="py-4 px-6 text-sm font-medium text-gray-900 whitespace-nowrap dark:text-white">
                Apple MacBook Pro 17"
              </td>
              <td className="py-4 px-6 text-sm text-gray-500 whitespace-nowrap dark:text-gray-400">
                Sliver
              </td>
              <td className="py-4 px-6 text-sm text-gray-500 whitespace-nowrap dark:text-gray-400">
                Laptop
              </td>
              <td className="py-4 px-6 text-sm text-gray-500 whitespace-nowrap dark:text-gray-400">
                $2999
              </td>
              <td className="py-4 px-6 text-sm font-medium text-right whitespace-nowrap">
                <a
                  href="#"
                  className="text-blue-600 dark:text-blue-500 hover:underline"
                >
                  Edit
                </a>
              </td>
            </tr>{" "}
            <tr className="border-b odd:bg-white even:bg-gray-50 odd:dark:bg-gray-800 even:dark:bg-gray-700 dark:border-gray-600">
              <td className="py-4 px-6 text-sm font-medium text-gray-900 whitespace-nowrap dark:text-white">
                Apple MacBook Pro 17"
              </td>
              <td className="py-4 px-6 text-sm text-gray-500 whitespace-nowrap dark:text-gray-400">
                Sliver
              </td>
              <td className="py-4 px-6 text-sm text-gray-500 whitespace-nowrap dark:text-gray-400">
                Laptop
              </td>
              <td className="py-4 px-6 text-sm text-gray-500 whitespace-nowrap dark:text-gray-400">
                $2999
              </td>
              <td className="py-4 px-6 text-sm font-medium text-right whitespace-nowrap">
                <a
                  href="#"
                  className="text-blue-600 dark:text-blue-500 hover:underline"
                >
                  Edit
                </a>
              </td>
            </tr>{" "}
            <tr className="border-b odd:bg-white even:bg-gray-50 odd:dark:bg-gray-800 even:dark:bg-gray-700 dark:border-gray-600">
              <td className="py-4 px-6 text-sm font-medium text-gray-900 whitespace-nowrap dark:text-white">
                Apple MacBook Pro 17"
              </td>
              <td className="py-4 px-6 text-sm text-gray-500 whitespace-nowrap dark:text-gray-400">
                Sliver
              </td>
              <td className="py-4 px-6 text-sm text-gray-500 whitespace-nowrap dark:text-gray-400">
                Laptop
              </td>
              <td className="py-4 px-6 text-sm text-gray-500 whitespace-nowrap dark:text-gray-400">
                $2999
              </td>
              <td className="py-4 px-6 text-sm font-medium text-right whitespace-nowrap">
                <a
                  href="#"
                  className="text-blue-600 dark:text-blue-500 hover:underline"
                >
                  Edit
                </a>
              </td>
            </tr>{" "}
            <tr className="border-b odd:bg-white even:bg-gray-50 odd:dark:bg-gray-800 even:dark:bg-gray-700 dark:border-gray-600">
              <td className="py-4 px-6 text-sm font-medium text-gray-900 whitespace-nowrap dark:text-white">
                Apple MacBook Pro 17"
              </td>
              <td className="py-4 px-6 text-sm text-gray-500 whitespace-nowrap dark:text-gray-400">
                Sliver
              </td>
              <td className="py-4 px-6 text-sm text-gray-500 whitespace-nowrap dark:text-gray-400">
                Laptop
              </td>
              <td className="py-4 px-6 text-sm text-gray-500 whitespace-nowrap dark:text-gray-400">
                $2999
              </td>
              <td className="py-4 px-6 text-sm font-medium text-right whitespace-nowrap">
                <a
                  href="#"
                  className="text-blue-600 dark:text-blue-500 hover:underline"
                >
                  Edit
                </a>
              </td>
            </tr>{" "}
            <tr className="border-b odd:bg-white even:bg-gray-50 odd:dark:bg-gray-800 even:dark:bg-gray-700 dark:border-gray-600">
              <td className="py-4 px-6 text-sm font-medium text-gray-900 whitespace-nowrap dark:text-white">
                Apple MacBook Pro 17"
              </td>
              <td className="py-4 px-6 text-sm text-gray-500 whitespace-nowrap dark:text-gray-400">
                Sliver
              </td>
              <td className="py-4 px-6 text-sm text-gray-500 whitespace-nowrap dark:text-gray-400">
                Laptop
              </td>
              <td className="py-4 px-6 text-sm text-gray-500 whitespace-nowrap dark:text-gray-400">
                $2999
              </td>
              <td className="py-4 px-6 text-sm font-medium text-right whitespace-nowrap">
                <a
                  href="#"
                  className="text-blue-600 dark:text-blue-500 hover:underline"
                >
                  Edit
                </a>
              </td>
            </tr>{" "}
            <tr className="border-b odd:bg-white even:bg-gray-50 odd:dark:bg-gray-800 even:dark:bg-gray-700 dark:border-gray-600">
              <td className="py-4 px-6 text-sm font-medium text-gray-900 whitespace-nowrap dark:text-white">
                Apple MacBook Pro 17"
              </td>
              <td className="py-4 px-6 text-sm text-gray-500 whitespace-nowrap dark:text-gray-400">
                Sliver
              </td>
              <td className="py-4 px-6 text-sm text-gray-500 whitespace-nowrap dark:text-gray-400">
                Laptop
              </td>
              <td className="py-4 px-6 text-sm text-gray-500 whitespace-nowrap dark:text-gray-400">
                $2999
              </td>
              <td className="py-4 px-6 text-sm font-medium text-right whitespace-nowrap">
                <a
                  href="#"
                  className="text-blue-600 dark:text-blue-500 hover:underline"
                >
                  Edit
                </a>
              </td>
            </tr>{" "}
            <tr className="border-b odd:bg-white even:bg-gray-50 odd:dark:bg-gray-800 even:dark:bg-gray-700 dark:border-gray-600">
              <td className="py-4 px-6 text-sm font-medium text-gray-900 whitespace-nowrap dark:text-white">
                Apple MacBook Pro 17"
              </td>
              <td className="py-4 px-6 text-sm text-gray-500 whitespace-nowrap dark:text-gray-400">
                Sliver
              </td>
              <td className="py-4 px-6 text-sm text-gray-500 whitespace-nowrap dark:text-gray-400">
                Laptop
              </td>
              <td className="py-4 px-6 text-sm text-gray-500 whitespace-nowrap dark:text-gray-400">
                $2999
              </td>
              <td className="py-4 px-6 text-sm font-medium text-right whitespace-nowrap">
                <a
                  href="#"
                  className="text-blue-600 dark:text-blue-500 hover:underline"
                >
                  Edit
                </a>
              </td>
            </tr>{" "}
            <tr className="border-b odd:bg-white even:bg-gray-50 odd:dark:bg-gray-800 even:dark:bg-gray-700 dark:border-gray-600">
              <td className="py-4 px-6 text-sm font-medium text-gray-900 whitespace-nowrap dark:text-white">
                Apple MacBook Pro 17"
              </td>
              <td className="py-4 px-6 text-sm text-gray-500 whitespace-nowrap dark:text-gray-400">
                Sliver
              </td>
              <td className="py-4 px-6 text-sm text-gray-500 whitespace-nowrap dark:text-gray-400">
                Laptop
              </td>
              <td className="py-4 px-6 text-sm text-gray-500 whitespace-nowrap dark:text-gray-400">
                $2999
              </td>
              <td className="py-4 px-6 text-sm font-medium text-right whitespace-nowrap">
                <a
                  href="#"
                  className="text-blue-600 dark:text-blue-500 hover:underline"
                >
                  Edit
                </a>
              </td>
            </tr>{" "}
            <tr className="border-b odd:bg-white even:bg-gray-50 odd:dark:bg-gray-800 even:dark:bg-gray-700 dark:border-gray-600">
              <td className="py-4 px-6 text-sm font-medium text-gray-900 whitespace-nowrap dark:text-white">
                Apple MacBook Pro 17"
              </td>
              <td className="py-4 px-6 text-sm text-gray-500 whitespace-nowrap dark:text-gray-400">
                Sliver
              </td>
              <td className="py-4 px-6 text-sm text-gray-500 whitespace-nowrap dark:text-gray-400">
                Laptop
              </td>
              <td className="py-4 px-6 text-sm text-gray-500 whitespace-nowrap dark:text-gray-400">
                $2999
              </td>
              <td className="py-4 px-6 text-sm font-medium text-right whitespace-nowrap">
                <a
                  href="#"
                  className="text-blue-600 dark:text-blue-500 hover:underline"
                >
                  Edit
                </a>
              </td>
            </tr>
            <tr className="border-b odd:bg-white even:bg-gray-50 odd:dark:bg-gray-800 even:dark:bg-gray-700 dark:border-gray-600">
              <td className="py-4 px-6 text-sm font-medium text-gray-900 whitespace-nowrap dark:text-white">
                Apple Imac 27"
              </td>
              <td className="py-4 px-6 text-sm text-gray-500 whitespace-nowrap dark:text-gray-400">
                White
              </td>
              <td className="py-4 px-6 text-sm text-gray-500 whitespace-nowrap dark:text-gray-400">
                Desktop Pc
              </td>
              <td className="py-4 px-6 text-sm text-gray-500 whitespace-nowrap dark:text-gray-400">
                $1999
              </td>
              <td className="py-4 px-6 text-sm font-medium text-right whitespace-nowrap">
                <a
                  href="#"
                  className="text-blue-600 dark:text-blue-500 hover:underline"
                >
                  Edit
                </a>
              </td>
            </tr>
            <tr className="border-b odd:bg-white even:bg-gray-50 odd:dark:bg-gray-800 even:dark:bg-gray-700 dark:border-gray-600">
              <td className="py-4 px-6 text-sm font-medium text-gray-900 whitespace-nowrap dark:text-white">
                iPhone 13 Pro
              </td>
              <td className="py-4 px-6 text-sm text-gray-500 whitespace-nowrap dark:text-gray-400">
                White
              </td>
              <td className="py-4 px-6 text-sm text-gray-500 whitespace-nowrap dark:text-gray-400">
                Phone
              </td>
              <td className="py-4 px-6 text-sm text-gray-500 whitespace-nowrap dark:text-gray-400">
                $999
              </td>
              <td className="py-4 px-6 text-sm font-medium text-right whitespace-nowrap">
                <a
                  href="#"
                  className="text-blue-600 dark:text-blue-500 hover:underline"
                >
                  Edit
                </a>
              </td>
            </tr>
            <tr className="border-b odd:bg-white even:bg-gray-50 odd:dark:bg-gray-800 even:dark:bg-gray-700 dark:border-gray-600">
              <td className="py-4 px-6 text-sm font-medium text-gray-900 whitespace-nowrap dark:text-white">
                Apple Magic Mouse 2
              </td>
              <td className="py-4 px-6 text-sm text-gray-500 whitespace-nowrap dark:text-gray-400">
                White
              </td>
              <td className="py-4 px-6 text-sm text-gray-500 whitespace-nowrap dark:text-gray-400">
                Accessories
              </td>
              <td className="py-4 px-6 text-sm text-gray-500 whitespace-nowrap dark:text-gray-400">
                $99
              </td>
              <td className="py-4 px-6 text-sm font-medium text-right whitespace-nowrap">
                <a
                  href="#"
                  className="text-blue-600 dark:text-blue-500 hover:underline"
                >
                  Edit
                </a>
              </td>
            </tr>
            <tr className="border-b odd:bg-white even:bg-gray-50 odd:dark:bg-gray-800 even:dark:bg-gray-700 dark:border-gray-600">
              <td className="py-4 px-6 text-sm font-medium text-gray-900 whitespace-nowrap dark:text-white">
                Apple Watch Series 7
              </td>
              <td className="py-4 px-6 text-sm text-gray-500 whitespace-nowrap dark:text-gray-400">
                Pink
              </td>
              <td className="py-4 px-6 text-sm text-gray-500 whitespace-nowrap dark:text-gray-400">
                Accessories
              </td>
              <td className="py-4 px-6 text-sm text-gray-500 whitespace-nowrap dark:text-gray-400">
                $599
              </td>
              <td className="py-4 px-6 text-sm font-medium text-right whitespace-nowrap">
                <a
                  href="#"
                  className="text-blue-600 dark:text-blue-500 hover:underline"
                >
                  Edit
                </a>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
      <nav className="flex py-2 -space-x-px justify-center">
        <a
          href="#"
          className="items-center px-2 py-2 rounded-l-md border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-50"
        >
          <svg
            className="h-5 w-5"
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 20 20"
            fill="currentColor"
            aria-hidden="true"
          >
            <path
              fillRule="evenodd"
              d="M12.707 5.293a1 1 0 010 1.414L9.414 10l3.293 3.293a1 1 0 01-1.414 1.414l-4-4a1 1 0 010-1.414l4-4a1 1 0 011.414 0z"
              clipRule="evenodd"
            />
          </svg>
        </a>
        <a
          href="#"
          aria-current="page"
          className="bg-indigo-50 border-indigo-500 text-indigo-600 inline-flex items-center px-4 py-2 border text-sm font-medium"
        >
          {" "}
          1{" "}
        </a>
        <a
          href="#"
          className="bg-white border-gray-300 text-gray-500 hover:bg-gray-50 inline-flex items-center px-4 py-2 border text-sm font-medium"
        >
          {" "}
          2{" "}
        </a>
        <a
          href="#"
          className="bg-white border-gray-300 text-gray-500 hover:bg-gray-50 hidden md:inline-flex items-center px-4 py-2 border text-sm font-medium"
        >
          {" "}
          3{" "}
        </a>
        <a
          href="#"
          className="bg-white border-gray-300 text-gray-500 hover:bg-gray-50 hidden md:inline-flex items-center px-4 py-2 border text-sm font-medium"
        >
          {" "}
          8{" "}
        </a>
        <a
          href="#"
          className="bg-white border-gray-300 text-gray-500 hover:bg-gray-50 inline-flex items-center px-4 py-2 border text-sm font-medium"
        >
          {" "}
          9{" "}
        </a>
        <a
          href="#"
          className="bg-white border-gray-300 text-gray-500 hover:bg-gray-50 inline-flex items-center px-4 py-2 border text-sm font-medium"
        >
          {" "}
          10{" "}
        </a>
        <a
          href="#"
          className="inline-flex items-center px-2 py-2 rounded-r-md border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-50"
        >
          <svg
            className="h-5 w-5"
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 20 20"
            fill="currentColor"
            aria-hidden="true"
          >
            <path
              fillRule="evenodd"
              d="M7.293 14.707a1 1 0 010-1.414L10.586 10 7.293 6.707a1 1 0 011.414-1.414l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0z"
              clipRule="evenodd"
            />
          </svg>
        </a>
      </nav>
    </>
  );
};

export default Users;
