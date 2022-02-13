interface FormSubmitProperties {
    children?: string;
    className?: string;
}

const FormSubmit = ({children, className}: FormSubmitProperties) => {
    return (
        <button className={className}>{children}</button>
    )
}

export default FormSubmit