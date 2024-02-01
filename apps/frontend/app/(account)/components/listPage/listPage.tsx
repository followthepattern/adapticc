import Pagination from "./pagination/pagination";

import { useEffect } from "react";
import { ListQueryParams, ListQueryResult, ListResponse, OrderBy } from "@/graphql/utils/list";
import { ListPageComponentProperties } from "./listPageWrapper/listPageWrapper";
import TableSkeleton from "../skeletons/tableSkeleton";

interface TableProperties<T> {
    entities: T[]
}

interface ListProperties<T> extends ListPageComponentProperties {
    onPageChange: (page: number) => void
    useList: () => ListQueryResult<ListResponse<T>>
    tableComponent: React.ComponentType<TableProperties<T>>
}

export function calculateMaxPage(count: number, pageSize: number): number {
    return Math.ceil(count / pageSize);
}

export default function List<T>(props: ListProperties<T>) {
    const [executeQuery, { data, error, loading, itemNotFound }] = props.useList();

    const orderBy: OrderBy[] = [];

    if (props.sortProps.sortLabel) {
        orderBy.push({
            name: props.sortProps.sortLabel.code,
            desc: !props.sortProps.sortLabel.asc,
        })
    }

    useEffect(() => {
        const listQueryParams: ListQueryParams = {
            page: props.paginationProperties.page,
            pageSize: props.paginationProperties.pageSize,
            search: props.filterProps.searchString,
            orderBy: orderBy,
        }
        executeQuery(listQueryParams);
    }, [props.searchParams])

    if (loading) {
        return <TableSkeleton />
    }

    if (error) {
        return <div>{JSON.stringify(error)}</div>
    }

    if (itemNotFound) {
        // notFound();
    }

    const entities = data?.data ?? [];
    const count = data?.count ?? 0;

    const maxPage = calculateMaxPage(count, props.paginationProperties.pageSize);

    const Table = props.tableComponent;

    return (
        <>
            <Table entities={entities} />
            {maxPage > 1 && <div className="pt-5">
                <Pagination
                    currentPage={props.paginationProperties.page}
                    maxPage={maxPage}
                    onClick={props.onPageChange}
                />
            </div>
            }
        </>
    )
}