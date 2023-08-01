'use client';

import { useLocation, useSearchParams } from 'react-router-dom'

import { GetSearch } from "@/lib/pagination";
import SectionHeader from '../components/sectionHeader/sectionHeader';
import List from '../components/list/list';
import { PAGE_DEFAULT } from '@/lib/constants';

export default function Products() {
  const resourceName = "Products";

  const [searchParams, setSearchParams] = useSearchParams();
  const searchString = GetSearch(searchParams);

  const resourceUrl = useLocation().pathname;

  const searchInputFieldOnChange = (searchString: string) => {
    searchParams.set("search", searchString);
    searchParams.set("page", PAGE_DEFAULT.toString());
    setSearchParams(searchParams);
  }

  return (
    <div>
      <SectionHeader
        resourceName={resourceName}
        resourceUrl={resourceUrl}
        searchInputOnChange={searchInputFieldOnChange}
        searchInput={searchString}
      />
      <div className="mt-8 flow-root overflow-hidden">
        <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
          <List />
        </div>
      </div>
    </div>
  )
}