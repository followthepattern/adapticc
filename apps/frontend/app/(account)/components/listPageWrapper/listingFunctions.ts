import { SetURLSearchParams } from "react-router-dom";

export const PAGE_DEFAULT = 1;
export const PAGESIZE_DEFAULT = 10;

export interface SortLabel {
    name: string
    asc: boolean
    code: string
}

export const SetPageParams = (
    searchParams: URLSearchParams,
    setSearchParams: SetURLSearchParams,
    page: number,
) => {
    searchParams.set("page", page.toString());
    setSearchParams(searchParams);
}

export const SetSearchPatternParams = (
    searchParams: URLSearchParams,
    setSearchParams: SetURLSearchParams,
    searchString: string,
) => {
    searchParams.set("search", searchString);
    searchParams.set("page", PAGE_DEFAULT.toString());
    setSearchParams(searchParams);
}

export const SetSortPatternParrams = (
    searchParams: URLSearchParams,
    setSearchParams: SetURLSearchParams,
    sortLabel: SortLabel,
) => {
    const url = `${sortLabel.code}_${sortLabel.asc ? "asc" : "desc"}`
    searchParams.set("sort", url)
    setSearchParams(searchParams);
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

export function GetSortLabel(searchParams: URLSearchParams): SortLabel | undefined {
    const strValue = searchParams.get("sort")

    if (typeof (strValue) !== "string") {
        return undefined;
    }

    if (strValue.length < 1) {
        return undefined;
    }

    const tags = strValue.split("_");

    if (tags.length < 1) {
        return undefined;
    }

    const result: SortLabel = {
        code: tags[0],
        asc: tags[1] == "asc",
        name: ""
    }

    return result
}