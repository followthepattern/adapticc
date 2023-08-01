import Pagination from "./pagination";
import Table from "./table";

import { useEffect } from "react";
import useListProduct from "../../products/components/listProduct";
import { CalculateMaxPage, GetPageFromSearchParams, GetPageSizeFromSearchParams, GetSearch } from "@/lib/pagination";
import { Link, useLocation, useNavigate, useSearchParams } from "react-router-dom";

function getTargetUrl(path: string, params: URLSearchParams, page: string) {
    params.set("page", page);

    return `${path}?${params.toString()}`
}

export default function List() {
    const [executeQuery, { data, error, loading, itemNotFound }] = useListProduct();
    const [searchParams] = useSearchParams();
    const navigate = useNavigate();
    const pathName = useLocation().pathname;

    const page = GetPageFromSearchParams(searchParams);
    const pageSize = GetPageSizeFromSearchParams(searchParams);
    const search = GetSearch(searchParams);

    useEffect(() => {
        executeQuery({ page, pageSize, search });
    }, [page, pageSize, search])

    if (loading) {
        return <div>Loading...</div>
    }

    if (error) {
        return <div>{JSON.stringify(error)}</div>
    }

    if (itemNotFound) {
        // notFound();
    }

    const entities = data?.data ?? [];
    const count = data?.count ?? 0;

    const maxPage = CalculateMaxPage(count, pageSize);

    return (
        <>
            <Table entities={entities}/>
            {maxPage > 1 && <Pagination
                currentPage={page}
                maxPage={maxPage}
                onClick={(t: number) => { navigate(getTargetUrl(pathName, searchParams, t.toString())) }}
            />}
        </>
    )
}