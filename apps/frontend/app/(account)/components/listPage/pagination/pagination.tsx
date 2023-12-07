import { ArrowLongLeftIcon, ArrowLongRightIcon } from '@heroicons/react/20/solid'
import classNames from 'classnames';

interface PaginationProperties {
    currentPage: number
    maxPage: number
    onClick: (page: number) => void
}

function calculatePaginationSymbol2(totalPages: number, currentPage: number): Array<number> {
    const pageNumberLimit = 5;
    let pagination: Array<number> = [];
    const limitPerSide = Math.floor(pageNumberLimit / 2);

    const pageFrom = Math.max(Math.min(currentPage - limitPerSide, totalPages - pageNumberLimit + 1), 1)
    const pageTo = Math.min(pageFrom + pageNumberLimit - 1, totalPages);


    for (let page = pageFrom; page <= pageTo; page++) {
        pagination.push(page);
    }

    return pagination;
}

interface PageNumberSymbolProperties {
    page: number
    onClick: (page: number) => void
    currentPage: number
    isFirst: boolean
    isLast: boolean
}

function PageNumberSymbol({ onClick, page, currentPage, isFirst, isLast }: PageNumberSymbolProperties) {
    const isCurrent = page == currentPage;
    return (
        <button
            onClick={() => onClick(page)}
            className={classNames("py-2 px-4 border-y",
                {
                    "border-gray-100 focus:bg-gray-200 hover:bg-gray-100 hover:text-gray-700": !isCurrent,
                    "border-blue-500 text-blue-600 bg-blue-50": isCurrent,
                    "rounded-l-lg border-l": isFirst,
                    "rounded-r-lg border-r": isLast,
                }
            )}
            aria-current="page"
        >
            {page}
        </button>
    )
}

interface PageSymbolProperties {
    paginationSymbols: number[]
    onClick: (page: number) => void
    currentPage: number
}

function PageSymbols({ paginationSymbols, onClick, currentPage }: PageSymbolProperties) {
    return (
        <>
            {paginationSymbols.map((symbol, index) => {
                return (
                    <PageNumberSymbol
                        key={index}
                        onClick={onClick}
                        page={symbol}
                        currentPage={currentPage}
                        isFirst={index === 0}
                        isLast={index === paginationSymbols.length - 1} />
                )
            })}
        </>
    )
}

export default function Pagination(props: PaginationProperties) {
    const paginationSymbols = calculatePaginationSymbol2(props.maxPage, props.currentPage)

    const arrowButton = "py-2 px-4 border border-gray-100 rounded-lg hover:bg-gray-100 hover:text-gray-700 focus:bg-gray-200"
    const arrowIcon = "w-5 h-5 text-gray-700"

    return (
        <nav className="flex justify-center border-gray-100">
            <div className="flex-auto hidden w-0 sm:flex">
                <button
                    onClick={() => props.onClick(props.currentPage - 1)}
                    className={classNames(arrowButton, { hidden: props.currentPage <= 1 })}
                >
                    <ArrowLongLeftIcon className={arrowIcon} aria-hidden="true" />
                </button>
            </div>
            <div className="flex">
                <PageSymbols paginationSymbols={paginationSymbols} onClick={props.onClick} currentPage={props.currentPage} />
            </div>
            <div className="justify-end flex-auto hidden w-0 sm:flex">
                <button
                    className={classNames(arrowButton, { hidden: props.currentPage >= props.maxPage })}
                    onClick={() => props.onClick(props.currentPage + 1)}
                >
                    <ArrowLongRightIcon className={arrowIcon} aria-hidden="true" />
                </button>
            </div>
        </nav>
    )
}
