import { ArrowLongLeftIcon, ArrowLongRightIcon } from '@heroicons/react/20/solid'
import classNames from 'classnames';

interface PaginationProperties {
    currentPage: number
    maxPage: number
    onClick: (page: number) => void
}

function calculatePaginationSymbol(totalPages: number, currentPage: number): Array<number | string> {
    const pagination: Array<number | string> = [];

    const startPage = 1;
    const endPage = totalPages;
    const pageNumberLimit = 5;

    // only limited pages can appear
    const maxVisiblePages = (totalPages > pageNumberLimit) ? pageNumberLimit : totalPages;

    // first and the last always appear in the list
    const maxVisibleBetweenPages = maxVisiblePages - 2;

    const showFirstDots = totalPages > pageNumberLimit && currentPage - Math.floor(maxVisibleBetweenPages / 2) - 1 > startPage;
    const showLastDots = totalPages > pageNumberLimit && currentPage + Math.floor(maxVisibleBetweenPages / 2) < endPage - 1;

    pagination.push(startPage);
    if (showFirstDots) {
        pagination.push("...");
    }

    const PageFrom = showFirstDots ? Math.min(currentPage - Math.floor(maxVisibleBetweenPages / 2), endPage - maxVisibleBetweenPages) : startPage + 1;
    const PageTo = showLastDots ? Math.max(currentPage + Math.floor(maxVisibleBetweenPages / 2), startPage + maxVisibleBetweenPages) : endPage - 1;

    for (let page = PageFrom; page <= PageTo; page++) {
        pagination.push(page);
    }

    if (showLastDots) {
        pagination.push("...");
    }
    pagination.push(endPage);

    return pagination;
}

function PageDotSymbol() {
    return (
        <span className="inline-flex items-center border-t-2 border-transparent px-4 pt-4 text-sm font-medium text-gray-500">
            ...
        </span>
    )
}

interface PageNumberSymbolProperties {
    page: number
    onClick: (page: number) => void
    currentPage: number
}

function PageNumberSymbol({ onClick, page, currentPage }: PageNumberSymbolProperties) {
    const isCurrent = page == currentPage;
    return (
        <button
            onClick={() => onClick(page)}
            className={classNames("inline-flex items-center border-t-2 border-transparent px-4 pt-4 text-sm font-medium text-gray-500 hover:border-gray-300 hover:text-gray-700",
                {
                    "border-indigo-500": isCurrent,
                    "text-indigo-600": isCurrent,
                }
            )}
            aria-current="page"
        >
            {page}
        </button>
    )
}

interface PageSymbolProperties {
    paginationSymbols: (string | number)[]
    onClick: (page: number) => void
    currentPage: number
}

function PageSymbols({ paginationSymbols, onClick, currentPage }: PageSymbolProperties) {
    let keyCount = 0;
    return (
        <>
            {paginationSymbols.map(symbol => {
                keyCount++
                if (typeof (symbol) === "string") {
                    return (
                        <PageDotSymbol key={keyCount} />
                    )
                }
                return (
                    <PageNumberSymbol key={keyCount} onClick={onClick} page={symbol} currentPage={currentPage} />
                )
            })}
        </>
    )
}

export default function Pagination(props: PaginationProperties) {
    const paginationSymbols = calculatePaginationSymbol(props.maxPage, props.currentPage)

    return (
        <nav className="flex items-center justify-between border-t border-gray-200 px-4 sm:px-0">
            <div className="-mt-px flex w-0 flex-1">
                <button
                    onClick={() => props.onClick(props.currentPage - 1)}
                    className={classNames("inline-flex items-center border-t-2 border-transparent pr-1 pt-4 text-sm font-medium text-gray-500 hover:border-gray-300 hover:text-gray-700", {hidden: props.currentPage <= 1})}
                >
                    <ArrowLongLeftIcon className="mr-3 h-5 w-5 text-gray-400" aria-hidden="true" />
                    Previous
                </button>
            </div>
            <div className="hidden md:-mt-px md:flex">
                <PageSymbols paginationSymbols={paginationSymbols} onClick={props.onClick} currentPage={props.currentPage} />
            </div>
            <div className="-mt-px flex w-0 flex-1 justify-end">
                <button
                    className={classNames("inline-flex items-center border-t-2 border-transparent pl-1 pt-4 text-sm font-medium text-gray-500 hover:border-gray-300 hover:text-gray-700", {hidden: props.currentPage >= props.maxPage})}
                    onClick={() => props.onClick(props.currentPage + 1)}
                >
                    Next
                    <ArrowLongRightIcon className="ml-3 h-5 w-5 text-gray-400" aria-hidden="true" />
                </button>
            </div>
        </nav>
    )
}
