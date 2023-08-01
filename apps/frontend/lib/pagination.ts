import { PAGESIZE_DEFAULT, PAGE_DEFAULT } from "./constants";

function getNumberFromSearchParams(key: string, searchParams: { [key: string]: string | string[] | undefined }, defaultValue: number): number {
    const strValue = searchParams[key];

    if (typeof (strValue) !== "string") {
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

function getNumberFromURLSearchParams(key: string, searchParams: URLSearchParams, defaultValue: number): number {
    const strValue = searchParams.get(key)

    if (typeof (strValue) !== "string") {
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

export function GetPageFromSearchParams(searchParams: URLSearchParams): number {
    return getNumberFromURLSearchParams("page", searchParams, PAGE_DEFAULT);
}

export function GetPageSizeFromSearchParams(searchParams: URLSearchParams): number {
    return getNumberFromURLSearchParams("pageSize", searchParams, PAGESIZE_DEFAULT);
}

export function GetSearch(searchParams: URLSearchParams): string {
    const strValue = searchParams.get("search")

    if (typeof (strValue) !== "string") {
        return "";
    }

    if (strValue.length == 0) {
        return "";
    }

    return strValue
}

export function CalculateMaxPage(count: number, pageSize: number): number {
    return Math.ceil(count / pageSize);
}