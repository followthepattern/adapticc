import SectionHeader from '../components/listPage/sectionHeader/sectionHeader';
import List from '../components/listPage/listPage';
import { ListPageComponentProperties } from '../components/listPage/listPageWrapper/listPageWrapper';
import { SortLabel, SetPageParams, SetSearchPatternParams, SetSortPatternParrams } from '../components/listPage/listPageWrapper/listingFunctions';
import useListProduct from './hooks/listProduct';
import CreateTable, { CreateTableProperties } from '../components/listPage/table/table';
import { Product } from '@/models/product';

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

function productViewLink(product: Product): string {
  return `/products/${product.id}`
}

function productEditLink(product: Product): string {
  return `/products/${product.id}/edit`
}

function getCreateTableProperties(): CreateTableProperties<Product> {
  return {
    headerColumns: ["Title", "Description"],
    getViewLink: productViewLink,
    getEditLink: productEditLink,
    getCells: (entity) => [entity.title ?? "", entity.description ?? ""]
  }
}

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

  const createTableProperties = getCreateTableProperties();

  const productTable = CreateTable(createTableProperties);

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
      <div className="flow-root mt-8 overflow-hidden">
        <div className="px-4 mx-auto max-w-7xl sm:px-6 lg:px-8">
          <List {...props}
            sortProps={{ sortLabel: selectedSortLabel }}
            onPageChange={pageOnChange}
            useList={useListProduct}
            tableComponent={productTable}
          />
        </div>
      </div>
    </div>
  )
}