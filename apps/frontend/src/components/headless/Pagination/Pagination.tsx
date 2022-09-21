export interface PaginationData {
  current: number;
  first: number;
  last: number;
  limit: number;
}

interface PaginationProperties {
  className?: string;
  numberButtonClassName?: string;
  actionButtonClassName?: string;
  currentButtonClassName?: string;
  renderUrl: (page: number) => string;
  paginationData: PaginationData;
}

function getPageNumbers(
  current: number,
  first: number,
  last: number,
  limit: number
): number[] {
  let result: number[] = [];
  const min = Math.max(first, current - limit);
  const max = Math.min(last, current + limit);
  console.info("max", max);
  console.info("min", min);

  for (let i = min; i <= max; i++) {
    result.push(i);
  }
  return result;
}

const Pagination = ({
  className,
  numberButtonClassName,
  actionButtonClassName,
  currentButtonClassName,
  paginationData,
  renderUrl,
}: PaginationProperties) => {
  let numbers = getPageNumbers(
    paginationData.current,
    paginationData.first,
    paginationData.last,
    paginationData.limit
  );

  const showPrev = paginationData.current != paginationData.first;
  const showNext = paginationData.current != paginationData.last;

  const showFirst =
    paginationData.current - paginationData.first > paginationData.limit;
  const showLast =
    paginationData.last - paginationData.current > paginationData.limit;

  return (
    <nav className={className}>
      {showFirst && (
        <a
          className={actionButtonClassName}
          href={renderUrl(paginationData.first)}
        >
          First
        </a>
      )}
      {showPrev && (
        <a
          className={actionButtonClassName}
          href={renderUrl(paginationData.current - 1)}
        >
          Prev
        </a>
      )}
      {numbers.map((number) => {
        return (
          <a
            className={
              number == paginationData.current
                ? currentButtonClassName
                : numberButtonClassName
            }
            href={renderUrl(number)}
          >
            {number}
          </a>
        );
      })}
      {showNext && (
        <a
          className={actionButtonClassName}
          href={renderUrl(paginationData.current + 2)}
        >
          Next
        </a>
      )}
      {showLast && (
        <a
          className={actionButtonClassName}
          href={renderUrl(paginationData.last)}
        >
          Last
        </a>
      )}
    </nav>
  );
};

export default Pagination;
