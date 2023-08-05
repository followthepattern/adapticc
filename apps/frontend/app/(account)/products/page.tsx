'use client';

import { useLocation, useSearchParams } from 'react-router-dom'

import { GetSearch, GetSort } from "@/lib/pagination";
import SectionHeader from '../components/sectionHeader/sectionHeader';
import List from '../components/list/list';
import { PAGE_DEFAULT } from '@/lib/constants';
import { SortLabel } from '../components/sectionHeader/components/sortButton';

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

export default function Products() {
  const [searchParams, setSearchParams] = useSearchParams();
  const searchString = GetSearch(searchParams);
  const initSort = GetSort(searchParams);

  const resourceUrl = useLocation().pathname;

  const searchInputFieldOnChange = (searchString: string) => {
    searchParams.set("search", searchString);
    searchParams.set("page", PAGE_DEFAULT.toString());
    setSearchParams(searchParams);
  }

  const sortOnChange = (sortLabel: SortLabel) => {
    const url = `${sortLabel.code}_${sortLabel.asc? "asc": "desc"}`

    searchParams.set("sort", url)

    setSearchParams(searchParams);
  }

  return (
    <div>
      <SectionHeader
        resourceName={RESOURCE_NAME}
        resourceUrl={resourceUrl}
        searchInputOnChange={searchInputFieldOnChange}
        sortOnChange={sortOnChange}
        searchInput={searchString}
        sortByLables={sortByLables}
        selectedSortLabel={initSort}
      />
      <div className="mt-8 flow-root overflow-hidden">
        <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
          <List />
        </div>
      </div>
    </div>
  )
}