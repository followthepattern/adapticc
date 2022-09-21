interface ActionsDropdownProperties {
    className?: string;
    actions: (data: any) => void[]
}

const ActionsDropdown = ({className}: ActionsDropdownProperties) => {
    return (<div className={className}></div>)
}

export default ActionsDropdown;