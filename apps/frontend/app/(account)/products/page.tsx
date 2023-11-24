import SectionHeader from '../components/listPage/sectionHeader/sectionHeader';
import List from '../components/listPage/listPage';
import { ListPageComponentProperties } from '../components/listPage/listPageWrapper/listPageWrapper';
import { SortLabel, SetPageParams, SetSearchPatternParams, SetSortPatternParrams } from '../components/listPage/listPageWrapper/listingFunctions';
import useListProduct from './hooks/listProduct';
import ProductTable from './hooks/productTable';

export const RESOURCE_NAME = "Products"
export const RESOURCE_URL = "/products"

const sortByLables: SortLabel[] = [
  {
    code: "id",
    name: "ID",
    asc: true,
  },
  {
    code: "title",
    name: "Title",
    asc: true,
  }
];

export default function Products(props: ListPageComponentProperties) {
  const sortOnChange = (sortLabel: SortLabel) => {
    SetSortPatternParrams(props.searchParams, props.setSearchParams, sortLabel);
  }

  const searchInputOnChange = (searchString: string) => {
    SetSearchPatternParams(props.searchParams, props.setSearchParams, searchString);
  }

  const pageOnChange = (page: number) => {
    SetPageParams(props.searchParams, props.setSearchParams, page);
  }

  const selectedSortLabel = sortByLables.find(l => l.code == props.sortProps.sortLabel?.code);

  return (
    <div>
      <SectionHeader
        resourceName={RESOURCE_NAME}
        resourceUrl={RESOURCE_URL}
        searchInputOnChange={searchInputOnChange}
        sortOnChange={sortOnChange}
        searchInput={props.filterProps.searchString}
        sortByLables={sortByLables}
        selectedSortLabel={selectedSortLabel}
      />
      <div className="mt-8 flow-root overflow-hidden">
        <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
          <List {...props}
            sortProps={{ sortLabel: selectedSortLabel }}
            onPageChange={pageOnChange}
            useList={useListProduct}
            tableComponent={ProductTable}
          />
        </div>
      </div>
    </div>
  )
}