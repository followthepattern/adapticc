interface ActionButtonProperties {
    className?: string;
    children?: any
    action: ActionFunc
}

interface ActionFunc {
    (): void
}

const ActionButton = ({children, className, action}: ActionButtonProperties) => {
    return (<button onClick={action} className={className}>{children}</button>)
}

export default ActionButton;