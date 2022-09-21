interface SearchProperties {
    className?: string;
    placeholder?: string;
}

const Filters = ({className, placeholder}: SearchProperties) => {
    return (<input
          type="text"
          className={className}
          placeholder={placeholder}
        />)
}

export default Filters;