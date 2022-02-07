interface HeaderProperties {
  children?: any;
  className?: string;
}

export const Header = ({
  children,
  className,
}: HeaderProperties) => {
  return (
    <div className={className}>
        {children}
    </div>
  );
};
