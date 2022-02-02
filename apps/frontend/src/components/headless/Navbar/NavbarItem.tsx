interface NavbarItemProperties {
    path: string
    className: string
    children: any
}

const NavbarItem = (props: NavbarItemProperties) => {
    return (
        <a href={props.path} className={props.className}>{props.children}</a>
    )
}

export default NavbarItem;