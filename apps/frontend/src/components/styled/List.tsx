import HeadlessList from "../headless/List/List";
import { PaginationData } from "../headless/Pagination/Pagination";

export interface ListProperties {
  columns: string[];
  rows: Row[];
  paginationData: PaginationData;
  resourceName: string;
}

interface Row {
  cells: string[];
  action: string;
}

const List = ({
  columns,
  rows,
  paginationData,
  resourceName,
}: ListProperties) => {
  const Header = HeadlessList.Table.Header;
  const Body = HeadlessList.Table.Body;
  const TR = HeadlessList.Table.Header.TR;
  const TH = HeadlessList.Table.Header.TR.TH;
  const TD = HeadlessList.Table.Header.TR.TD;
  const ActionButton = HeadlessList.ActionButton;

  return (
    <HeadlessList>
      <div className="px-1">
        <HeadlessList.Filters
          className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-80 pl-10 p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
          placeholder="Search for items"
        ></HeadlessList.Filters>
      </div>
      <div className="py-2 sm:rounded-lg overflow-x-auto">
        <HeadlessList.Table className="min-w-full">
          <Header className="bg-gray-100 dark:bg-gray-700">
            <TR key="0">
              {columns.map((column, index) => (
                <TH
                  key={index}
                  className="py-3 px-6 text-xs font-medium tracking-wider text-left text-gray-700 uppercase dark:text-gray-400"
                >
                  {column}
                </TH>
              ))}
            </TR>
          </Header>
          <Body>
            {rows.map((row, index) => {
              console.info(index);
              return (
                <TR
                  trKey={index}
                  className="border-b odd:bg-white even:bg-gray-50 odd:dark:bg-gray-800 even:dark:bg-gray-700 dark:border-gray-600"
                >
                  {row.cells.map((cell, rowIndex) => {
                    return (
                      <TD
                        key={rowIndex}
                        className="py-4 px-6 text-sm font-medium text-gray-900 whitespace-nowrap dark:text-white"
                      >
                        {cell}
                      </TD>
                    );
                  })}
                  <TD className="py-4 px-6 text-sm font-medium text-right whitespace-nowrap">
                    <ActionButton className="font-medium text-blue-600 dark:text-blue-500 hover:underline" action={() => alert(row)}>
                      Edit
                    </ActionButton>
                  </TD>
                </TR>
              );
            })}
          </Body>
        </HeadlessList.Table>
      </div>
      <HeadlessList.Pagination
        renderUrl={(page: number) => {
          return `./${resourceName}/${page}`;
        }}
        paginationData={paginationData}
        numberButtonClassName="bg-white border-gray-300 text-gray-500 hover:bg-gray-50 hidden md:inline-flex items-center px-4 py-2 border text-sm font-medium"
        actionButtonClassName="bg-white border-gray-300 text-gray-500 hover:bg-gray-50 md:inline-flex items-center px-4 py-2 border text-sm font-medium"
        currentButtonClassName="bg-indigo-50 border-indigo-500 text-indigo-600 md:inline-flex items-center px-4 py-2 border text-sm font-medium"
        className="flex py-2 -space-x-px justify-center"
      />
    </HeadlessList>
  );
};

export default List;
