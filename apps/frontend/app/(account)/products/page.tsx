'use client';

import { ListPageProperties } from "@/components/listPage";
import Link from "next/link";

import Pagination from "../components/pagination";
import { notFound, usePathname, useRouter } from "next/navigation";
import { CalculateMaxPage, GetPageFromSearchParams, GetPageSizeFromSearchParams } from "@/lib/pagination";
import { stringify } from "querystring";
import SectionHeading from "../components/sectionHeading/sectionHeading";
import useListProduct from "./components/listProduct";

function getTargetUrl(path: string, searchParams: string, page: string) {
  const params = new URLSearchParams(searchParams);

  params.set("page", page);

  return `${path}?${params.toString()}`
}

export default function Page({ searchParams }: ListPageProperties) {
  const resourceName = "Products";

  const [executeQuery, { data, called, error, loading, itemNotFound }] = useListProduct();

  const pathName = usePathname();

  const router = useRouter();

  const page = GetPageFromSearchParams(searchParams);
  const pageSize = GetPageSizeFromSearchParams(searchParams);


  if (!called) {
    executeQuery(page, pageSize);
  }

  if (loading) {
    return <div>Loading...</div>
  }

  if (error) {
    return <div>{JSON.stringify(error)}</div>
  }

  if (itemNotFound) {
    notFound();
  }

  const entities = data?.data ?? [];
  const count = data?.count ?? 0;

  const maxPage = CalculateMaxPage(count, pageSize);

  return (
    <div>
      <SectionHeading resourceName={resourceName} resourceUrl={pathName}/>
      <div className="mt-8 flow-root overflow-hidden">
        <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
          <table className="w-full text-left">
            <thead className="bg-white">
              <tr>
                <th scope="col" className="relative isolate py-3.5 pr-3 text-left text-sm font-semibold text-gray-900">
                  Title
                  <div className="absolute inset-y-0 right-full -z-10 w-screen border-b border-b-gray-200" />
                  <div className="absolute inset-y-0 left-0 -z-10 w-screen border-b border-b-gray-200" />
                </th>
                <th
                  scope="col"
                  className="hidden px-3 py-3.5 text-left text-sm font-semibold text-gray-900 md:table-cell"
                >
                  Description
                </th>
                <th scope="col" className="relative py-3.5 pl-3">
                  <span className="sr-only">Edit</span>
                </th>
              </tr>
            </thead>
            <tbody>
              {entities.map((product) => (
                <tr key={product.id}>
                  <td className="relative py-4 pr-3 text-sm font-medium text-gray-900">
                    <Link href={`/products/${product.id}`} className="hover:text-indigo-900">
                      {product.title}
                    </Link>
                    <div className="absolute bottom-0 right-full h-px w-screen bg-gray-100" />
                    <div className="absolute bottom-0 left-0 h-px w-screen bg-gray-100" />
                  </td>
                  <td className="hidden px-3 py-4 text-sm text-gray-500 md:table-cell">{product.description}</td>
                  <td className="relative py-4 pl-3 text-right text-sm font-medium">
                    <Link href={`/products/${product.id}/edit`} className="text-indigo-600 hover:text-indigo-900">
                      Edit<span className="sr-only">, {product.title}</span>
                    </Link>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      </div>
      {maxPage > 1 && <Pagination
        currentPage={page}
        maxPage={maxPage}
        onClick={(t: number) => { router.push(getTargetUrl(pathName, stringify(searchParams), t.toString())) }}
      />}
    </div>
  )
}