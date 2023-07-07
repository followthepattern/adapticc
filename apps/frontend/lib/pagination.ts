function getNumberFromSearchParams(key: string, searchParams: { [key: string]: string | string[] | undefined }, defaultValue: number): number {
    const strValue = searchParams[key];

    if (typeof(strValue) !== "string") {
        return defaultValue;
    }

    if (strValue.length == 0) {
        return defaultValue;
    }

    const intValue = parseInt(strValue);

    if (Number.isNaN(intValue)) {
        return defaultValue;
    }

    if (intValue < 1) {
        return defaultValue;
    }

    return intValue;
}

export function GetPageFromSearchParams(searchParams: { [key: string]: string | string[] | undefined }): number {
    return getNumberFromSearchParams("page", searchParams, 1);
}

export function GetPageSizeFromSearchParams(searchParams: { [key: string]: string | string[] | undefined }): number {
    return getNumberFromSearchParams("pageSize", searchParams, 10);
}

export function CalculateMaxPage(count: number, pageSize: number): number {
    return Math.ceil(count / pageSize);
}