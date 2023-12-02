import React from "react";
import classNames from "classnames";

const Input = React.forwardRef((props: React.DetailedHTMLProps<React.InputHTMLAttributes<HTMLInputElement>, HTMLInputElement>, ref:React.ForwardedRef<HTMLInputElement>) => {
    return (
        <input
            ref={ref}
            {...props}
            className={classNames(
                props.className,
                "block w-full mt-2 border border-gray-300 rounded-lg ring-1 ring-inset ring-gray-100 placeholder:text-gray-400 focus:ring-0 focus:ring-inset",
                {
                    "cursor-not-allowed bg-gray-100 text-gray-600": props.disabled
                }
            )}
        />
    )
})

export default Input;