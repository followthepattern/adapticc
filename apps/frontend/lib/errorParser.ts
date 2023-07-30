import { WRONG_EMAIL_OR_PASSWORD } from "./constants";

export function ErrorParser(errorMsg:string): string {
    if (errorMsg === "WRONG_EMAIL_OR_PASSWORD") {
        return WRONG_EMAIL_OR_PASSWORD;
    }

    return errorMsg;
}